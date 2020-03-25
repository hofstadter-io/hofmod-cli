package schema


GoReleaser :: {
  Draft: bool | *true
  Author: string
  Homepage: string

  Brew: {
    GitHubOwner: string
    GitHubRepoName: string
    GitHubUsername: string
    GitHubEmail: string
  }

  ...
}
