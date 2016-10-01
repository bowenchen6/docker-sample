# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

RUN go get -u github.com/garyburd/redigo/redis
RUN go get -u github.com/gorilla/mux
# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/bowen/server

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/bowen/server

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/server

# Document that the service listens on port 8080.
EXPOSE 8080