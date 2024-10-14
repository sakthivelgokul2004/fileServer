FROM golang:1.23-alpine

WORKDIR /app
COPY go.mod go.sum ./
COPY .   .

RUN go mod tidy
RUN go install github.com/air-verse/air@latest


CMD  air -c  .air.toml

