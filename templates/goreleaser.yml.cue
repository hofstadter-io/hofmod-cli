package templates

GoReleaserTemplate :: """
build:
  main: main.go
  binary: {% .CLI.cliName %}
  ldflags: -s -w -X main.builddate={{.Date}}
  env:
    - CGO_ENABLED=0
  goos:
    - darwin
    - linux
    - windows
  goarch:
    - amd64
    - arm64

archive:
  format: tar.gz
  format_overrides:
    - goos: windows
      format: zip
  name_template: "{{.Binary}}_{{.Version}}_{{.Os}}-{{.Arch}}"
  replacements:
    amd64: 64bit
    arm: ARM
    arm64: ARM64
    darwin: MacOS
    linux: Linux
    windows: Windows
  files:
    - README.md
    - LICENSE

brew:
  name: {% .CLI.cliName %}
  github:
    owner: {% .CLI.GoReleaser.Brew.GitHubOwner %}
    name: {%  .CLI.GoReleaser.Brew.GitHubRepoName %}
  folder: Formula
  commit_author:
    name: {% .CLI.GoReleaser.Brew.GitHubUsername %}
    email: {% .CLI.GoReleaser.Brew.GitHubEmail %}
  homepage: {% .CLI.GoReleaser.Homepage %}
  description: '{% .CLI.Short %}'

release:
  draft: {% .CLI.GoReleaser.Draft %}

"""
