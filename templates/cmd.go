{{ if .CMD.Parent }}
package cmd{{ .CMD.Parent.Name }}
{{ else }}
package cmd
{{ end }}

import (
	{{ if .CMD.HasAnyRun }}
	"fmt"
	"os"
	{{end}}
	{{ if .CLI.Telemetry }}
	"strings"
	{{end}}

  "github.com/spf13/cobra"
  {{ if or .CMD.Flags .CMD.Pflags }}
  "github.com/spf13/viper"
  {{ end }}

	{{ if .CMD.Imports }}
	{{ range $i, $I := .CMD.Imports }}
	{{ $I.As }} "{{ $I.Path }}"
	{{ end }}
	{{ end }}

	{{ if .CMD.Commands }}
  {{ if .CMD.Parent.Parent.Parent.Parent }}
	"{{ .CLI.Package }}/cmd/{{ .CMD.Parent.Parent.Parent.Parent.Name }}/{{ .CMD.Parent.Parent.Parent.Name }}/{{ .CMD.Parent.Parent.Name }}/{{ .CMD.Parent.Name }}/{{ .CMD.cmdName }}"
  {{ else if .CMD.Parent.Parent.Parent }}
	"{{ .CLI.Package }}/cmd/{{ .CMD.Parent.Parent.Parent.Name }}/{{ .CMD.Parent.Parent.Name }}/{{ .CMD.Parent.Name }}/{{ .CMD.cmdName }}"
  {{ else if .CMD.Parent.Parent }}
	"{{ .CLI.Package }}/cmd/{{ .CMD.Parent.Parent.Name }}/{{ .CMD.Parent.Name }}/{{ .CMD.cmdName }}"
  {{ else if .CMD.Parent }}
	"{{ .CLI.Package }}/cmd/{{ .CMD.Parent.Name }}/{{ .CMD.cmdName }}"
  {{ else }}
	"{{ .CLI.Package }}/cmd/{{ .CMD.cmdName }}"
  {{ end }}
	{{ end }}

	{{ if .CLI.Telemetry }}
	"{{ .CLI.Package }}/ga"
	{{end}}
	{{ if .CMD.Pflags }}
	"{{ .CLI.Package }}/pflags"
	{{ end }}
)

{{ if .CMD.Long }}
var {{ .CMD.Name }}Long = `{{ .CMD.Long }}`
{{ end }}

{{ template "flag-vars" .CMD }}
{{ template "flag-init" .CMD }}
{{ template "pflag-init" .CMD }}

{{ if .CMD.PersistentPrerun }}
func {{ .CMD.CmdName }}PersistentPreRun({{- template "lib-args.go" . -}}) (err error) {
	{{ if .CMD.PersistentPrerunBody }}
	{{ .CMD.PersistentPrerunBody }}
	{{ end }}

	return err
}
{{ end }}

{{ if .CMD.Prerun}}
func {{ .CMD.CmdName }}PreRun({{- template "lib-args.go" . -}}) (err error) {
	{{ if .CMD.PrerunBody }}
	{{ .CMD.PrerunBody }}
	{{ end }}

	return err
}
{{ end }}

{{ if not .CMD.OmitRun}}
func {{ .CMD.CmdName }}Run({{ template "lib-args.go" . -}}) (err error) {

	{{ if .CMD.Body}}
	{{ .CMD.Body}}
	{{ end }}

	return err
}
{{ end }}

{{ if .CMD.PersistentPostrun}}
func {{ .CMD.CmdName }}PersistentPostRun({{- template "lib-args.go" . -}}) (err error) {

	{{ if .CMD.PersistentPostrunBody}}
	{{ .CMD.PersistentPostrunBody}}
	{{ end }}

	return err
}
{{ end }}

{{ if .CMD.Postrun}}
func {{ .CMD.CmdName }}PostRun({{- template "lib-args.go" . -}}) (err error) {

	{{ if .CMD.PostrunBody }}
	{{ .CMD.PostrunBody }}
	{{ end }}

	return err
}
{{ end }}
var {{ .CMD.CmdName }}Cmd = &cobra.Command{

  {{ if .CMD.Usage}}
  Use: "{{ .CMD.Usage }}",
  {{ else }}
  Use: "{{ .CMD.Name }}",
  {{ end }}

	{{ if .CMD.Hidden }}
	Hidden: true,
	{{ end }}

	{{ if .CMD.Aliases }}
	Aliases: []string{
		{{range $i, $AL := .CMD.Aliases}}"{{$AL}}",
		{{end}}
	},
	{{ end }}

  {{ if .CMD.Short}}
  Short: "{{ .CMD.Short }}",
  {{ end }}

  {{ if .CMD.Long }}
  Long: {{ .CMD.Name }}Long,
  {{ end }}

  {{ if .CMD.PersistentPrerun }}
  PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
    {{ template "args-parse" .CMD.Args }}

		err = {{ .CMD.CmdName }}PersistentPreRun({{ template "lib-call.go" . }})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
  },
  {{ end }}

{{ if or .CMD.Prerun .CLI.Telemetry}}
  PreRun: func(cmd *cobra.Command, args []string) {
		{{ if .CLI.Telemetry }}
		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c, "<omit>", 0)
		{{ end }}

		{{ if .CMD.Prerun}}
		var err error
    {{ template "args-parse" .CMD.Args }}

		err = {{ .CMD.CmdName }}PreRun({{ template "lib-call.go" . }})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		{{ end }}
  },
  {{ end }}

  {{ if not .CMD.OmitRun}}
  Run: func(cmd *cobra.Command, args []string) {
		var err error
    {{ template "args-parse" .CMD.Args }}

		err = {{ .CMD.CmdName }}Run({{ template "lib-call.go" . }})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
  },
  {{ end }}

  {{ if .CMD.PersistentPostrun}}
  PersistentPostRun: func(cmd *cobra.Command, args []string) {
		var err error
    {{ template "args-parse" .CMD.Args }}

		err = {{ .CMD.CmdName }}PersistentPostRun({{ template "lib-call.go" . }})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
  },
  {{ end }}

  {{ if .CMD.Postrun}}
  PostRun: func(cmd *cobra.Command, args []string) {
		var err error
    {{ template "args-parse" .CMD.Args }}

		err = {{ .CMD.CmdName }}PostRun({{ template "lib-call.go" . }})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
  },
  {{ end }}
}

func init() {
	{{ if .CMD.CustomHelp }}
	help := func (cmd *cobra.Command, args []string) {
		fu := {{ $.CMD.CmdName }}Cmd.Flags().FlagUsages()
		ch := strings.Replace({{ $.CMD.CmdName }}CustomHelp, "<<flag-usage>>", fu, 1)
		fmt.Println(ch)
	}
	usage := func (cmd *cobra.Command) error {
		fu := {{ $.CMD.CmdName }}Cmd.Flags().FlagUsages()
		ch := strings.Replace({{ $.CMD.CmdName }}CustomHelp, "<<flag-usage>>", fu, 1)
		fmt.Println(ch)
		return fmt.Errorf("unknown HOF command")
	}
	{{ else }}
	help := {{ $.CMD.CmdName }}Cmd.HelpFunc()
	usage := {{ $.CMD.CmdName }}Cmd.UsageFunc()
	{{ end }}

	{{ if .CLI.Telemetry }}
	thelp := func (cmd *cobra.Command, args []string) {
		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c + "/help", "<omit>", 0)
		help(cmd, args)
	}
	tusage := func (cmd *cobra.Command) error {
		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c + "/help", "<omit>", 0)
		return usage(cmd)
	}
	{{ $.CMD.CmdName }}Cmd.SetHelpFunc(thelp)
	{{ $.CMD.CmdName }}Cmd.SetUsageFunc(tusage)
	{{ else }}
	{{ $.CMD.CmdName }}Cmd.SetHelpFunc(help)
	{{ $.CMD.CmdName }}Cmd.SetUsageFunc(usage)
	{{ end }}

{{if .CMD.Commands}}
  {{- range $i, $C := .CMD.Commands }}
  {{ $.CMD.CmdName }}Cmd.AddCommand(cmd{{ $.CMD.cmdName }}.{{ $C.CmdName }}Cmd)
  {{- end}}
{{ end }}
}

{{ if .CMD.CustomHelp }}
const {{ $.CMD.CmdName }}CustomHelp = `{{ .CMD.CustomHelp }}`
{{ end }}
