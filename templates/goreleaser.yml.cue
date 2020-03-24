package templates

ReleasesTemplate :: GoReleaserTemplate

GoReleaserTemplate :: """
build:
  main: main.go
  binary: {% .CLI.cliName %}
  ldflags: -s -w -X {% .CLI.Package %}/commands.Version={{.Version}} -X {% .CLI.Package %}/commands.Commit={{.ShortCommit}} -X {% .CLI.Package %}/commands.BuildDate={{.Date}}
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
    owner: {% .CLI.Releases.Brew.GitHubOwner %}
    name: {%  .CLI.Releases.Brew.GitHubRepoName %}
  folder: Formula
  commit_author:
    name: {% .CLI.Releases.Brew.GitHubUsername %}
    email: {% .CLI.Releases.Brew.GitHubEmail %}
  homepage: {% .CLI.Releases.Homepage %}
  description: '{% .CLI.Short %}'

release:
  draft: {% .CLI.Releases.Draft %}

"""
