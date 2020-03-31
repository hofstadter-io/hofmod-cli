package templates

ReleasesTemplate :: GoReleaserTemplate

GoReleaserTemplate :: """
project_name: "{% .CLI.cliName %}"

before:
  hooks:
    - go mod vendor

builds:
- binary: "{% .CLI.cliName %}"
  main: main.go

  ldflags:
    - -s -w
    - -X {% .CLI.Package %}/commands.Version={{.Version}}
    - -X {% .CLI.Package %}/commands.Commit={{.FullCommit}}
    - -X {% .CLI.Package %}/commands.BuildDate={{.Date}}
    - -X {% .CLI.Package %}/commands.GoVersion=go1.14
    - -X {% .CLI.Package %}/commands.BuildOS={{.Os}}
    - -X {% .CLI.Package %}/commands.BuildArch={{.Arch}}
    - -X {% .CLI.Package %}/commands.BuildArm={{.Arm}}

  env:
  - CGO_ENABLED=0

  goos:
    - darwin
    - linux
    - windows
  goarch:
    - amd64
    - arm64
    - arm

  ignore:
  - goos: linux
    goarch: arm


snapshot:
  name_template: "{{ .Tag }}-SNAPSHOT-{{.ShortCommit}}"

archives:
- format: binary
  replacements:
    darwin: Darwin
    linux: Linux
    windows: Windows
    amd64: x86_64
  # Needed hack for binary only uploads
  # For more information, check #602
  files:
   - "thisfiledoesnotexist*"


checksum:
  name_template: '{{ .ProjectName }}_{{ .Version }}_checksums.txt'

changelog:
  sort: asc
  filters:
    exclude:
    - '^docs:'
    - '^test:'


release:
  disable: {% .CLI.Releases.Disable %}
  draft: {% .CLI.Releases.Draft %}
  github:
    owner: {% .CLI.Releases.GitHub.Owner %}
    name: {%  .CLI.Releases.GitHub.Repo %}

dockers:
- binaries:
    - {% .CLI.cliName %}
  skip_push: true
  dockerfile: ci/docker/Dockerfile.jessie
  image_templates:
  - "{% .CLI.Releases.Docker.Repo %}/{{.ProjectName}}:{{.Tag}}"
  - "{% .CLI.Releases.Docker.Repo %}/{{.ProjectName}}:v{{ .Major }}.{{ .Minor }}"
  - "{% .CLI.Releases.Docker.Repo %}/{{.ProjectName}}:v{{ .Major }}"
  - "{% .CLI.Releases.Docker.Repo %}/{{.ProjectName}}:latest"

  - "{% .CLI.Releases.Docker.Repo %}/{{.ProjectName}}:{{.Tag}}-debian"
  - "{% .CLI.Releases.Docker.Repo %}/{{.ProjectName}}:v{{ .Major }}.{{ .Minor }}-debian"
  - "{% .CLI.Releases.Docker.Repo %}/{{.ProjectName}}:v{{ .Major }}-debian"
  - "{% .CLI.Releases.Docker.Repo %}/{{.ProjectName}}:latest-debian"


- binaries:
    - {% .CLI.cliName %}
  skip_push: true
  dockerfile: ci/docker/Dockerfile.scratch
  image_templates:
  - "{% .CLI.Releases.Docker.Repo %}/{{.ProjectName}}:{{.Tag}}-scrath"
  - "{% .CLI.Releases.Docker.Repo %}/{{.ProjectName}}:v{{ .Major }}.{{ .Minor }}-scrath"
  - "{% .CLI.Releases.Docker.Repo %}/{{.ProjectName}}:v{{ .Major }}-scrath"
  - "{% .CLI.Releases.Docker.Repo %}/{{.ProjectName}}:latest-scrath"

"""
