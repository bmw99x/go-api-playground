# Start from a Debian-based image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container's workspace.
WORKDIR /go/src/app
COPY . .

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go mod tidy
RUN go build -o app ./cmd
## Run the command by default when the container starts.
CMD ["./app"]

EXPOSE 8080