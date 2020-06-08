package gen

import (
	"list"

  hof "github.com/hofstadter-io/hof/schema"

  "github.com/hofstadter-io/hofmod-cli/schema"
  "github.com/hofstadter-io/hofmod-cli/templates"
)

#HofGenerator: hof.#HofGenerator & {
  Cli: schema.#Cli
  Outdir?: string

  OutdirConfig: {
    CiOutdir: string | *"\(Outdir)/ci/\(In.CLI.cliName)"
    CliOutdir: string | *"\(Outdir)/cmd/\(In.CLI.cliName)"
    CmdOutdir: string | *"\(Outdir)/cmd/\(In.CLI.cliName)/cmd"
    FlagsOutdir: string | *"\(Outdir)/cmd/\(In.CLI.cliName)/flags"
  }

  // Internal
  In: {
    CLI: Cli
  }

  basedir: "cmd/\(In.CLI.cliName)"

  PackageName: "github.com/hofstadter-io/hofmod-cli"

  PartialsDir:  "./partials/"
  TemplatesDir: "./templates/"
  TemplatesDirConfig: {
    "**/goreleaser.yml": {
      AltDelims: true
      LHS2_D: "{%"
      RHS2_D: "%}"
      LHS3_D: "{%%"
      RHS3_D: "%%}"
    }
  }

  // Combine everything together and output files that might need to be generated
  _All: [
    _OnceFiles,
    _S1_Cmds,
    _S2_Cmds,
    _S3_Cmds,
    _S4_Cmds,
    _S5_Cmds,
    _S1_Flags,
    _S2_Flags,
    _S3_Flags,
    _S4_Flags,
    _S5_Flags,
  ]

  Out: [...hof.#HofGeneratorFile] & list.FlattenN(_All , 1)

  // Files that are not repeatedly used, they are generated once for the whole CLI
  _OnceFiles: [...hof.#HofGeneratorFile] & [
    {
      TemplateName: "main.go"
      Filepath: "\(OutdirConfig.CliOutdir)/main.go"
    },
    {
      TemplateName: "root.go"
      Filepath: "\(OutdirConfig.CmdOutdir)/root.go"
    },
    {
      TemplateName: "root_test.go"
      Filepath: "\(OutdirConfig.CmdOutdir)/root_test.go"
    },
    {
      TemplateName: "hls/cli/root_help.hls"
      Filepath: "\(OutdirConfig.CmdOutdir)/hls/cli/root/help.hls"
    },
    {
      TemplateName: "flags.go"
      Filepath: "\(OutdirConfig.FlagsOutdir)/root.go"
    },
    {
      if In.CLI.VersionCommand != _|_ {
        TemplateName: "version.go"
        Filepath: "\(OutdirConfig.CmdOutdir)/version.go"
      }
    },
    {
      if In.CLI.VersionCommand != _|_ {
        TemplateName: "verinfo.go"
        Filepath: "\(OutdirConfig.CliOutdir)/verinfo/verinfo.go"
      }
    },
    {
      if In.CLI.Updates != _|_ {
        TemplateName: "update.go"
        Filepath: "\(OutdirConfig.CmdOutdir)/update.go"
      }
    },
    {
      if In.CLI.CompletionCommands != _|_ {
        TemplateName: "completions.go"
        Filepath: "\(OutdirConfig.CmdOutdir)/completions.go"
      }
    },
    {
      if In.CLI.Telemetry != _|_ {
        TemplateName: "ga.go"
        Filepath: "\(OutdirConfig.CliOutdir)/ga/ga.go"
      }
    },
    {
      if In.CLI.Releases != _|_ {
        ( hof.#HofGeneratorFile & {
          TemplateName:  "goreleaser.yml"
          Filepath:  "\(OutdirConfig.CliOutdir)/.goreleaser.yml"
        })
      }
    },
    {
      if In.CLI.Releases != _|_ {
        Template:  templates.DockerfileJessie
        Filepath:  "\(OutdirConfig.CiOutdir)/docker/Dockerfile.jessie"
      }
    },
    {
      if In.CLI.Releases != _|_ {
        Template:  templates.DockerfileScratch
        Filepath:  "\(OutdirConfig.CiOutdir)/docker/Dockerfile.scratch"
      }
    },

  ]

  // Sub command tree
  _S1_Cmds: [...hof.#HofGeneratorFile] & list.FlattenN([[
    for _, C in Cli.Commands
    {
      In: {
        CMD: {
          C
          PackageName: "cmd"
        }
      }
      TemplateName: "cmd.go"
      Filepath: "\(OutdirConfig.CmdOutdir)/\(In.CMD.Name).go"
		}
	], [
    for _, C in Cli.Commands if C.OmitTests == _|_
	  {
      In: {
        CMD: {
          C
          PackageName: "cmd"
        }
      }
      TemplateName: "cmd_test.go"
      Filepath: "\(OutdirConfig.CmdOutdir)/\(In.CMD.Name)_test.go"
		}
	], [
    for _, C in Cli.Commands if C.OmitTests == _|_
	  {
      In: {
        CMD: C
      }
      TemplateName: "hls/cli/cmd_help.hls"
      Filepath: "\(OutdirConfig.CmdOutdir)/hls/cli/\(In.CMD.cmdName)/help.hls"
		}
	]], 1)

  _S2C: [ for P in _S1_Cmds if len(P.In.CMD.Commands) > 0 {
    [ for C in P.In.CMD.Commands { C,  Parent: { Name: P.In.CMD.Name } }]
  }]
  _S2_Cmds: [...hof.#HofGeneratorFile] & [ // List comprehension
    for _, C in list.FlattenN(_S2C, 1)
    {
      In: {
        CMD: C
      }
      TemplateName: "cmd.go"
      Filepath: "\(OutdirConfig.CmdOutdir)/\(C.Parent.Name)/\(C.Name).go"
    }
  ]

  _S3C: [ for P in _S2_Cmds if len(P.In.CMD.Commands) > 0 {
    [ for C in P.In.CMD.Commands { C,  Parent: { Name: P.In.CMD.Name, Parent: P.In.CMD.Parent } }]
  }]
  _S3_Cmds: [...hof.#HofGeneratorFile] & [ // List comprehension
    for _, C in list.FlattenN(_S3C, 1)
    {
      In: {
        CMD: C
      }
      TemplateName: "cmd.go"
      Filepath: "\(OutdirConfig.CmdOutdir)/\(C.Parent.Parent.Name)/\(C.Parent.Name)/\(C.Name).go"
    }
  ]

  _S4C: [ for P in _S3_Cmds if len(P.In.CMD.Commands) > 0 {
    [ for C in P.In.CMD.Commands { C,  Parent: { Name: P.In.CMD.Name, Parent: P.In.CMD.Parent } }]
  }]
  _S4_Cmds: [...hof.#HofGeneratorFile] & [ // List comprehension
    for _, C in list.FlattenN(_S4C, 1)
    {
      In: {
        CMD: C
      }
      TemplateName: "cmd.go"
      Filepath: "\(OutdirConfig.CmdOutdir)/\(C.Parent.Parent.Parent.Name)/\(C.Parent.Parent.Name)/\(C.Parent.Name)/\(C.Name).go"
    }
  ]

  _S5C: [ for P in _S4_Cmds if len(P.In.CMD.Commands) > 0 {
    [ for C in P.In.CMD.Commands { C,  Parent: { Name: P.In.CMD.Name, Parent: P.In.CMD.Parent } }]
  }]
  _S5_Cmds: [...hof.#HofGeneratorFile] & [ // List comprehension
    for _, C in list.FlattenN(_S5C, 1)
    {
      In: {
        CMD: C
      }
      TemplateName: "cmd.go"
      Filepath: "\(OutdirConfig.CmdOutdir)/\(C.Parent.Parent.Parent.Parent.Name)/\(C.Parent.Parent.Parent.Name)/\(C.Parent.Parent.Name)/\(C.Parent.Name)/\(C.Name).go"
    }
  ]


  // Persistent Flags
  _S1_Flags: [...hof.#HofGeneratorFile] & [ // List comprehension
    for _, C in Cli.Commands if C.Pflags != _|_ || C.Flags != _|_
    {
      In: {
        // CLI
        CMD: {
          C
          PackageName: "flags"
        }
      }
      TemplateName: "flags.go"
      Filepath: "\(OutdirConfig.FlagsOutdir)/\(In.CMD.Name).go"
    }
  ]

  _S2F: [ for P in _S1_Flags if len(P.In.CMD.Commands) > 0 {
    [ for C in P.In.CMD.Commands if C.Pflags != _|_ || C.Flags != _|_ { C,  Parent: { Name: P.In.CMD.Name } }]
  }]
  _S2_Flags: [...hof.#HofGeneratorFile] & [ // List comprehension
    for _, C in list.FlattenN(_S2F, 1)
    {
      In: {
        CMD: {
          C
          PackageName: "flags"
        }
      }
      TemplateName: "flags.go"
      Filepath: "\(OutdirConfig.FlagsOutdir)/\(C.Parent.Name)__\(C.Name).go"
    }
  ]

  _S3F: [ for P in _S2_Flags if len(P.In.CMD.Commands) > 0 {
    [ for C in P.In.CMD.Commands if C.Pflags != _|_ || C.Flags != _|_ { C,  Parent: { Name: P.In.CMD.Name, Parent: P.In.CMD.Parent } }]
  }]
  _S3_Flags: [...hof.#HofGeneratorFile] & [ // List comprehension
    for _, C in list.FlattenN(_S3F, 1)
    {
      In: {
        CMD: {
          C
          PackageName: "flags"
        }
      }
      TemplateName: "flags.go"
      Filepath: "\(OutdirConfig.FlagsOutdir)/\(C.Parent.Parent.Name)__\(C.Parent.Name)__\(C.Name).go"
    }
  ]

  _S4F: [ for P in _S3_Flags if len(P.In.CMD.Commands) > 0 {
    [ for C in P.In.CMD.Commands if C.Pflags != _|_ || C.Flags != _|_ { C,  Parent: { Name: P.In.CMD.Name, Parent: P.In.CMD.Parent } }]
  }]
  _S4_Flags: [...hof.#HofGeneratorFile] & [ // List comprehension
    for _, C in list.FlattenN(_S4F, 1)
    {
      In: {
        CMD: {
          C
          PackageName: "flags"
        }
      }
      TemplateName: "flags.go"
      Filepath: "\(OutdirConfig.FlagsOutdir)/\(C.Parent.Parent.Parent.Name)__\(C.Parent.Parent.Name)__\(C.Parent.Name)__\(C.Name).go"
    }
  ]

  _S5F: [ for P in _S4_Flags if len(P.In.CMD.Commands) > 0 {
    [ for C in P.In.CMD.Commands if C.Pflags != _|_ || C.Flags != _|_ { C,  Parent: { Name: P.In.CMD.Name, Parent: P.In.CMD.Parent } }]
  }]
  _S5_Flags: [...hof.#HofGeneratorFile] & [ // List comprehension
    for _, C in list.FlattenN(_S5F, 1)
    {
      In: {
        CMD: {
          C
          PackageName: "flags"
        }
      }
      TemplateName: "flags.go"
      Filepath: "\(OutdirConfig.FlagsOutdir)/\(C.Parent.Parent.Parent.Parent.Name)__\(C.Parent.Parent.Parent.Name)__\(C.Parent.Parent.Name)__\(C.Parent.Name)__\(C.Name).go"
    }
  ]
}

