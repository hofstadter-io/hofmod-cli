package cli

import (
	"list"

  "github.com/hofstadter-io/cuemod--cli-golang/gen"
  "github.com/hofstadter-io/cuemod--cli-golang/schema"
)

Schema : schema.Cli

Generator : {
  Cli: schema.Cli

  // Files that are not repeatedly used, they are generated once for the whole CLI
  _OnceIn: {
    In: {
      CLI: Cli
    }
  }
  _OnceFiles: [ G & _OnceIn for _, G in gen.OnceFiles ]

  // Sub command tree
  _Commands: [ // List comprehension
    {
      gen.CommandGen & {
        In: {
          CLI: Cli
          CMD: C & {
            PackageName: "commands"
          }
        }
      },
    }
    for _, C in Cli.Commands
  ]

  _SubCmds:  [[C & { Parent: P.In.CMD } for _, C in P.In.CMD.Commands] for _, P in _Commands]
  _SubCommands: [ // List comprehension
    {
      gen.CommandGen & {
        In: {
          CLI: Cli
          CMD: C
        }
      },
    }
    for _, C in list.FlattenN( _SubCmds, 1)
  ]

  _SubSubCmds:  [[C & { Parent: P.In.CMD } for _, C in P.In.CMD.Commands] for _, P in _SubCommands]
  _SubSubCommands: [ // List comprehension
    {
      gen.CommandGen & {
        In: {
          CLI: Cli
          CMD: C
        }
      },
    }
    for _, C in list.FlattenN( _SubSubCmds, 1)
  ]

  // Combine everything together and output files that might need to be generated
  _All: [_OnceFiles, _Commands, _SubCommands, _SubSubCommands]
  Out: list.FlattenN(_All , 1)
}
