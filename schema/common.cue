package schema

common : {
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

  Pflags?:   [...Flag]
  Flags?:    [...Flag]
  Args?:     [...Arg]
  Commands:  [...Command] | *[]
}

