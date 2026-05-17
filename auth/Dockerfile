FROM golang:1.22.1-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://mirror.abrha.net/repository/go/
RUN go env -w GOSUMDB=off
RUN go build -o main cmd/main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/envs/ ./envs/

CMD [ "/app/main" ]