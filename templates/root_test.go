package cmd_test

import (
	"testing"

	"github.com/hofstadter-io/hof/script/runtime"
	"github.com/hofstadter-io/hof/lib/yagu"

	"{{ .CLI.Package }}/cmd"
)

func init() {
	// ensure our root command is setup
	cmd.RootInit()
}

func TestScriptRootCliTests(t *testing.T) {
	// setup some directories
	workdir := ".workdir/cli/root"
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
		Dir: "hls/cli/root",
		WorkdirRoot: workdir,
	})
}

