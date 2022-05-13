package templates


DockerfileMaintainer: """
MAINTAINER {{ .CLI.Releases.Docker.Maintainer }}
"""

DockerfileWorkdir: """
VOLUME ["/work"]
WORKDIR /work
"""


DockerfileDebian: """
FROM debian:{{ .CLI.Releases.Docker.Versions.debian }}
\(DockerfileMaintainer)

COPY {{ .CLI.cliName }} /usr/local/bin
ENTRYPOINT ["{{ .CLI.cliName }}"]

\(DockerfileWorkdir)

"""

DockerfileAlpine: """
FROM alpine:{{ .CLI.Releases.Docker.Versions.alpine }}
\(DockerfileMaintainer)

COPY {{ .CLI.cliName }} /usr/local/bin
ENTRYPOINT ["{{ .CLI.cliName }}"]

\(DockerfileWorkdir)

"""


DockerfileScratch: """
FROM scratch
\(DockerfileMaintainer)

COPY {{ .CLI.cliName }} /
ENTRYPOINT ["/{{ .CLI.cliName }}"]

\(DockerfileWorkdir)

"""

