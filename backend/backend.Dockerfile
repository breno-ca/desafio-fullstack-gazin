FROM golang:1.24-alpine3.20 AS builder

WORKDIR /build

COPY . .

RUN chmod +x entrypoint.sh
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go mod download
RUN go build -o main .

FROM alpine:3.20

WORKDIR /run

COPY --from=builder \
	/build/internal/database/migrations/mysql \
	/run/internal/database/migrations/mysql
COPY --from=builder /build/entrypoint.sh /run/
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /build/main /run/

EXPOSE 8080

ENTRYPOINT ["/run/entrypoint.sh"]
CMD ["/run/main"]
