package schema

import (
  "strings"
)

Cli :: common & {
  Name:     string
  cliName:  strings.ToCamel(Name)
  CliName:  strings.ToTitle(Name)

  Package:  string
}
