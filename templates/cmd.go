{{ if .CMD.Parent }}
package cmd{{ .CMD.Parent.Name }}
{{ else }}
package cmd
{{ end }}

import (
	{{ if .CMD.CustomHelp }}
	{{ if not .CMD.HasAnyRun }}
	"fmt"
	{{end}}
	{{end}}
	{{ if .CMD.HasAnyRun }}
	"fmt"
	"os"
	{{end}}

  "github.com/spf13/cobra"

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

	{{/* hack */}}
	{{ if .CMD.Flags }}
	"{{ .CLI.Package }}/flags"
	{{ else if .CMD.Pflags }}
	"{{ .CLI.Package }}/flags"
	{{ else if .CMD.Topics }}
	"{{ .CLI.Package }}/flags"
	{{ else if .CMD.Examples }}
	"{{ .CLI.Package }}/flags"
	{{ else if .CMD.Tutorials }}
	"{{ .CLI.Package }}/flags"
	{{ end }}
	{{ if .CLI.Telemetry }}
	"{{ .CLI.Package }}/ga"
	{{end}}
)

{{ if .CMD.Long }}
var {{ .CMD.Name }}Long = `{{ .CMD.Long }}`
{{ end }}

{{ template "pflag-init" .CMD }}
{{ template "flag-init" .CMD }}

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
	{{ else }}
	// you can safely comment this print out
	fmt.Println("not implemented")
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
		ga.SendCommandPath(cmd.CommandPath())
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
	extra := func(cmd *cobra.Command) bool {
		{{ if .CMD.Topics }}
		if flags.PrintSubject("Topics\n", "", flags.RootPflags.Topic, {{ .CMD.CmdName }}Topics) {
			return true
		}
		{{ end }}

		{{ if .CMD.Examples }}
		if flags.PrintSubject("Examples\n", "", flags.RootPflags.Example, {{ .CMD.CmdName }}Examples) {
			return true
		}
		{{ end }}

		{{ if .CMD.Tutorials }}
		if flags.PrintSubject("Tutorials\n", "", flags.RootPflags.Tutorial, {{ .CMD.CmdName }}Tutorials) {
			return true
		}
		{{ end }}

		return false
	}
	{{ if .CMD.CustomHelp }}
	help := func (cmd *cobra.Command, args []string) {
		if extra(cmd) {
			return
		}
		fu := {{ $.CMD.CmdName }}Cmd.Flags().FlagUsages()
		ch := strings.Replace({{ $.CMD.CmdName }}CustomHelp, "<<flag-usage>>", fu, 1)
		fmt.Println(ch)
		{{ if .CMD.TBDLong }}
		fmt.Println("\nstatus: {{ .CMD.TBDLong }}")
		{{ end }}
	}
	usage := func (cmd *cobra.Command) error {
		if extra(cmd) {
			return nil
		}
		fu := {{ $.CMD.CmdName }}Cmd.Flags().FlagUsages()
		ch := strings.Replace({{ $.CMD.CmdName }}CustomHelp, "<<flag-usage>>", fu, 1)
		fmt.Println(ch)
		{{ if .CMD.TBDLong }}
		fmt.Println("\nstatus: {{ .CMD.TBDLong }}")
		{{ end }}
		return fmt.Errorf("unknown command %q", cmd.Name())
	}
	{{ else }}
	ohelp := {{ $.CMD.CmdName }}Cmd.HelpFunc()
	ousage := {{ $.CMD.CmdName }}Cmd.UsageFunc()
	help := func (cmd *cobra.Command, args []string) {
		if extra(cmd) {
			return
		}
		ohelp(cmd, args)
	}
	usage := func(cmd *cobra.Command) error {
		if extra(cmd) {
			return nil
		}
		return ousage(cmd)
	}
	{{ end }}

	{{ if .CLI.Telemetry }}
	thelp := func (cmd *cobra.Command, args []string) {
		ga.SendCommandPath(cmd.CommandPath() + " help")
		help(cmd, args)
	}
	tusage := func (cmd *cobra.Command) error {
		ga.SendCommandPath(cmd.CommandPath() + " usage")
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

{{ if .CMD.Topics }}
var {{ .CMD.CmdName }}Topics = map[string]string {
  {{- range $k, $v := .CMD.Topics }}
  "{{ $k }}": `{{ replace $v "`" "¡" -1 }}`,
  {{- end}}
}
{{ end }}

{{ if .CMD.Examples }}
var {{ .CMD.CmdName }}Examples = map[string]string {
  {{- range $k, $v := .CMD.Examples }}
  "{{ $k }}": `{{ replace $v "`" "¡" -1 }}`,
  {{- end}}
}
{{ end }}

{{ if .CMD.Tutorials }}
var {{ .CMD.CmdName }}Tutorials = map[string]string {
  {{- range $k, $v := .CMD.Tutorials }}
  "{{ $k }}": `{{ replace $v "`" "¡" -1 }}`,
  {{- end}}
}
{{ end }}
