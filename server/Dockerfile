FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
COPY .   .

#RUN go mod tidy

CMD  air -c  .air.toml

