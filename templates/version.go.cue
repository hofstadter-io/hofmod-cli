package templates

VersionCommandTemplate :: """
package commands

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	Version   = "dev"
	Commit    = "dirty"
	BuildDate = "unknown"
)

var VersionLong = `Print the build version for {{ .CLI.cliName }}`

var VersionCmd = &cobra.Command{

	Use: "version",

	Aliases: []string{
		"ver",
	},

	Short: "print the version",

	Long: VersionLong,

	Run: func(cmd *cobra.Command, args []string) {
    fmt.Printf("Version:    %v\\nCommit:     %v\\nBuildDate:  %v\\n", Version, Commit, BuildDate)
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}
"""

