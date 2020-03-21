package gen

import (
  "text/template"

  "github.com/hofstadter-io/cuemod--cli-golang/schema"
  "github.com/hofstadter-io/cuemod--cli-golang/templates"
)

MainGen : {
  In: {
    CLI: schema.Cli
  }
  Template: templates.MainTemplate
  Filename: "main.go"
  Out: template.Execute(Template, In)
}

