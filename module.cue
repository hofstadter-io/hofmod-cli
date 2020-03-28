package cli

import (
	"list"

  "github.com/hofstadter-io/hofmod-cli/schema"
  "github.com/hofstadter-io/hofmod-cli/templates"
)

HofGenerator :: {
  Cli: schema.Cli
  Outdir?: string

  In: {
    CLI: Cli
  }

  // Files that are not repeatedly used, they are generated once for the whole CLI
  _OnceFiles: [
    {
      Template: templates.MainTemplate
      Filepath: "main.go"
    },
    {
      Template: templates.RootTemplate
      Filepath: "commands/root.go"
    },
    {
      if In.CLI.VersionCommand != _|_ {
        Template: templates.VersionCommandTemplate
        Filepath: "commands/version.go"
      }
    },
    {
      if In.CLI.BashCompletion != _|_ {
        Template: templates.BashCompletionTemplate
        Filepath: "commands/bash-completion.go"
      }
    },
    {
      if In.CLI.Releases != _|_ {
        Template: templates.ReleasesTemplate
        Filepath: ".goreleaser.yml"
        Alt:      true
      }
    },

  ]

  // Sub command tree
  _Commands: [ // List comprehension
    {
      In: {
        CMD: {
          C
          PackageName: "commands"
        }
      }
      Template: templates.CommandTemplate
      Filepath: "commands/\(In.CMD.Name).go"
    }
    for _, C in Cli.Commands
  ]

  _SubCommands: [ // List comprehension
    {
      In: {
        CMD: C
      }
      Template: templates.CommandTemplate
      Filepath: "commands/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"
    }
    for _, C in list.FlattenN([[{ C,  Parent: { Name: P.In.CMD.Name } } for _, C in P.In.CMD.Commands ] for _, P in _Commands if P.In.CMD.Commands != _|_ ], 1)
  ]

  _SubSubCommands: [ // List comprehension
    {
      In: {
        CMD: C
      }
      Template: templates.CommandTemplate
      Filepath: "commands/\(In.CMD.Parent.Parent.Name)/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"
    }
    for _, C in list.FlattenN([[{ C,  Parent: { Name: P.In.CMD.Name, Parent: P.In.CMD.Parent } } for _, C in P.In.CMD.Commands ] for _, P in _SubCommands if P.In.CMD.Commands != _|_ ], 1)
  ]

  // SubSubSubCommand
  // Filepath: "commands/\(In.CMD.Parent.Parent.Parent.Name)/\(In.CMD.Parent.Parent.Name)/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"

  // Combine everything together and output files that might need to be generated
  Out: list.FlattenN([_OnceFiles, _Commands, _SubCommands, _SubSubCommands] , 1)
}


CueGenerator :: {
  Cli: Schema

  In: {
    CLI: Cli
  }

  // Files that are not repeatedly used, they are generated once for the whole CLI
  _OnceFiles: [
    {
      In: {
        CLI: Cli
      }
      Template: templates.MainTemplate
      Filepath: "main.go"
    },
    {
      In: {
        CLI: Cli
      }
      Template: templates.RootTemplate
      Filepath: "commands/root.go"
    },
    {
      In: {
        CLI: Cli
      }
      if In.CLI.VersionCommand != _|_ {
        Template: templates.VersionCommandTemplate
        Filepath: "commands/version.go"
      }
    },
    {
      In: {
        CLI: Cli
      }
      if In.CLI.BashCompletion != _|_ {
        Template: templates.BashCompletionTemplate
        Filepath: "commands/bash-completion.go"
      }
    },
    {
      In: {
        CLI: Cli
      }
      if In.CLI.Releases != _|_ {
        Template: templates.ReleasesTemplate
        Filepath: ".goreleaser.yml"
        Alt:      true
      }
    },

  ]

  // Sub command tree
  _Commands: [ // List comprehension
    {
      In: {
        CLI: Cli
        CMD: {
          C
          PackageName: "commands"
        }
      }
      Template: templates.CommandTemplate
      Filepath: "commands/\(In.CMD.Name).go"
    }
    for _, C in Cli.Commands
  ]

  _SubCommands: [ // List comprehension
    {
      In: {
        CLI: Cli
        CMD: C
      }
      Template: templates.CommandTemplate
      Filepath: "commands/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"
    }
    for _, C in list.FlattenN([[{ C,  Parent: { Name: P.In.CMD.Name } } for _, C in P.In.CMD.Commands] for _, P in _Commands], 1)
  ]

  _SubSubCommands: [ // List comprehension
    {
      In: {
        CLI: Cli
        CMD: C
      }
      Template: templates.CommandTemplate
      Filepath: "commands/\(In.CMD.Parent.Parent.Name)/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"
    }
    for _, C in list.FlattenN([[{ C,  Parent: { Name: P.In.CMD.Name, Parent: P.In.CMD.Parent } } for _, C in P.In.CMD.Commands] for _, P in _SubCommands], 1)
  ]

  // SubSubSubCommand
  // Filepath: "commands/\(In.CMD.Parent.Parent.Parent.Name)/\(In.CMD.Parent.Parent.Name)/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"

  // Combine everything together and output files that might need to be generated
  Out: list.FlattenN([_OnceFiles, _Commands, _SubCommands, _SubSubCommands] , 1)
}
