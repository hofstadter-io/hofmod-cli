package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

  "github.com/spf13/cobra"
	"github.com/parnurzeal/gorequest"
)

var UpdateLong = `Print the build version for {{ .CLI.cliName }}`

var (
	UpdateCheckFlag bool
	UpdateServeFlag bool

	UpdateStarted bool
	UpdateErrored bool
	UpdateChecked bool
	UpdateAvailable *ProgramVersion
	UpdateData []interface{}
)

func init() {
	UpdateCmd.Flags().BoolVarP(&UpdateCheckFlag, "check", "", false, "set to only check for an update")
	UpdateCmd.Flags().BoolVarP(&UpdateServeFlag, "serve", "", false, "start an update checking server")
}

const checkURL = `{{ .CLI.Updates.CheckURL }}`
const fetchURL = `{{ .CLI.Updates.FetchURL }}`

const updateMessage = `
Updates available. v%s -> %s (latest)

  run '{{ .CLI.cliName }} update' to get the latest.

`

// TODO, add a curl to the above? or os specific?

var UpdateCmd = &cobra.Command{

	Use: "update",

	Short: "update the dma tool",

	Long: UpdateLong,

	Run: func(cmd *cobra.Command, args []string) {
		if UpdateServeFlag {
			err := ServeUpdates()
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}

			return
		}

		latest, err := CheckUpdate(true)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		// Semver Check?
		cur := ProgramVersion{Version: "v" + Version}
		if latest.Version == cur.Version {
			return
		} else {
			// This will actually be called on the root persistent post run anyway
			// fmt.Printf(updateMessage, Version, latest.Version)
			if UpdateCheckFlag {
				return
			}
		}

		err = InstallLatest()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	},
}

func init() {
	RootCmd.AddCommand(UpdateCmd)

	{{ if not .CLI.Updates.Disabled }}
	go CheckUpdate(false)
	{{ end }}
}

type ProgramVersion struct {
	Version string
	URL     string
}

func CheckUpdate(manual bool) (ver ProgramVersion, err error) {
	UpdateStarted = true
	cur := ProgramVersion{Version: "v" + Version}

	req := gorequest.New().Get(checkURL).
		Query("current="+cur.Version).
		Query("manual="+fmt.Sprint(manual))
	{{ if not .CLI.Updates.TelemetryDisabled }}
	req.Query("args=" + url.QueryEscape(strings.Join(os.Args, " ")))
	{{ end }}
	resp, b, errs := req.EndBytes()
	UpdateErrored = true

	check := "http2: server sent GOAWAY and closed the connection"
	if len(errs) != 0 && !strings.Contains(errs[0].Error(), check) {
		// fmt.Println("errs:", errs)
		return ver, errs[0]
	}

	if len(errs) != 0 || resp.StatusCode >= 500 {
		return ver, fmt.Errorf("Internal Error: " + string(b))
	}
	if resp.StatusCode >= 400 {
		if resp.StatusCode == 404 {
			return ver, fmt.Errorf("No releases available :[")
		}
		return ver, fmt.Errorf("Bad Request: " + string(b))
	}
	UpdateErrored = false

	if !manual {
		ver.Version = string(b)
		UpdateChecked = true

		// Semver Check?
		if ver.Version != cur.Version {
			UpdateAvailable = &ver
		}

		return ver, nil
	}

	var gh map[string]interface{}
	err = json.Unmarshal(b, &gh)
	if err != nil {
		return ver, err
	}

	nameI, ok := gh["name"]
	if !ok {
		return ver, fmt.Errorf("Internal Error: could not find version in update check response")
	}
	name, ok := nameI.(string)
	if !ok {
		return ver, fmt.Errorf("Internal Error: version is not a string in update check response")
	}

	// TODO, pick correct URL from reponse by build os / arch

	ver.Version = name
	UpdateChecked = true

	// Semver Check?
	if ver.Version != cur.Version {
		UpdateAvailable = &ver
		aI, ok := gh["assets"]
		if ok {
			a, aok := aI.([]interface{})
			if aok {
				UpdateData = a
			}
		}
	}

	return ver, nil
}

func WaitPrintUpdateAvail() {
	for i := 0; i < 20 && !UpdateStarted && !UpdateChecked && !UpdateErrored; i++ {
		time.Sleep(50 * time.Millisecond)
	}
	PrintUpdateAvailable()
}

func PrintUpdateAvailable() {
	if UpdateChecked && UpdateAvailable != nil {
		fmt.Printf(updateMessage, Version, UpdateAvailable.Version)
	}
}

func InstallLatest() (err error) {
	fmt.Printf("Installing {{ .CLI.cliName }}@%s\n", UpdateAvailable.Version)

	if UpdateData == nil {
		return fmt.Errorf("No update available")
	}
	/*
	vers, err := json.MarshalIndent(UpdateData, "", "  ")
	if err == nil {
		fmt.Println(string(vers))
	}
	*/

	fmt.Println("OS/Arch", BuildOS, BuildArch)

	url := ""
	for _, Asset := range UpdateData {
		asset := Asset.(map[string]interface{})
		U := asset["browser_download_url"].(string)
		u := strings.ToLower(U)

		fmt.Println("  url:", u)

		osOk, archOk := false, false

		switch BuildOS {
		case "linux":
			if strings.Contains(u, "windows") {
				osOk = true
			}

		case "darwin":
			if strings.Contains(u, "darwin") {
				osOk = true
			}

		case "windows":
			if strings.Contains(u, "windows") {
				osOk = true
			}
		}

		switch BuildArch {
			case "amd64":
			if strings.Contains(u, "x86_64"){
				archOk = true
			}
			case "arm64":
			if strings.Contains(u, "arm64"){
				archOk = true
			}
			case "arm":
			if strings.Contains(u, "arm") && !strings.Contains(u, "arm64"){
				archOk = true
			}
		}

		if osOk && archOk {
			url = u
		}
	}

	fmt.Println("URL: ", url)

	return nil
}

func ServeUpdates() (err error) {
	fmt.Printf("Serving...\n")

	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		log.Println("CHECK", r.RemoteAddr, q)

		if q.Get("manual") != "true" {
			fmt.Fprint(w, "v" + Version)
			return
		}

		// url := "https://api.github.com/repos/{{ .CLI.Releases.GitHub.Owner }}/{{ .CLI.Releases.GitHub.Repo }}/releases/latest"
		url := "https://api.github.com/repos/{{ .CLI.Releases.GitHub.Owner }}/hof/releases/latest"
		req := gorequest.New().Get(url)
		resp, body, errs := req.End()

		check := "http2: server sent GOAWAY and closed the connection"
		if len(errs) != 0 && !strings.Contains(errs[0].Error(), check) {
			fmt.Println("errs:", errs)
			fmt.Println("resp:", resp)
			fmt.Println("body:", body)
			return
		}

		if len(errs) != 0 || resp.StatusCode >= 500 {
			http.Error(w, body, resp.StatusCode)
			return
			// return ver, fmt.Errorf("Internal Error: " + body)
		}
		if resp.StatusCode >= 400 {
			http.Error(w, body, resp.StatusCode)
			return
		}

		fmt.Fprint(w, body)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

	return nil
}

