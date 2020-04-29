package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
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

var checkURL = `{{ .CLI.Updates.CheckURL }}`

func init () {
	if Commit == "Dirty" {
		checkURL = `{{ .CLI.Updates.DevCheckURL }}`
	}
}

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
	if !manual && os.Getenv("{{ .CLI.CLI_NAME }}_UPDATES_DISABLED") != "" {
		return
	}
	UpdateStarted = true
	cur := ProgramVersion{Version: "v" + Version}

	req := gorequest.New().Get(checkURL).
		Query("current="+cur.Version).
		Query("manual="+fmt.Sprint(manual))
	{{ if not .CLI.Updates.TelemetryDisabled }}
	if os.Getenv("{{ .CLI.CLI_NAME }}_TELEMETRY_DISABLED") != "" {
		req.Query("args=" + url.QueryEscape(strings.Join(os.Args, " ")))
	}
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

		osOk, archOk := false, false

		switch BuildOS {
		case "linux":
			if strings.Contains(u, "linux") {
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
			break
		}
	}

	fmt.Println("Download URL: ", url, "\n")

		switch BuildOS {
		case "linux":
			fallthrough
		case "darwin":

			return downloadAndInstall(url)

		case "windows":
			fmt.Println("Please downlaod and install manually from the link above.\n")
			return nil
		}


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

		url := "https://api.github.com/repos/{{ .CLI.Releases.GitHub.Owner }}/{{ .CLI.Releases.GitHub.Repo }}/releases/latest"
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


func downloadAndInstall(url string) error {
	req := gorequest.New().Get(url)
	resp, content, errs := req.EndBytes()

	check := "http2: server sent GOAWAY and closed the connection"
	if len(errs) != 0 && !strings.Contains(errs[0].Error(), check) {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		return errs[0]
	}

	if len(errs) != 0 || resp.StatusCode >= 400 {
		return fmt.Errorf("Error %v - %s", resp.StatusCode, string(content))
	}

	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		return err
	}

	// defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		return err
	}
	if err := tmpfile.Close(); err != nil {
		return err
	}

	ex, err := os.Executable()
	if err != nil {
		return err
	}

	real, err := filepath.EvalSymlinks(ex)
	if err != nil {
		return err
	}

	// Sudo copy the file
	cmd := exec.Command("/bin/sh", "-c",
		fmt.Sprintf("export OWNER=$(ls -l %s | awk '{ print $3 \":\" $4 }') && sudo mv %s %s.backup && sudo cp %s %s && sudo chown $OWNER %s && sudo chmod 0755 %s && sudo rm %s.backup",
			real, // get owner
			real, real, // mv
			tmpfile.Name(), real, // cp
			real, // chown
			real, // chmod
			real, // rm
		),
	)

	// prep stdin for password
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", stdoutStderr)
	if err != nil {
		return err
	}

	UpdateAvailable = nil
	UpdateData = nil
	return nil
}
