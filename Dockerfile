FROM golang:1.16.2-alpine3.13 as builder

COPY go.mod go.sum /app/
WORKDIR /app/src/api
RUN go mod download

WORKDIR /app
COPY . .
WORKDIR /app/src/api
RUN go build -o main

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /app/src/api /app/
WORKDIR /app

ENV GIN_MODE="release"
CMD ["./main"]
