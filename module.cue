package cli

import (
	"list"

  hof "github.com/hofstadter-io/hof/schema"

  "github.com/hofstadter-io/hofmod-cli/schema"
  "github.com/hofstadter-io/hofmod-cli/templates"
)

HofGenerator :: hof.HofGenerator & {
  Cli: schema.Cli
  Outdir?: string

  // Internal
  In: {
    CLI: Cli
  }

  PackageName: "github.com/hofstadter-io/hofmod-cli"

  PartialsDir:  "/partials/"
  TemplatesDir: "/templates/"

  // Combine everything together and output files that might need to be generated
  _All: [
    _OnceFiles,
    _Commands,
    _SubCommands,
    _SubSubCommands,
    _CommandsLib,
    _SubCommandsLib,
    _SubSubCommandsLib,
  ]
  Out: [...hof.HofGeneratorFile] & list.FlattenN(_All , 1)

  // Files that are not repeatedly used, they are generated once for the whole CLI
  _OnceFiles: [...hof.HofGeneratorFile] & [
    {
      TemplateName: "main.go"
      Filepath: "main.go"
    },
    {
      TemplateName: "root.go"
      Filepath: "commands/root.go"
    },
    {
      if In.CLI.VersionCommand != _|_ {
        TemplateName: "version.go"
        Filepath: "commands/version.go"
      }
    },
    {
      if In.CLI.BashCompletion != _|_ {
        TemplateName: "bash-completions.go"
        Filepath: "commands/bash-completion.go"
      }
    },

    {
      if In.CLI.Releases != _|_ {
        TemplateName:  "goreleaser.yml"
        Filepath:  ".goreleaser.yml"
        TemplateConfig: {
          AltDelims: true
          LHS2_D: "{%"
          RHS2_D: "%}"
          LHS3_D: "{%%"
          RHS3_D: "%%}"
        }
      }
    },

    {
      if In.CLI.Releases != _|_ {
        Template:  templates.DockerfileJessie
        Filepath:  "ci/docker/Dockerfile.jessie"
      }
    },
    {
      if In.CLI.Releases != _|_ {
        Template:  templates.DockerfileScratch
        Filepath:  "ci/docker/Dockerfile.scratch"
      }
    },

  ]

  // Sub command tree
  _Commands: [...hof.HofGeneratorFile] & [ // List comprehension
    {
      In: {
        // CLI
        CMD: {
          C
          PackageName: "commands"
        }
      }
      TemplateName: "cmd.go"
      Filepath: "commands/\(In.CMD.Name).go"
    }
    for _, C in Cli.Commands
  ]

  _SubCommands: [...hof.HofGeneratorFile] & [ // List comprehension
    {
      In: {
        CMD: C
      }
      TemplateName: "cmd.go"
      Filepath: "commands/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"
    }
    for _, C in list.FlattenN([[{ C,  Parent: { Name: P.In.CMD.Name } } for _, C in P.In.CMD.Commands ] for _, P in _Commands if P.In.CMD.Commands != _|_ ], 1)
  ]

  _SubSubCommands: [...hof.HofGeneratorFile] & [ // List comprehension
    {
      In: {
        CMD: C
      }
      TemplateName: "cmd.go"
      Filepath: "commands/\(In.CMD.Parent.Parent.Name)/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"
    }
    for _, C in list.FlattenN([[{ C,  Parent: { Name: P.In.CMD.Name, Parent: P.In.CMD.Parent } } for _, C in P.In.CMD.Commands ] for _, P in _SubCommands if P.In.CMD.Commands != _|_ ], 1)
  ]

  // SubSubSubCommand
  // Filepath: "commands/\(In.CMD.Parent.Parent.Parent.Name)/\(In.CMD.Parent.Parent.Name)/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"

  _CommandsLib: [...hof.HofGeneratorFile] & [ // List comprehension
    {
      In: {
        CMD: {
          C
          PackageName: "commands"
        }
      }
      TemplateName: "cmdlib.go"
      Filepath: "lib/commands/\(In.CMD.Name).go"
    }
    for _, C in Cli.Commands
  ]

  _SubCommandsLib: [...hof.HofGeneratorFile] & [ // List comprehension
    {
      In: {
        CMD: C
      }
      TemplateName: "cmdlib.go"
      Filepath: "lib/commands/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"
    }
    for _, C in list.FlattenN([[{ C,  Parent: { Name: P.In.CMD.Name } } for _, C in P.In.CMD.Commands ] for _, P in _CommandsLib if P.In.CMD.Commands != _|_ ], 1)
  ]

  _SubSubCommandsLib: [...hof.HofGeneratorFile] & [ // List comprehension
    {
      In: {
        CMD: C
      }
      TemplateName: "cmdlib.go"
      Filepath: "lib/commands/\(In.CMD.Parent.Parent.Name)/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"
    }
    for _, C in list.FlattenN([[{ C,  Parent: { Name: P.In.CMD.Name, Parent: P.In.CMD.Parent } } for _, C in P.In.CMD.Commands ] for _, P in _SubCommandsLib if P.In.CMD.Commands != _|_ ], 1)
  ]

}
