package cli

import (
  "github.com/hofstadter-io/dsl-cli:cli"
  "github.com/hofstadter-io/dsl-cli/schema"
)

GEN : cli.Generator & {
  Cli: CLI
}

CLI : cli.Schema & {
  Name: "example"
  Package: "github.com/hof-lang/cuemod--cli-golang/example"
  Short: "Example CLI with hof-lang/cuemod--cli-golang"
  OmitRun: true

  Commands: [
    schema.Command & {
      Name: "echo"
    },
    schema.Command & {
      Name: "hello"
    },
    schema.Command & {
      Name: "config"
      Commands: [
        schema.Command & {
          Name: "init"
        },
        schema.Command & {
          Name: "get"
        },
        schema.Command & {
          Name: "set"
        },
      ]
    },
  ]
}
