package creators

import (
	"github.com/hofstadter-io/hof/schema/common"
	"github.com/hofstadter-io/hof/schema/gen"
)

Creator: gen.#Generator & {
	@gen(creator)

	Create: {
		Message: {
			let name = Input.name
			Before: "Creating a new Go Cli"
			After: """
			Your new Cli generator is ready, run the following
			to generate the code, build the binary, and run \(name).

			now run 'make first'    (cd to the --outdir if used)
			"""
		}

		Args: [...string]
		if len(Args) > 0 {
			Input: name: Args[0]
		}

		Input: {
			name:      string
			repo:      string
			about:     string
			releases:  bool | *false
			updates:   bool | *false
			telemetry: bool | *false
		}

		Prompt: [{
			Name:       "name"
			Type:       "input"
			Prompt:     "What is your CLI named"
			Required:   true
			Validation: common.NameLabel
		},{
			Name:       "repo"
			Type:       "input"
			Prompt:     "Git repository"
			Default:    "github.com/user/repo"
			Validation: common.NameLabel
		},{
			Name:       "about"
			Type:       "input"
			Prompt:     "Tell us a bit about it..."
			Required:   true
			Validation: common.NameLabel
		},{
			Name:       "releases"
			Type:       "confirm"
			Prompt:     "Enable GoReleaser tooling"
			Default:    true
		},

		if Input.releases == true {
			Name:       "updates"
			Type:       "confirm"
			Prompt:     "Enable self updating"
			Default:    true
		}

		if Input.releases == true {
			Name:       "telemetry"
			Type:       "confirm"
			Prompt:     "Enable telemetry"
		}
		]
	}

	In: {
		Create.Input
		...
	}

	Out: [...gen.#File] & [ 
		for file in [
			// "debug.yaml",
			"cue.mod/module.cue",
			"cli.cue",
			"Makefile",
		]{ TemplatePath: file, Filepath: file }
	]

	Templates: [{
		Globs: ["creators/templates/**/*"]
		TrimPrefix: "creators/templates/"
	}]
	Statics: []
	Partials: []

	EmbeddedTemplates: {
		"debug.yaml": {
			Content: """
			{{ yaml . }}
			"""
		}
	}
}
