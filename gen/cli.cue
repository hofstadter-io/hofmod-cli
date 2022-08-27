package gen

import (
	"list"

	hof "github.com/hofstadter-io/hof/schema/gen"

	"github.com/hofstadter-io/hofmod-cli/schema"
	"github.com/hofstadter-io/hofmod-cli/templates"
)

#HofGenerator: #Generator
#Generator: hof.#Generator & {
	Cli:     schema.#Cli
	Outdir?: string | *"./"

	OutdirConfig: {
		CiOutdir:    string | *"ci/\(In.CLI.cliName)"
		CliOutdir:   string | *"cmd/\(In.CLI.cliName)"
		CmdOutdir:   string | *"cmd/\(In.CLI.cliName)/cmd"
		FlagsOutdir: string | *"cmd/\(In.CLI.cliName)/flags"
	}

	// Internal
	In: {
		CLI: Cli
	}

	basedir: "cmd/\(In.CLI.cliName)"

	PackageName: "github.com/hofstadter-io/hofmod-cli"

	Templates: [{
		Globs: ["templates/*.*"]
		TrimPrefix: "templates/"
	}, {
		Globs: ["templates/alt/*"]
		TrimPrefix: "templates/"
		Delims: {
			LHS: "{%"
			RHS: "%}"
		}
	}]

	// Combine everything together and output files that might need to be generated
	All: [
		for _, F in OnceFiles {F},
		for _, F in S1_Cmds {F},
		for _, F in S2_Cmds {F},
		for _, F in S3_Cmds {F},
		for _, F in S4_Cmds {F},
		for _, F in S5_Cmds {F},
		for _, F in S1_Flags {F},
		for _, F in S2_Flags {F},
		for _, F in S3_Flags {F},
		for _, F in S4_Flags {F},
		for _, F in S5_Flags {F},
	]

	// Out: OnceFiles
	// Out: [...hof.#HofGeneratorFile] & OnceFiles
	// Out: list.FlattenN(All , 1)
	Out: [...hof.#HofGeneratorFile] & All

	// Files that are not repeatedly used, they are generated once for the whole CLI
	OnceFiles: [...hof.#HofGeneratorFile] & [
			{
			TemplatePath: "main.go"
			Filepath:     "\(OutdirConfig.CliOutdir)/main.go"
		},
		{
			TemplatePath: "root.go"
			Filepath:     "\(OutdirConfig.CmdOutdir)/root.go"
		},
		{
			TemplatePath: "flags.go"
			Filepath:     "\(OutdirConfig.FlagsOutdir)/root.go"
		},
		{
			if (In.CLI.VersionCommand & true) != _|_ {
				TemplatePath: "version.go"
				Filepath:     "\(OutdirConfig.CmdOutdir)/version.go"
			}
		},
		{
			if (In.CLI.VersionCommand & true) != _|_ {
				TemplatePath: "verinfo.go"
				Filepath:     "\(OutdirConfig.CliOutdir)/verinfo/verinfo.go"
			}
		},
		{
			if (In.CLI.Updates & true) != _|_ {
				TemplatePath: "update.go"
				Filepath:     "\(OutdirConfig.CmdOutdir)/update.go"
			}
		},
		{
			if (In.CLI.CompletionCommands & true) != _|_ {
				TemplatePath: "completions.go"
				Filepath:     "\(OutdirConfig.CmdOutdir)/completions.go"
			}
		},
		{
			if In.CLI.Telemetry != _|_ {
				TemplatePath: "ga.go"
				Filepath:     "\(OutdirConfig.CliOutdir)/ga/ga.go"
			}
		},
		{
			if In.CLI.Releases != _|_ {
				TemplateContent: templates.DockerfileDebian
				Filepath:        "\(OutdirConfig.CiOutdir)/docker/Dockerfile.debian"
			}
		},
		{
			if In.CLI.Releases != _|_ {
				TemplateContent: templates.DockerfileAlpine
				Filepath:        "\(OutdirConfig.CiOutdir)/docker/Dockerfile.alpine"
			}
		},
		{
			if In.CLI.Releases != _|_ {
				TemplateContent: templates.DockerfileScratch
				Filepath:        "\(OutdirConfig.CiOutdir)/docker/Dockerfile.scratch"
			}
		},
		{
			if In.CLI.Releases != _|_ {
				TemplatePath: "alt/goreleaser.yml"
				Filepath:     "\(OutdirConfig.CliOutdir)/.goreleaser.yml"
			}
		},

	]

	// Sub command tree
	S1_Cmds: [...hof.#HofGeneratorFile] & [
			for _, C in Cli.Commands {
			In: {
				CMD: {
					C
					PackageName: "cmd"
				}
			}
			TemplatePath: "cmd.go"
			Filepath:     "\(OutdirConfig.CmdOutdir)/\(In.CMD.Name).go"
		},
	]

	S2C: list.FlattenN([ for P in S1_Cmds if len(P.In.CMD.Commands) > 0 {
		[ for C in P.In.CMD.Commands {C, Parent: {Name: P.In.CMD.Name}}]
	}], 1)
	S2_Cmds: [...hof.#HofGeneratorFile] & [ // List comprehension
			for _, C in S2C {
			In: {
				CMD: C
			}
			TemplatePath: "cmd.go"
			Filepath:     "\(OutdirConfig.CmdOutdir)/\(C.Parent.Name)/\(C.Name).go"
		},
	]

	S3C: list.FlattenN([ for P in S2_Cmds if len(P.In.CMD.Commands) > 0 {
		[ for C in P.In.CMD.Commands {C, Parent: {Name: P.In.CMD.Name, Parent: P.In.CMD.Parent}}]
	}], 1)
	S3_Cmds: [...hof.#HofGeneratorFile] & [ // List comprehension
			for _, C in S3C {
			In: {
				CMD: C
			}
			TemplatePath: "cmd.go"
			Filepath:     "\(OutdirConfig.CmdOutdir)/\(C.Parent.Parent.Name)/\(C.Parent.Name)/\(C.Name).go"
		},
	]

	S4C: list.FlattenN([ for P in S3_Cmds if len(P.In.CMD.Commands) > 0 {
		[ for C in P.In.CMD.Commands {C, Parent: {Name: P.In.CMD.Name, Parent: P.In.CMD.Parent}}]
	}], 1)
	S4_Cmds: [...hof.#HofGeneratorFile] & [ // List comprehension
			for _, C in S4C {
			In: {
				CMD: C
			}
			TemplatePath: "cmd.go"
			Filepath:     "\(OutdirConfig.CmdOutdir)/\(C.Parent.Parent.Parent.Name)/\(C.Parent.Parent.Name)/\(C.Parent.Name)/\(C.Name).go"
		},
	]

	S5C: list.FlattenN([ for P in S4_Cmds if len(P.In.CMD.Commands) > 0 {
		[ for C in P.In.CMD.Commands {C, Parent: {Name: P.In.CMD.Name, Parent: P.In.CMD.Parent}}]
	}], 1)
	S5_Cmds: [...hof.#HofGeneratorFile] & [ // List comprehension
			for _, C in S5C {
			In: {
				CMD: C
			}
			TemplatePath: "cmd.go"
			Filepath:     "\(OutdirConfig.CmdOutdir)/\(C.Parent.Parent.Parent.Parent.Name)/\(C.Parent.Parent.Parent.Name)/\(C.Parent.Parent.Name)/\(C.Parent.Name)/\(C.Name).go"
		},
	]

	// Persistent Flags
	S1_Flags: [...hof.#HofGeneratorFile] & [ // List comprehension
			for _, C in Cli.Commands if C.Pflags != _|_ || C.Flags != _|_ {
			In: {
				// CLI
				CMD: {
					C
					PackageName: "flags"
				}
			}
			TemplatePath: "flags.go"
			Filepath:     "\(OutdirConfig.FlagsOutdir)/\(In.CMD.Name).go"
		},
	]

	S2F: list.FlattenN([ for P in S1_Flags if len(P.In.CMD.Commands) > 0 {
		[ for C in P.In.CMD.Commands if C.Pflags != _|_ || C.Flags != _|_ {C, Parent: {Name: P.In.CMD.Name}}]
	}], 1)
	S2_Flags: [...hof.#HofGeneratorFile] & [ // List comprehension
			for _, C in S2F {
			In: {
				CMD: {
					C
					PackageName: "flags"
				}
			}
			TemplatePath: "flags.go"
			Filepath:     "\(OutdirConfig.FlagsOutdir)/\(C.Parent.Name)__\(C.Name).go"
		},
	]

	S3F: list.FlattenN([ for P in S2_Flags if len(P.In.CMD.Commands) > 0 {
		[ for C in P.In.CMD.Commands if C.Pflags != _|_ || C.Flags != _|_ {C, Parent: {Name: P.In.CMD.Name, Parent: P.In.CMD.Parent}}]
	}], 1)
	S3_Flags: [...hof.#HofGeneratorFile] & [ // List comprehension
			for _, C in S3F {
			In: {
				CMD: {
					C
					PackageName: "flags"
				}
			}
			TemplatePath: "flags.go"
			Filepath:     "\(OutdirConfig.FlagsOutdir)/\(C.Parent.Parent.Name)__\(C.Parent.Name)__\(C.Name).go"
		},
	]

	S4F: list.FlattenN([ for P in S3_Flags if len(P.In.CMD.Commands) > 0 {
		[ for C in P.In.CMD.Commands if C.Pflags != _|_ || C.Flags != _|_ {C, Parent: {Name: P.In.CMD.Name, Parent: P.In.CMD.Parent}}]
	}], 1)
	S4_Flags: [...hof.#HofGeneratorFile] & [ // List comprehension
			for _, C in S4F {
			In: {
				CMD: {
					C
					PackageName: "flags"
				}
			}
			TemplatePath: "flags.go"
			Filepath:     "\(OutdirConfig.FlagsOutdir)/\(C.Parent.Parent.Parent.Name)__\(C.Parent.Parent.Name)__\(C.Parent.Name)__\(C.Name).go"
		},
	]

	S5F: list.FlattenN([ for P in S4_Flags if len(P.In.CMD.Commands) > 0 {
		[ for C in P.In.CMD.Commands if C.Pflags != _|_ || C.Flags != _|_ {C, Parent: {Name: P.In.CMD.Name, Parent: P.In.CMD.Parent}}]
	}], 1)
	S5_Flags: [...hof.#HofGeneratorFile] & [ // List comprehension
			for _, C in S5F {
			In: {
				CMD: {
					C
					PackageName: "flags"
				}
			}
			TemplatePath: "flags.go"
			Filepath:     "\(OutdirConfig.FlagsOutdir)/\(C.Parent.Parent.Parent.Parent.Name)__\(C.Parent.Parent.Parent.Name)__\(C.Parent.Parent.Name)__\(C.Parent.Name)__\(C.Name).go"
		},
	]

	...
}
