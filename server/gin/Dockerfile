FROM golang:1.23 AS builder

WORKDIR /app

COPY ../../go.mod ../../go.sum ./

RUN go mod download

COPY ./server/gin /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /datetime-server .

FROM alpine:latest

WORKDIR /root/

RUN apk --no-cache add ca-certificates

COPY --from=builder /datetime-server .

CMD ["./datetime-server"]
