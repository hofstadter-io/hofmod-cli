package cli

import (
	"list"

  "github.com/hofstadter-io/cuemod--cli-golang/gen"
  "github.com/hofstadter-io/cuemod--cli-golang/schema"
  "github.com/hofstadter-io/cuemod--cli-golang/templates"
)

Schema :: schema.Cli

// Generator :: OptimizedGenerator
Generator :: OrigGenerator

OrigGenerator :: {
  Cli: schema.Cli

  // Files that are not repeatedly used, they are generated once for the whole CLI
  // Sub command tree
  OnceFiles: [ // List comprehension
    {
      In: {
        CLI: Cli
      }
      Template: templates.MainTemplate
      Filename: "main.go"
    },
    {
      In: {
        CLI: Cli
      }
      Template: templates.RootTemplate
      Filename: "commands/root.go"
    },
    {
      In: {
        CLI: Cli
      }
      if In.CLI.VersionCommand != _|_ {
        Template: templates.VersionCommandTemplate
        Filename: "commands/version.go"
      }
    },
    {
      In: {
        CLI: Cli
      }
      if In.CLI.BashCompletion != _|_ {
        Template: templates.BashCompletionTemplate
        Filename: "commands/bash-completion.go"
      }
    },
    {
      In: {
        CLI: Cli
      }
      Template: templates.ToolTemplate
      Filename: "cue_tool.cue"
    },
    {
      In: {
        CLI: Cli
      }
      if In.CLI.Releases != _|_ {
        Template: templates.ReleasesTemplate
        Filename: ".goreleaser.yml"
        Alt:      true
      }
    },

  ]

  // Sub command tree
  Commands: [ // List comprehension
    {
      In: {
        CLI: Cli
        CMD: C & {
          PackageName: "commands"
        }
      }
      Template: templates.CommandTemplate
      Filename: "commands/\(In.CMD.Name).go"
    }
    for _, C in Cli.Commands
  ]

  SubCommands: [ // List comprehension
    {
      In: {
        CLI: Cli
        CMD: C
      }
      Template: templates.CommandTemplate
      Filename: "commands/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"
    }
    for _, C in list.FlattenN([[C & { Parent: P.In.CMD } for _, C in P.In.CMD.Commands] for _, P in Commands], 1)
  ]

  SubSubCommands: [ // List comprehension
    {
      In: {
        CLI: Cli
        CMD: C
      }
      Template: templates.CommandTemplate
      Filename: "commands/\(In.CMD.Parent.Parent.Name)/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"
    }
    for _, C in list.FlattenN([[C & { Parent: P.In.CMD } for _, C in P.In.CMD.Commands] for _, P in SubCommands], 1)
  ]

  // SubSubSubCommand
  // Filename: "commands/\(In.CMD.Parent.Parent.Parent.Name)/\(In.CMD.Parent.Parent.Name)/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"

  // Combine everything together and output files that might need to be generated
  All: [OnceFiles, Commands, SubCommands, SubSubCommands]
  Out: list.FlattenN(All , 1)
}

OptimizedGenerator :: {
  Cli: schema.Cli

  // Files which only need to be generated once
  //   see gen/oncefiles.cue
  OnceFiles: gen.OnceFiles

  // Files for commmands like turules
  Commands:        list.FlattenN([[c & { Parent: C } for _, c in C.Commands] for _, C in Cli.Commands], -1)
  SubCommands:     list.FlattenN([[c & { Parent: C } for _, c in C.Commands] for _, C in Commands], -1)
  SubSubCommands:  list.FlattenN([[c & { Parent: C } for _, c in C.Commands] for _, C in SubCommands], -1)

  // Combine everything together and output files that might need to be generated
  All: [OnceFiles, Commands, SubCommands, SubSubCommands]
  Out: list.FlattenN(All , 1)
}


