FROM golang:1.22 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGI_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest as production
COPY --from=builder /app .
CMD ["./app"]



