package schema

import (
  "strings"
)

Cli :: {
  Name:     string
  cliName:  strings.ToCamel(Name)
  CliName:  strings.ToTitle(Name)

  Package:  string

  VersionCommand: bool | *true
  BashCompletion: bool | *true

  Releases?: GoReleaser

  // Debugging
  EnablePProf: bool | *false

  Common

  ...
}
