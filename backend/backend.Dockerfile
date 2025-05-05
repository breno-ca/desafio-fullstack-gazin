FROM golang:1.24-alpine3.20 AS builder

WORKDIR /build

COPY . .

RUN go mod download
RUN go build -o main .

FROM alpine:3.20

WORKDIR /run

COPY --from=builder /build/main .

EXPOSE 8080

CMD ["./main"]

