FROM golang:alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -ldflags="-s -w" -o /app/main ./cmd/main/main.go

FROM alpine:latest

RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]