# hofmod-cli

A [hof](https://github.com/hofstadter-io/hof) generator for creating advanced Golang CLIs.

Design your CLI structure, arguments, flags, and a whole host of addons
and then generate the implementation. Changed your mind about what your
CLI should look like? Redesign, regenerate, and keep adding you custom code.

### Features:

- Quickly architect your CLI commands, arguments, flags, configuration
- Built on the fantastic [spf13/cobra](https://github.com/ andspf13/cobra)library for Golang CLIs
- Cross-platform builds and releases using GoReleaser, GitHub, and Docker
- Supports config files in local project falling back to OS specific application dir
- Your CLI will self check for updates and can self install with a user command
- Shell auto completion for bash, fish, zsh, and power shell
- Advanced help system with support for custom overviews, extra topics, and examples
- Telemetry systems which can hook up to Google Analytics
- Golang pprof and many other ENV VARs to control inner behavior

### Sites to see:

- [Schema](./schema) - the design spec your write a CLI in
- [Generator](./gen) - [hof](https://github.com/hofstadter-io/hof) generator definition you invoke
- [Templates](./templates) and [partials](./partials) - files which implement the code
- [Example](https://github.com/hofstadter-io/hof) - the [hof](https://github.com/hofstadter-io/hof) tool leverages and powers this, see the `hof.cue` and `design` directory for relevant files

### Usage

You'll need the [hof](https://github.com/hofstadter-io/hof) tool installed.
You can download `hof` from [the releases page](https://github.com/hofstadter-io/hof/releases).

Let's start a new project:

```
# Start a project
hof init github.com/verdverm/my-cli
cd my-cli
```

Add the following to the `cue.mods` file (same format as `go.mod`)

```
module github.com/verdverm/my-cli

cue v0.2.0

require (
    github.com/hofstadter-io/hofmod-cli v0.5.8
)
```

To fetch the module, run:

```
hof mod vendor cue

# and after the next file, run
hof gen
```

Create a file named `cli.cue` and add the following content:

```
package cli

import (
	"github.com/hofstadter-io/hofmod-cli/schema"

	"github.com/hofstadter-io/hof/design/cli/cmds"
)

# Typically we put the cli under a nested directory
#Outdir: "./cmd/hof"

#CLI: schema.#Cli & {
    # Name and package path (matches outdir)
	Name:    "hof"
	Package: "github.com/hofstadter-io/hof/cmd/hof"

    # Usage and help
	Usage:      "hof"
	Short:      "Polyglot Development Tool and Framework"
	Long:       Short
	CustomHelp: #RootCustomHelp

    # Print the help when no subcommands are supplied
	OmitRun: true

    # Command stage hooks
	PersistentPrerun:     true
    # You can write code here or...
	PersistentPrerunBody: "runtime.Init()"
 
	PersistentPostrun: true
    # ...or add custom code right in the output

    # Persistent flags work for all subcommands too
    Pflags: [{
		Name:    "labels"
		Long:    "label"
		Short:   "l"
		Type:    "[]string"
		Default: "nil"
		Help:    "Labels for use across all commands"
	}, {
		Name:    "config"
		Long:    "config"
		Short:   ""
		Type:    "string"
		Default: ""
		Help:    "Path to a hof configuration file"
	}, ...]

    # Subcommands and nested down as far as you need
	Commands: [

	]

	//
	// Addons
	//
	Releases: #CliReleases
	Updates:  true

    ...
}
```

(this was adapted from the [hof](https://github.com/hofstadter-io/hof t) tool)

Now run `hof gen` to generate your code.
Try adding implementation, and then build:

```
go build -o my-cli cmd/my-cli/main.go

./my-cli
```

Update your designs, rerun `hof gen`, rebuild
and keep iterating away!

