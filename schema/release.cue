package schema


#GoReleaser: {
  Disable: bool | *true
  Draft: bool | *true
  Author: string
  Homepage: string

  GitHub: {
    Owner: string
    Repo: string
    URL: "https://github.com/\(Owner)/\(Repo)"
    Rel: "https://api.github.com/repos/\(Owner)/\(Repo)/releases"
  }

  Docker: {
    Maintainer: string
    Repo: string
    Versions: {
      debian: string | *"11.3"
      alpine: string | *"3.15"
    }
  }
}
