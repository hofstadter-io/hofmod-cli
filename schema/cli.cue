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

  Updates?: {
    _Releases: Releases // To require releases when updates are enabled?
    CheckURL: string
    DevCheckURL: string
    AvailableVersions: [...string]

    CheckingDisabled: bool | *false
  }

  Telemetry?: string
  TelemetryIdDir: string | *".\(cliName)"

  // Debugging
  EnablePProf: bool | *false

  #Common

  ...
}
