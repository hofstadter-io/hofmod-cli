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

  "github.com/spf13/cobra"
  {{ if or .CMD.Flags .CMD.Pflags }}
  "github.com/spf13/viper"
  {{ end }}

	{{ if .CMD.Commands }}
  {{ if .CMD.Parent.Parent.Parent }}
	"{{ .CLI.Package }}/cmd/{{ .CMD.Parent.Parent.Parent.Name }}/{{ .CMD.Parent.Parent.Name }}/{{ .CMD.Parent.Name }}/{{ .CMD.cmdName }}"
  {{ else if .CMD.Parent.Parent }}
	"{{ .CLI.Package }}/cmd/{{ .CMD.Parent.Parent.Name }}/{{ .CMD.Parent.Name }}/{{ .CMD.cmdName }}"
  {{ else if .CMD.Parent }}
	"{{ .CLI.Package }}/cmd/{{ .CMD.Parent.Name }}/{{ .CMD.cmdName }}"
  {{ else }}
	"{{ .CLI.Package }}/cmd/{{ .CMD.cmdName }}"
  {{ end }}
	{{ end }}

	/*
		{{ .CMD.PersistentPrerun }}
		{{ .CMD.Prerun }}
		{{ .CMD.OmitRun }}
		{{ .CMD.PersistentPostrun }}
		{{ .CMD.Postrun }}
		{{ .CMD.HasAnyRun }}
	*/

	{{ if .CMD.HasAnyRun }}
  {{ if .CMD.Parent.Parent.Parent }}
	"{{ .CLI.Package }}/lib/cmd/{{ .CMD.Parent.Parent.Parent.Name }}/{{ .CMD.Parent.Parent.Name }}/{{ .CMD.Parent.Name }}"
  {{ else if .CMD.Parent.Parent }}
	"{{ .CLI.Package }}/lib/cmd/{{ .CMD.Parent.Parent.Name }}/{{ .CMD.Parent.Name }}"
  {{ else if .CMD.Parent }}
	"{{ .CLI.Package }}/lib/cmd/{{ .CMD.Parent.Name }}"
  {{ else }}
	"{{ .CLI.Package }}/lib/cmd"
  {{ end }}
	{{ end }}
)

{{ if .CMD.Long }}
var {{ .CMD.Name }}Long = `{{ .CMD.Long }}`
{{ end }}

{{ template "flag-vars" .CMD }}
{{ template "flag-init" .CMD }}

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

		{{ if .CMD.Parent }}
		err = libcmd{{ .CMD.Parent.Name }}.{{ .CMD.CmdName }}PersistentPreRun({{ template "lib-call.go" . }})
		{{ else }}
		err = libcmd.{{ .CMD.CmdName }}PersistentPreRun({{ template "lib-call.go" . }})
		{{ end }}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
  },
  {{ end }}

  {{ if .CMD.Prerun }}
  PreRun: func(cmd *cobra.Command, args []string) {
		var err error
    {{ template "args-parse" .CMD.Args }}

		{{ if .CMD.Parent }}
		err = libcmd{{ .CMD.Parent.Name }}.{{ .CMD.CmdName }}PreRun({{ template "lib-call.go" . }})
		{{ else }}
		err = libcmd.{{ .CMD.CmdName }}PreRun({{ template "lib-call.go" . }})
		{{ end }}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
  },
  {{ end }}

  {{ if not .CMD.OmitRun}}
  Run: func(cmd *cobra.Command, args []string) {
		var err error
    {{ template "args-parse" .CMD.Args }}

		{{ if .CMD.Parent }}
		err = libcmd{{ .CMD.Parent.Name }}.{{ .CMD.CmdName }}Run({{ template "lib-call.go" . }})
		{{ else }}
		err = libcmd.{{ .CMD.CmdName }}Run({{ template "lib-call.go" . }})
		{{ end }}

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

		{{ if .CMD.Parent }}
		err = libcmd{{ .CMD.Parent.Name }}.{{ .CMD.CmdName }}PostRun({{ template "lib-call.go" . }})
		{{ else }}
		err = libcmd.{{ .CMD.CmdName }}PostRun({{ template "lib-call.go" . }})
		{{ end }}

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

		{{ if .CMD.Parent }}
		err = libcmd{{ .CMD.Parent.Name }}.{{ .CMD.CmdName }}PostRun(A, F)
		{{ else }}
		err = libcmd.{{ .CMD.CmdName }}PostRun(A, F)
		{{ end }}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
  },
  {{ end }}
}

{{if .CMD.Commands}}
func init() {
	{{- range $i, $C := .CMD.Commands }}
  {{ $.CMD.CmdName }}Cmd.AddCommand(cmd{{ $.CMD.cmdName }}.{{ $C.CmdName }}Cmd)
	{{- end}}
}
{{ end }}
