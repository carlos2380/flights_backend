FROM golang:1.21-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /flights cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /flights .

ENV PORT=8000

EXPOSE 8000

CMD ["./flights"]
