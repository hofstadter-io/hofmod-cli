package cmd

import (
  "fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"{{ .CLI.Package }}/ga"
	"{{ .CLI.Package }}/verinfo"
)

const versionMessage = `
Version:     v%s
Commit:      %s

BuildDate:   %s
GoVersion:   %s
OS / Arch:   %s %s

{{ with .CLI.Releases }}
Author:   {{ .Author }}
Homepage: {{ .Homepage }}
GitHub:   {{ .GitHub.URL }}
{{ end }}
`

var VersionLong = `Print the build version for {{ .CLI.cliName }}`

var VersionCmd = &cobra.Command{

	Use: "version",

	Aliases: []string{
		"ver",
	},

	Short: "print the version",

	Long: VersionLong,

	Run: func(cmd *cobra.Command, args []string) {
		{{ if .CLI.ConfigDir }}
		s, e := os.UserConfigDir()
		fmt.Printf("{{ .CLI.Name }} ConfigDir %q %v\n", filepath.Join(s, "{{ .CLI.ConfigDir }}"), e)
		{{ end }}

    fmt.Printf(
      versionMessage,
      verinfo.Version,
      verinfo.Commit,
      verinfo.BuildDate,
      verinfo.GoVersion,
      verinfo.BuildOS,
      verinfo.BuildArch,
    )
	},
}

func init() {
	help := VersionCmd.HelpFunc()
	usage := VersionCmd.UsageFunc()

	{{ if .CLI.Telemetry }}
	thelp := func (cmd *cobra.Command, args []string) {
		if VersionCmd.Name() == cmd.Name() {
			ga.SendCommandPath("version help")
		}
		help(cmd, args)
	}
	tusage := func (cmd *cobra.Command) error {
		if VersionCmd.Name() == cmd.Name() {
			ga.SendCommandPath("version usage")
		}
		return usage(cmd)
	}
	VersionCmd.SetHelpFunc(thelp)
	VersionCmd.SetUsageFunc(tusage)
	{{ else }}
	VersionCmd.SetHelpFunc(help)
	VersionCmd.SetUsageFunc(usage)
	{{ end }}
}
