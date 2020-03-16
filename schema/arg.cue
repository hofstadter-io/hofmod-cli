package schema

import (
  "strings"
)

ArgType ::
  "string"  |
  "int"

Arg :: ArgOpen & {}

ArgOpen : {
  Name:     string
  argName:  strings.ToCamel(Name)
  ArgName:  strings.ToTitle(Name)

  Type:       ArgType
  Default?:   _
  Required?:  bool
  Rest?:      bool
  Help:       string
}

