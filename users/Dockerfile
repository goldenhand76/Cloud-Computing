FROM golang:1.24-alpine3.23 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go



FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

RUN chmod +x /app/main

EXPOSE 8083
CMD [ "/app/main" ]