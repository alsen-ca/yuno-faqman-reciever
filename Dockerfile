FROM golang:1.25.5-trixie

WORKDIR /app

COPY main.go .
COPY go.mod .
COPY go.sum .
COPY internal/ .

# TODO