# syntax=docker/dockerfile:1
FROM golang:1.16-alpine
ENV GO111MODULE  on
WORKDIR /source
COPY . .
RUN go mod download
RUN go build -o /linear cmd/linear-pr-checker/main.go
ENTRYPOINT ["/linear"]