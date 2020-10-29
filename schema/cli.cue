package schema

import (
  "strings"
)

#Cli: {
  Name:     string
  cliName:  strings.ToCamel(Name)
  CliName:  strings.ToTitle(Name)
  CLI_NAME: strings.ToUpper(Name)

  Package:  string

  Updates: bool | *true
  VersionCommand: bool | *true
  CompletionCommands: bool | *true

  Releases?: #GoReleaser

  ConfigDir: string | *"\(cliName)"

	// directory of files to embed into the binary
	EmbedDir?: string

  Telemetry?: string
  TelemetryIdDir: string | *".\(cliName)"

  // Debugging
  EnablePProf: bool | *false

  #Common
}
