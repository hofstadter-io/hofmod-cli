package cmd

import (
	"fmt"
	"os"
	{{ if .CLI.EnablePProf }}
	"log"
	"runtime/pprof"
	{{end}}
	{{ if .CLI.CustomHelp }}
	"strings"
	{{ end }}

	"github.com/rogpeppe/go-internal/testscript"
	"github.com/spf13/cobra"

	{{ if .CLI.Imports }}
	{{ range $i, $I := .CLI.Imports }}
	{{ $I.As }} "{{ $I.Path }}"
	{{ end }}
	{{ end }}

	{{ if .CLI.Telemetry }}
	"{{ .CLI.Package }}/ga"
	{{end}}
	{{ if or .CLI.Flags .CLI.Pflags }}
	"{{ .CLI.Package }}/flags"
	{{ end }}
)

{{ if .CLI.Long }}
var {{ .CLI.Name }}Long = `{{ .CLI.Long }}`
{{ end }}

{{ template "flag-init" .CLI }}
{{ template "pflag-init" .CLI }}

{{ if .CLI.PersistentPrerun }}
func RootPersistentPreRun({{- template "lib-args.go" . -}}) (err error) {
	{{ if .CLI.PersistentPrerunBody }}
	{{ .CLI.PersistentPrerunBody }}
	{{ end }}

	return err
}
{{ end }}

{{ if .CLI.Prerun }}
func RootPreRun({{- template "lib-args.go" . -}}) (err error) {
	{{ if .CLI.PrerunBody }}
	{{ .CLI.PrerunBody }}
	{{ end }}

	return err
}
{{ end }}

{{ if not .CLI.OmitRun}}
func RootRun({{ template "lib-args.go" . -}}) (err error) {

	{{ if .CLI.Body}}
	{{ .CLI.Body}}
	{{ end }}

	return err
}
{{ end }}

{{ if or .CLI.PersistentPostrun .CLI.Updates}}
func RootPersistentPostRun({{- template "lib-args.go" . -}}) (err error) {

	{{ if .CLI.Updates }}
	WaitPrintUpdateAvailable()
	{{ end }}

	{{ if .CLI.PersistentPostrunBody}}
	{{ .CLI.PersistentPostrunBody}}
	{{ end }}

	return err
}
{{ end }}

{{ if .CLI.Postrun}}
func RootPostRun({{- template "lib-args.go" . -}}) (err error) {

	{{ if .CLI.PostrunBody }}
	{{ .CLI.PostrunBody }}
	{{ end }}

	return err
}
{{ end }}

var RootCmd = &cobra.Command{

  {{ if .CLI.Usage}}
  Use: "{{ .CLI.Usage }}",
  {{ else }}
  Use: "{{ .CLI.Name }}",
  {{ end }}

  {{ if .CLI.Short}}
  Short: "{{ .CLI.Short }}",
  {{ end }}

  {{ if .CLI.Long }}
  Long: {{ .CLI.Name }}Long,
  {{ end }}

  {{ if .CLI.PersistentPrerun }}
  PersistentPreRun: func(cmd *cobra.Command, args []string) {
	var err error
    {{ template "args-parse" .CLI.Args }}

	err = RootPersistentPreRun({{ template "lib-call.go" .CLI.Args }})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
  },
  {{ end }}

{{ if or .CLI.Prerun .CLI.Telemetry}}
  PreRun: func(cmd *cobra.Command, args []string) {
	{{ if .CLI.Telemetry }}
	ga.SendGaEvent("root", "<omit>", 0)
	{{ end }}

	{{ if .CLI.Prerun}}
	var err error
    {{ template "args-parse" .CLI.Args }}

	err = RootPreRun({{ template "lib-call.go" .CLI.Args }})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	{{ end }}
  },
  {{ end }}

  {{ if not .CLI.OmitRun}}
  Run: func(cmd *cobra.Command, args []string) {
	var err error
    {{ template "args-parse" .CLI.Args }}

	err = RootRun({{ template "lib-call.go" .CLI.Args }})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
  },
  {{ end }}

  {{ if or .CLI.PersistentPostrun .CLI.Updates }}
  PersistentPostRun: func(cmd *cobra.Command, args []string) {
	var err error
    {{ template "args-parse" .CLI.Args }}

	err = RootPersistentPostRun({{ template "lib-call.go" .CLI.Args }})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
  },
  {{ end }}

  {{ if .CLI.Postrun }}
  PostRun: func(cmd *cobra.Command, args []string) {
	var err error
    {{ template "args-parse" .CLI.Args }}

	err = RootPostRun({{ template "lib-call.go" .CLI.Args }})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
  },
  {{ end }}
}

func RootInit() {
	{{ if .CLI.CustomHelp }}
	help := func (cmd *cobra.Command, args []string) {
		fu := RootCmd.Flags().FlagUsages()
		rh := strings.Replace(RootCustomHelp, "<<flag-usage>>", fu, 1)
		fmt.Println(rh)
		fmt.Println(cmd.Name(), "hof", args)
	}
	usage := func(cmd *cobra.Command) error {
		fu := RootCmd.Flags().FlagUsages()
		rh := strings.Replace(RootCustomHelp, "<<flag-usage>>", fu, 1)
		fmt.Println(rh)
		return fmt.Errorf("unknown {{ .CLI.cliName }} command")
	}
	{{ else }}
	help := RootCmd.HelpFunc()
	usage := RootCmd.UsageFunc()
	{{ end }}

	{{ if .CLI.Telemetry }}
	thelp := func (cmd *cobra.Command, args []string) {
		if RootCmd.Name() == cmd.Name() {
			ga.SendGaEvent("root/help", "<omit>", 0)
		}
		help(cmd, args)
	}
	tusage := func (cmd *cobra.Command) error {
		if RootCmd.Name() == cmd.Name() {
			ga.SendGaEvent("root/usage", "<omit>", 0)
		}
		return usage(cmd)
	}
	RootCmd.SetHelpFunc(thelp)
	RootCmd.SetUsageFunc(tusage)
	{{ else }}
	RootCmd.SetHelpFunc(help)
	RootCmd.SetUsageFunc(usage)
	{{ end }}


{{if .CLI.Updates}}
	RootCmd.AddCommand(UpdateCmd)
{{end}}
{{if .CLI.VersionCommand}}
	RootCmd.AddCommand(VersionCmd)
{{end}}
{{if .CLI.CompletionCommands}}
	RootCmd.AddCommand(CompletionCmd)
{{end}}

{{if .CLI.Commands}}
	{{range $i, $C := .CLI.Commands }}
	RootCmd.AddCommand({{ $C.CmdName }}Cmd)
	{{- end }}
{{ end }}
}

{{ if .CLI.CustomHelp }}
const RootCustomHelp = `{{ .CLI.CustomHelp }}`
{{ end }}

func RunExit() {
	if err := RunErr(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func RunInt() int {
	if err := RunErr(); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func RunErr() error {
	{{ if .CLI.EnablePProf }}
	if fn := os.Getenv("{{.CLI.CLI_NAME}}_CPU_PROFILE"); fn != "" {
		f, err := os.Create(fn)
		if err != nil {
			log.Fatal("Could not create file for CPU profile:", err)
		}
		defer f.Close()

		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal("Could not start CPU profile process:", err)
		}

		defer pprof.StopCPUProfile()
	}
	{{ end }}

	RootInit()
	return RootCmd.Execute()
}

func CallTS(ts* testscript.TestScript, args[]string) error {
	RootCmd.SetArgs(args)

	err := RootCmd.Execute()
	ts.Check(err)

	return err
}
