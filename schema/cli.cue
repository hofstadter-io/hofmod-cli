package schema

import (
  "strings"
)

Cli : Common & {
  Name:     string
  cliName:  strings.ToCamel(Name)
  CliName:  strings.ToTitle(Name)

  Package:  string

  GoReleaser: GoReleaser
}
