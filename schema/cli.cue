package schema

import (
  "strings"
)

Cli :: CliOpen & Common {}

CliOpen : CommonOpen & {
  Name:     string
  cliName:  strings.ToCamel(Name)
  CliName:  strings.ToTitle(Name)

  Package:  string
}
