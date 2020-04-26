package schema


#GoReleaser: {
  Disable: bool | *true
  Draft: bool | *true
  Author: string
  Homepage: string

  GitHub: {
    Owner: string
    Repo: string
  }

  Docker: {
    Maintainer: string
    Repo: string
  }

  ...
}
