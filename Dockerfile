# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/kahoona77/emerald
ADD ./emerald.conf /opt/emerald/

# Build the emerald command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go build -o /opt/emerald/emerald github.com/kahoona77/emerald

#copy assets
RUN cp -r /go/src/github.com/kahoona77/emerald/assets /opt/emerald/assets

# Run the emerald command by default when the container starts.
WORKDIR /opt/emerald
ENTRYPOINT ./emerald -conf /opt/emerald/emerald.conf

# Document that the service listens on port 8080.
EXPOSE 8080
