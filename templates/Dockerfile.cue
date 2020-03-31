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

COPY hof /usr/bin/local
ENTRYPOINT ["hof"]

\(DockerfileWorkdir)

"""


DockerfileScratch :: """
FROM scratch
\(DockerfileMaintainer)

COPY hof /
ENTRYPOINT ["/hof"]

\(DockerfileWorkdir)

"""

