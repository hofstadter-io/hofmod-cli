package schema

import "text/template"

#Import: {
  As: string | *""
  Path: string
}

#Common: {
  Name:     string
  Usage:    string | *Name
  Short:    string 
  Long:     string | *Short
	TBD:      string | *""

	Help: template.Execute("{{ printf \"%-15s %-5s %s\" .name .tbd .short }}", { tbd: TBD, name: Name, short: Short })
	CustomHelp?: string

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
  HasAnyRun: !OmitRun || PersistentPrerun || Prerun || Postrun || PersistentPostrun

  HasAnyFlags: bool
  HasAnyFlags: Pflags != _|_ || Flags != _|_

  Imports?:  [...#Import]
  Pflags?:   [...#Flag]
  Flags?:    [...#Flag]
  Args?:     [...#Arg]
  Commands:  [...#Command] | *[]

	Topics:    [string]: string
	Examples:  [string]: string
	Tutorials: [string]: string

  ...
}

