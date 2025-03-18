# Description: Dockerfile for building the go application
FROM golang:1.23.7-alpine3.21 AS builder

WORKDIR /app

COPY . .

RUN go build -o main main.go

# final stage
FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .


EXPOSE 8080

CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]
