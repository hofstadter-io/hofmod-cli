package gen

import (
  "github.com/hofstadter-io/cuemod--cli-golang/schema"
  "github.com/hofstadter-io/cuemod--cli-golang/templates"
)

CommandGen :: {
  In: {
    CLI: schema.Cli
    CMD: schema.Command
  }
  Template: templates.CommandTemplate
  if In.CMD.Parent == _|_ {
    Filename: "commands/\(In.CMD.Name).go"
  } 
  if In.CMD.Parent != _|_ {
    if In.CMD.Parent.Parent == _|_ {
      Filename: "commands/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"
    }
    if In.CMD.Parent.Parent != _|_ {
      if In.CMD.Parent.Parent.Parent == _|_ {
        Filename: "commands/\(In.CMD.Parent.Parent.Name)/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"
      }
      if In.CMD.Parent.Parent.Parent != _|_ {
        Filename: "commands/\(In.CMD.Parent.Parent.Parent.Name)/\(In.CMD.Parent.Parent.Name)/\(In.CMD.Parent.Name)/\(In.CMD.Name).go"
      }
    }
  }

  ...
}

