FROM golang:1.23 as builder

WORKDIR /app

COPY ../../go.mod ../../go.sum ./

RUN go mod download

COPY ./server/nethttp /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /datetime-server .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /datetime-server .

EXPOSE 8000

CMD ["./datetime-server"]
