# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.16

ENV GO111MODULE  on

# Create mh user and home folder
RUN groupadd --gid 9999 mhuser
RUN useradd --create-home --uid 9999 --gid 9999 --shell /bin/bash mhuser

# Create WORKDIR (working directory) for app and add code
WORKDIR /source
COPY . .

RUN go build -o linear cmd/linear-pr-checker/main.go

# RUN as the mhuser user
USER mhuser

# Run the housebot command by default when the container starts.
# runs command at most recent WORKDIR
ENTRYPOINT ["./linear"]