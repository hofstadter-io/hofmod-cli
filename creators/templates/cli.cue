package {{ .name }}

import (
	"github.com/hofstadter-io/hofmod-cli/gen"
	"github.com/hofstadter-io/hofmod-cli/schema"
)

"{{ .name }}": gen.#Generator & {
	@gen({{.name}},cli)
	Name: "{{ .name }}"
	Outdir: "./"
	Cli: cli
	WatchGlobs: ["./*.cue"]
}

cli: schema.#Cli & {

	Name:    "{{ .name }}"
	Module:  "{{ .repo }}"
	Package: "{{ .repo }}/cmd/{{ .name }}"

	Usage:      "{{ .name }}"
	Short:      "{{ .about }}"
	Long:       Short

	// set to true to print help from root command
	// and always assume subcommands are to be run
	// set to false to run code from the root command
	OmitRun: true

	Args: [
		// add any args here
	]

	Commands: [
		// add any subcommands here
	]

	// extras
	VersionCommand: true
	CompletionCommands: true
	{{ if .updates }}
	Updates: true
	{{ else }}
	Updates: false
	{{ end }}

	{{ if .telemetry }}
	// set your GA identifier here
	Telemetry: "ua-xxxxxx"
	{{ end }}

	{{ if .releases }}
	// GoReleaser configuration
	// see https://docs.hofstadter.io/... for options
	Releases: {
		Author: ""
		Homepage: ""

		GitHub: {
			Owner: ""
			Repo:  ""
		}

		Docker: {
			Maintainer: ""
			Repo: ""
		}
	}
	{{ end }}
}

