FROM golang:1.23.2-alpine3.19 AS builder

WORKDIR "/app"

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV PORT=8080

RUN apk add --no-cache git ca-certificates

COPY . .

RUN go build -o loyalty-system  ./cmd/main.go

FROM alpine:3.19

RUN apk update
RUN apk --no-cache add ca-certificates bash

WORKDIR /root/

COPY --from=builder /app/loyalty-system .
RUN chmod +x loyalty-system

EXPOSE ${PORT}

CMD ["./loyalty-system"]
