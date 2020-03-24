package gen

import (
  "github.com/hofstadter-io/cuelib/template"

  "github.com/hofstadter-io/cuemod--cli-golang/schema"
  "github.com/hofstadter-io/cuemod--cli-golang/templates"
)

MainGen : {
  In: {
    CLI: schema.Cli
  }
  Template: templates.MainTemplate
  Filename: "main.go"
  Out: (template.RenderTemplate & { Template: templates.MainTemplate, Values: In}).Out
}

ToolGen : {
  In: {
    CLI: schema.Cli
  }
  Template: templates.ToolTemplate
  Filename: "cli_tool.cue"
  Out: (template.RenderTemplate & { Template: templates.ToolTemplate, Values: In}).Out
}

GoReleaserGen : {
  In: {
    CLI: schema.Cli
  }
  Template: templates.GoReleaserTemplate
  Filename: ".goreleaser.yml"
  Out: (template.AltDelimTemplate & { Template: templates.GoReleaserTemplate, Values: In}).Out
}

