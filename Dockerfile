# Build Stage
FROM lacion/alpine-golang-buildimage:1.12.4 AS build-stage

LABEL app="build-inventory"
LABEL REPO="https://github.com/triardn/inventory"

ENV PROJPATH=/go/src/github.com/triardn/inventory

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/triardn/inventory
WORKDIR /go/src/github.com/triardn/inventory

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/triardn/inventory"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/inventory/bin

WORKDIR /opt/inventory/bin

COPY --from=build-stage /go/src/github.com/triardn/inventory/bin/inventory /opt/inventory/bin/
RUN chmod +x /opt/inventory/bin/inventory

# Create appuser
RUN adduser -D -g '' inventory
USER inventory

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/inventory/bin/inventory"]
