package schema

#Import: {
  As: string | *""
  Path: string
}

#Common: {
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

  HasAnyRun: bool
  HasAnyRun: PersistentPrerun || Prerun || !OmitRun || Postrun || PersistentPostrun

  HasAnyFlags: bool
  HasAnyFlags: Pflags != _|_ || Flags != _|_

  Imports?:  [...#Import]
  Pflags?:   [...#Flag]
  Flags?:    [...#Flag]
  Args?:     [...#Arg]
  Commands:  [...#Command] | *[...]

  ...
}

