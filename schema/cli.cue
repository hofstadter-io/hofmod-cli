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

  VersionCommand: bool | *true
  CompletionCommands: bool | *true

  Releases?: #GoReleaser

  ConfigDir: string | *"\(cliName)"

  Updates: bool | *true

  Telemetry?: string
  TelemetryIdDir: string | *".\(cliName)"

  // Debugging
  EnablePProf: bool | *false

  #Common

  ...
}
