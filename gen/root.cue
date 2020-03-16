package gen

import (
  "text/template"

  "github.com/hof-lang/cuemod--cli-golang/schema"
  "github.com/hof-lang/cuemod--cli-golang/templates"
)

RootGen :: RootGenOpen & {}

RootGenOpen : {
  In: {
    CLI: schema.Cli
  }
  Template: templates.RootTemplate
  Filename: "commands/root.go"
  Out: template.Execute(Template, In)
}

