package schema

import (
  "strings"
)

#FlagType:
  "string"  | "[]string"  |
  "int"     | "[]int"     |
  "float64" | "[]float64" |
  "bool"

#Flag: {
  Name:      string
  flagName:  strings.ToCamel(Name)
  FlagName:  strings.ToTitle(Name)

  Type:     #FlagType
  Default:  _
  Help:     string | *""
  Long:     string | *""
  Short:    string | *""

  ...
}

