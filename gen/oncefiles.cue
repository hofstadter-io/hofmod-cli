package gen

import (
  "github.com/hofstadter-io/cuemod--cli-golang/schema"
  "github.com/hofstadter-io/cuemod--cli-golang/templates"
)

OnceFiles :: [
  MainGen,
  RootGen,
  VersionGen,
  CompletionGen,
  ToolGen,
  ReleasesGen
]

MainGen :: {
  In: {
    CLI: schema.Cli | *{...}
  }
  Template: templates.MainTemplate
  Filename: "main.go"

  ...
}

RootGen :: {
  In: {
    CLI: schema.Cli | *{...}
  }
  Template: templates.RootTemplate
  Filename: "commands/root.go"

  ...
}

VersionGen :: {
  In: {
    CLI: schema.Cli | *{...}
  }
  if In.CLI.VersionCommand != _|_ {
    Template: templates.VersionCommandTemplate
    Filename: "commands/version.go"
  }

  ...
}

CompletionGen :: {
  In: {
    CLI: schema.Cli | *{...}
  }
  if In.CLI.BashCompletion != _|_ {
    Template: templates.BashCompletionTemplate
    Filename: "commands/bash-completion.go"
  }

  ...
}

ToolGen :: {
  In: {
    CLI: schema.Cli | *{...}
  }
  Template: templates.ToolTemplate
  Filename: "cue_tool.cue"

  ...
}

ReleasesGen :: {
  In: {
    CLI: schema.Cli | *{...}
  }
  if In.CLI.Releases != _|_ {
    Template: templates.ReleasesTemplate
    Filename: ".goreleaser.yml"
    Alt:      true
  }

  ...
}
