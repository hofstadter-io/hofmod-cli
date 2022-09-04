package cli

import (
	// "github.com/hofstadter-io/hof/schema/common"
	"github.com/hofstadter-io/hof/schema/gen"
)

Creator: gen.#Generator & {
	@gen(creator)

	Outdir: "./"

	CreateMessage: {
	  let name = CreateInput.name
		Before: "Creating a new Go Cli"
		After: """
		Your new Cli generator is ready, run the following
		to generate the code, build the binary, and run \(name).

		0. fill in any empty fields
		1. hof mod vendor cue
		2. hof gen
		3. go build -o \(name) ./cmd/\(name)
		4. ./\(name) or ./(name) -h
		"""
	}

	CreateInput: _
	CreatePrompt: [{
		Name:       "name"
		Type:       "input"
		Prompt:     "What is your CLI named"
		Required:   true
		// Validation: common.NameLabel
	},{
		Name:       "repo"
		Type:       "input"
		Prompt:     "Git repository"
		Default:    "github.com/user/repo"
		// Validation: common.NameLabel
	},{
		Name:       "about"
		Type:       "input"
		Prompt:     "Tell us a bit about it..."
		Required:   true
		// Validation: common.NameLabel
	},{
		Name:       "updates"
		Type:       "confirm"
		Prompt:     "Enable self updating"
		Default:    true
	},{
		Name:       "telemetry"
		Type:       "confirm"
		Prompt:     "Enable telemetry"
	},{
		Name:       "releases"
		Type:       "confirm"
		Prompt:     "Enable GoReleaser tooling"
		Default:    true
	}]

	In: {
		CreateInput
		...
	}

	Out: [...gen.#File] & [ {
		TemplatePath: "gen.cue"
		Filepath:     "gen.cue"
	}, {
		TemplatePath: "cli.cue"
		Filepath:     "cli.cue"
	}, {
		TemplatePath: "debug"
		Filepath:     "debug.yaml"
	}]

	gen.#EmptyTemplates

	EmbeddedTemplates: {
		debug: {
			Content: """
			{{ yaml . }}
			"""
		}
		"gen.cue": Content: #GenTemplate
		"cli.cue": Content: #CliTemplate
	}
}

#GenTemplate: """
package {{ .name }}

import (
	"github.com/hofstadter-io/hofmod-cli/gen"
)

"{{ .name }}": gen.#Generator & {
	@gen({{.name}},cli)
	Outdir: "./"
	"Cli": Cli
	WatchGlobs: ["./*.cue"]
}

"""

#CliTemplate: """
package {{ .name }}

import (
	"github.com/hofstadter-io/hofmod-cli/schema"
)

Cli: schema.#Cli & {

	Name:    "{{ .name }}"
	Package: "{{ .repo }}/cmd/{{ .name }}"

	Usage:      "{{ .name }}"
	Short:      "{{ .about }}"
	Long:       Short

	// set to true to print help from root command
	// and always assume subcommands are to be run
	OmitRun: false

	Args: [
		// add any args here
	]

	Commands: [
		// add any subcommands here
	]

	// extras
	VersionCommand: true
	CompletionCommands: true
	{{ if .updates }}Updates: true{{ end }}

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

"""
