package templates

import (
  "github.com/hofstadter-io/hofmod-cli/partials"
)

CommandTemplate :: partials.AllPartials + RealCommandTemplate

RealCommandTemplate :: """
{{ if .CMD.Parent }}
package {{ .CMD.Parent.Name }}
{{ else }}
package commands
{{ end }}

import (
  {{ if or .CMD.OmitRun .CMD.Body }}
  // hello... something might need to go here
  {{ else }}
  "fmt"
  {{end}}

  {{ $already := false }}
  {{ range $i, $A := .CMD.Args }}
    {{ if $already }}
    {{ else }}
      {{ if $A.Required }}
        {{ $already = true }}
        "os"
      {{ end }}
    {{ end }}
  {{ end }}

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
  {{ if .CMD.Parent.Parent.Parent }}
	"{{ .CLI.Package }}/commands/{{ .CMD.Parent.Parent.Parent.Name }}/{{ .CMD.Parent.Parent.Name }}/{{ .CMD.Parent.Name }}/{{ .CMD.cmdName }}"
  {{ else if .CMD.Parent.Parent }}
	"{{ .CLI.Package }}/commands/{{ .CMD.Parent.Parent.Name }}/{{ .CMD.Parent.Name }}/{{ .CMD.cmdName }}"
  {{ else if .CMD.Parent }}
	"{{ .CLI.Package }}/commands/{{ .CMD.Parent.Name }}/{{ .CMD.cmdName }}"
  {{ else }}
	"{{ .CLI.Package }}/commands/{{ .CMD.cmdName }}"
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
    {{ template "args-parse" .CMD.Args }}

    {{ if .CMD.PersistentPrerunBody }}
    {{ .CMD.PersistentPrerunBody }}
    {{ end }}
  },
  {{ end }}

  {{ if .CMD.Prerun }}
  PreRun: func(cmd *cobra.Command, args []string) {
    {{ template "args-parse" .CMD.Args }}

    {{ if .CMD.PrerunBody }}
    {{ .CMD.PrerunBody }}
    {{ end }}
  },
  {{ end }}

  {{ if not .CMD.OmitRun}}
  Run: func(cmd *cobra.Command, args []string) {
    {{ template "args-parse" .CMD.Args }}

    {{ if .CMD.Body}}
    {{ .CMD.Body}}
    {{ else }}

    // Default body
    {{ if .CMD.Parent.Parent.Parent }}
    fmt.Println("{{ .CLI.Name }} {{ .CMD.Parent.Parent.Name }} {{ .CMD.Parent.Parent.Name }} {{ .CMD.Parent.Name }} {{ .CMD.Name }}"{{- range $i, $C := .CMD.Args }}, {{ .Name }}{{ end }})
    {{ else if .CMD.Parent.Parent }}
    fmt.Println("{{ .CLI.Name }} {{ .CMD.Parent.Parent.Name }} {{ .CMD.Parent.Name }} {{ .CMD.Name }}"{{- range $i, $C := .CMD.Args }}, {{ .Name }}{{ end }})
    {{ else if .CMD.Parent }}
    fmt.Println("{{ .CLI.Name }} {{ .CMD.Parent.Name }} {{ .CMD.Name }}"{{- range $i, $C := .CMD.Args }}, {{ .Name }}{{ end }})
    {{ else }}
    fmt.Println("{{ .CLI.Name }} {{ .CMD.Name }}"{{- range $i, $C := .CMD.Args }}, {{ .Name }}{{ end }})
    {{ end }}

    {{ end }}
  },
  {{ end }}

  {{ if .CMD.PersistentPostrun}}
  PersistentPostRun: func(cmd *cobra.Command, args []string) {
    {{ template "args-parse" .CMD.Args }}

    {{ if .CMD.PersistentPostrunBody}}
    {{ .CMD.PersistentPostrunBody}}
    {{ end }}
  },
  {{ end }}

  {{ if .CMD.Postrun}}
  PostRun: func(cmd *cobra.Command, args []string) {
    {{ template "args-parse" .CMD.Args }}

    {{ if .CMD.PostrunBody }}
    {{ .CMD.PostrunBody }}
    {{ end }}
  },
  {{ end }}
}

{{if .CMD.Commands}}
func init() {
	{{- range $i, $C := .CMD.Commands }}
  {{ $.CMD.CmdName }}Cmd.AddCommand({{ $.CMD.cmdName }}.{{ $C.CmdName }}Cmd)
	{{- end}}
}
{{ end }}

"""
