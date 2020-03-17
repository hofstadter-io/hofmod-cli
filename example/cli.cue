package cli

import (
  "github.com/hofstadter-io/cuemod--cli-golang:cli"
  "github.com/hofstadter-io/cuemod--cli-golang/schema"
)

Outdir: "./example"

GEN : cli.Generator & {
  Cli: CLI
}

CLI : cli.Schema & {
  Name: "example"
  Package: "github.com/hof-lang/cuemod--cli-golang/example"

  Uage: "example"
  Short: "Example CLI with hof-lang/cuemod--cli-golang"
  Long:  Short

  OmitRun: true

  Pflags: [
    schema.Flag & {
      Name: "config"
      Type: "string"
      Default: ""
      Help: "Some config file path"
      Long: "config"
      Short: "c"
    }
  ]

  PersistentPrerun: true
  PersistentPrerunBody: """
    fmt.Println("PersistentPrerun", RootConfigPflag, args)
  """

  Commands: [
    {
      Name: "hello"
      Usage:  "hello <name>"
      Short:  "world?"
      Long:   Short

      Args: [
        schema.Arg & {
          Name: "name"
          Type: "string"
          Required: true
          Help: "name to say hi to"
        }
      ]

      Body: """
        fmt.Println("hello", name)
      """
    },
    {
      Name:   "echo"
      Usage:  "echo [words]"
      Short:  "echo a message."
      Long:   "echo a message, no need for \"\" around the words"

      Body: """
        fmt.Println(strings.Join(args))
      """
    },
    {
      Name: "config"
      Usage:  "config"
      Short:  "get/set configuration values"
      Long:   Short

      OmitRun: true

      Commands: [
        {
          Name: "init"
          Usage:  "init"
          Short:  "init configuration file"
          Long:   Short
        },
        {
          Name: "get"
          Usage:  "get <key>"
          Short:  "get configuration value"
          Long:   Short
        },
        {
          Name: "set"
          Usage:  "set"
          Short:  "set a configuration value"
          Long:   Short
          OmitRun: true
          Commands: [
            {
              Name: "host"
              Usage:  "host"
              Short:  "set host value"
              Long:   Short
            },
            {
              Name: "port"
              Usage:  "port"
              Short:  "set port value"
              Long:   Short
            },
            {
              Name: "schema"
              Usage:  "schema"
              Short:  "set http schema"
              Long:   Short
            },
          ]
        },
      ]
    },
  ]
}
