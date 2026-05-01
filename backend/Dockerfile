FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api ./cmd/api
RUN go build -o seed ./cmd/seed

FROM alpine:latest AS api

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/api .

EXPOSE 8080

CMD ["./api"]

FROM alpine:latest AS seed

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/seed .
COPY --from=builder /app/data ./data

CMD ["./seed"]