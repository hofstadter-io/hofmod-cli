package templates

VersionCommandTemplate :: """
package commands

import (
	"os"
  "time"

	"github.com/spf13/cobra"
)

var (
	Version   = "dev"
	Commit    = "dirty"
  BuiltBy   = os.Getenv("USER")
	BuildDate = time.Now().String()
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
    fmt.Printf("Version:    %v\\nCommit:     %v\\nBuiltBy:    %v\\nBuildDate:  %v\\n", Version, Commit, BuiltBy, BuildDate)
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}
"""

