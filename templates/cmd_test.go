{{ if .CMD.Parent }}
package cmd{{ .CMD.Parent.Name }}_test
{{ else }}
package cmd_test
{{ end }}

import (
	"testing"

	"github.com/hofstadter-io/hof/script/runtime"
	"github.com/hofstadter-io/hof/lib/yagu"

	{{ if .CMD.Commands }}

	{{ if .CMD.Parent.Parent.Parent.Parent }}
	"{{ .CLI.Package }}/cmd/{{ .CMD.Parent.Parent.Parent.Parent.Name }}/{{ .CMD.Parent.Parent.Parent.Name }}/{{ .CMD.Parent.Parent.Name }}/{{ .CMD.Parent.Name }}"
	{{ else if .CMD.Parent.Parent.Parent }}
	"{{ .CLI.Package }}/cmd/{{ .CMD.Parent.Parent.Parent.Name }}/{{ .CMD.Parent.Parent.Name }}/{{ .CMD.Parent.Name }}"
	{{ else if .CMD.Parent.Parent }}
	"{{ .CLI.Package }}/cmd/{{ .CMD.Parent.Parent.Name }}/{{ .CMD.Parent.Name }}"
	{{ else if .CMD.Parent }}
	"{{ .CLI.Package }}/cmd/{{ .CMD.Parent.Name }}"
	{{ else }}
	"{{ .CLI.Package }}/cmd"
	{{ end }}

	{{ else }}
	"{{ .CLI.Package }}/cmd"
	{{ end }}
)

func TestScript{{ .CMD.CmdName }}CliTests(t *testing.T) {
	// setup some directories
{{ if .CMD.Commands }}
  {{ if .CMD.Parent.Parent.Parent.Parent }}
	dir := "{{ .CMD.Parent.Parent.Parent.Parent.Name }}/{{ .CMD.Parent.Parent.Parent.Name }}/{{ .CMD.Parent.Parent.Name }}/{{ .CMD.Parent.Name }}/{{ .CMD.cmdName }}"
  {{ else if .CMD.Parent.Parent.Parent }}
	dir := "{{ .CMD.Parent.Parent.Parent.Name }}/{{ .CMD.Parent.Parent.Name }}/{{ .CMD.Parent.Name }}/{{ .CMD.cmdName }}"
  {{ else if .CMD.Parent.Parent }}
	dir := "{{ .CMD.Parent.Parent.Name }}/{{ .CMD.Parent.Name }}/{{ .CMD.cmdName }}"
  {{ else if .CMD.Parent }}
	dir := "{{ .CMD.Parent.Name }}/{{ .CMD.cmdName }}"
  {{ else }}
	dir := "{{ .CMD.cmdName }}"
  {{ end }}
{{ else }}
	dir := "{{ .CMD.cmdName }}"
{{ end }}
	workdir := ".workdir/cli/" + dir
	yagu.Mkdir(workdir)

	runtime.Run(t, runtime.Params{
		Setup: func (env *runtime.Env) error {
			// add any environment variables for your tests here
			{{ if .CLI.Telemetry }}
			env.Vars = append(env.Vars, "{{ .CLI.CLI_NAME }}_TELEMETRY_DISABLED=1")
			{{ end }}
			return nil
		},
		Funcs: map[string] func (ts* runtime.Script, args[]string) error {
			"__{{ .CLI.cliName }}": cmd.CallTS,
		},
		Dir: "hls/cli/{{.CMD.cmdName}}",
		WorkdirRoot: workdir,
	})
}

