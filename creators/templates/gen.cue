package {{ .name }}

import (
	"github.com/hofstadter-io/hofmod-cli/gen"
)

"{{ .name }}": gen.#Generator & {
	@gen({{.name}},cli)
	Name: "{{ .name }}"
	Outdir: "./"
	"Cli": Cli
	WatchGlobs: ["./*.cue"]
}
