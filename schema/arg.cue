package schema

import (
  "strings"
)

#ArgType:
  "string"  |
  "int"

#Arg: {
  Name:     string
  argName:  strings.ToCamel(Name)
  ArgName:  strings.ToTitle(Name)

  Type:       #ArgType
  Default?:   _
  Required?:  bool
  Rest?:      bool
  Help:       string | *""

  ...
}

