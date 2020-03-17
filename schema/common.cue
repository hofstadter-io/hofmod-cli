package schema

Import : {
  As: string | *""
  Path: string
}

Command : Common & {
  Usage:    string
  Short:    string
  Long:     string

  PersistentPrerun:   bool | *false
  Prerun:             bool | *false
  OmitRun:            bool | *false
  Postrun:            bool | *false
  PersistentPostrun:  bool | *false

  PersistentPrerunBody?:   string
  PrerunBody?:             string
  Body?:                   string
  PostrunBody?:            string
  PersistentPostrunBody?:  string

  Imports?:  [...ImportOpen]
  Pflags?:   [...FlagOpen]
  Flags?:    [...FlagOpen]
  Args?:     [...ArgOpen]
  Commands:  [...CommandOpen] | *[]
}

