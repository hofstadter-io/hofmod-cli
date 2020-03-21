package gen

import (
  "text/template"

  "github.com/hofstadter-io/cuemod--cli-golang/schema"
  "github.com/hofstadter-io/cuemod--cli-golang/templates"
)

RootGen : {
  In: {
    CLI: schema.Cli
  }
  Template: templates.RootTemplate
  Filename: "commands/root.go"
  Out: template.Execute(Template, In)
}

