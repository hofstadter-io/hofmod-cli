package schema

import (
  "strings"
)

#Cli: {
  Name:     string
  cliName:  strings.ToCamel(Name)
  CliName:  strings.ToTitle(Name)
  CLI_NAME: strings.ToUpper(Name)

	Module:   string
  Package:  string | *"\(Module)/cmd/\(Name)"

  Updates: bool | *true
  VersionCommand: bool | *true
  CompletionCommands: bool | *true

  Releases?: #GoReleaser

  ConfigDir: string | *"\(cliName)"

	// directory of files to embed into the binary
	EmbedDir?: string

  Telemetry?: string
	TelemetryAsk: string | *"""
	We only send the command run, no args or input.
	You can disable at any time by setting
	  \(CLI_NAME)_TELEMETRY_DISABLED=1

	Would you like to help by sharing very basic usage stats?
	"""
	// subdir under os.UserConfigDir
  TelemetryIdDir: string | *"\(cliName)"

  // Debugging
  EnablePProf: bool | *false

  #Common
}
