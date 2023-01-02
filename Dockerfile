FROM golang:1.18.9-alpine3.16 as builder
WORKDIR /app
COPY . .
RUN go build -ldflags="-w -s" -o main ./cmd/web

FROM alpine:3.16
RUN adduser -D -u 10000 adel
RUN mkdir /app/ && chown adel /app/
USER adel
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/pkg/calendar/static/favicon.ico .
EXPOSE "8000"
ENTRYPOINT ["./main"]