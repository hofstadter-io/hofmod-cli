package cmd_test

import (
	"testing"

	"github.com/hofstadter-io/hof/lib/gotils/testscript"
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

	testscript.Run(t, testscript.Params{
		Setup: func (env *testscript.Env) error {
			// add any environment variables for your tests here
			{{ if .CLI.Telemetry }}
			env.Vars = append(env.Vars, "{{ .CLI.CLI_NAME }}_TELEMETRY_DISABLED=1")
			{{ end }}
			return nil
		},
		Funcs: map[string] func (ts* testscript.TestScript, args[]string) error {
			"__{{ .CLI.cliName }}": cmd.CallTS,
		},
		Dir: "testscripts/cli/root",
		WorkdirRoot: workdir,
	})
}

