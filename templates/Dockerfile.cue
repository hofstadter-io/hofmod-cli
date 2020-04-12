package templates


DockerfileMaintainer :: """
MAINTAINER {{ .CLI.Releases.Docker.Maintainer }}
"""

DockerfileWorkdir :: """
VOLUME ["/work"]
WORKDIR /work
"""


DockerfileJessie :: """
FROM debian:jessie
\(DockerfileMaintainer)

COPY {{ .CLI.cliName }} /usr/bin/local
ENTRYPOINT ["{{ .CLI.cliName }}"]

\(DockerfileWorkdir)

"""


DockerfileScratch :: """
FROM scratch
\(DockerfileMaintainer)

COPY {{ .CLI.cliName }} /
ENTRYPOINT ["/{{ .CLI.cliName }}"]

\(DockerfileWorkdir)

"""

