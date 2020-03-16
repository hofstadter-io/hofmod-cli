package gen

import (
  "text/template"

  "github.com/hof-lang/cuemod--cli-golang/schema"
  "github.com/hof-lang/cuemod--cli-golang/templates"
)

MainGen : MainGenOpen & {}

MainGenOpen : {
  In: {
    CLI: schema.Cli
  }
  Template: templates.MainTemplate
  Filename: "main.go"
  Out: template.Execute(Template, In)
}

