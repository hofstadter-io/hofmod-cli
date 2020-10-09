package schema

import (
  "strings"
)

#Command: {
  Name:     string
  cmdName:  strings.ToCamel(Name)
  CmdName:  strings.ToTitle(Name)

  Hidden?:       bool
	OmitTests?:    bool
  Aliases?:      [...string]
  PackageName?:  string

  #Common
}

