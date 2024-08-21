FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
COPY go.mod go.sum ./
COPY .   .

#RUN go mod tidy
#RUN go mod vendor

CMD  air -c .air.toml

