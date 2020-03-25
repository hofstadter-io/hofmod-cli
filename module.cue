package cli

import (
	"list"

  "github.com/hofstadter-io/cuemod--cli-golang/gen"
  "github.com/hofstadter-io/cuemod--cli-golang/schema"
)

Schema :: schema.Cli

Generator :: OptimizedGenerator

OptimizedGenerator :: {
  Cli: schema.Cli

  // Files that are not repeatedly used, they are generated once for the whole CLI
  OnceIn: {
    In: {
      CLI: Cli
    }

    ...
  }
  OnceFiles: [ G & OnceIn for _, G in gen.OnceFiles ]

  // Sub command tree
  Commands: [ // List comprehension
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

  SubCmds:  [[C & { Parent: P.In.CMD } for _, C in P.In.CMD.Commands] for _, P in Commands]
  SubCommands: [ // List comprehension
    {
      gen.CommandGen & {
        In: {
          CLI: Cli
          CMD: C
        }
      },
    }
    for _, C in list.FlattenN( SubCmds, 1)
  ]

  SubSubCmds:  [[C & { Parent: P.In.CMD } for _, C in P.In.CMD.Commands] for _, P in SubCommands]
  SubSubCommands: [ // List comprehension
    {
      gen.CommandGen & {
        In: {
          CLI: Cli
          CMD: C
        }
      },
    }
    for _, C in list.FlattenN( SubSubCmds, 1)
  ]

  // Combine everything together and output files that might need to be generated
  All: [OnceFiles, Commands, SubCommands, SubSubCommands]
  Out: list.FlattenN(All , 1)
}


OrigGenerator :: {
  Cli: schema.Cli

  // Files that are not repeatedly used, they are generated once for the whole CLI
  OnceIn: {
    In: {
      CLI: Cli
    }

    ...
  }
  OnceFiles: [ G & OnceIn for _, G in gen.OnceFiles ]

  // Sub command tree
  Commands: [ // List comprehension
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

  SubCmds:  [[C & { Parent: P.In.CMD } for _, C in P.In.CMD.Commands] for _, P in Commands]
  SubCommands: [ // List comprehension
    {
      gen.CommandGen & {
        In: {
          CLI: Cli
          CMD: C
        }
      },
    }
    for _, C in list.FlattenN( SubCmds, 1)
  ]

  SubSubCmds:  [[C & { Parent: P.In.CMD } for _, C in P.In.CMD.Commands] for _, P in SubCommands]
  SubSubCommands: [ // List comprehension
    {
      gen.CommandGen & {
        In: {
          CLI: Cli
          CMD: C
        }
      },
    }
    for _, C in list.FlattenN( SubSubCmds, 1)
  ]

  // Combine everything together and output files that might need to be generated
  All: [OnceFiles, Commands, SubCommands, SubSubCommands]
  Out: list.FlattenN(All , 1)
}
