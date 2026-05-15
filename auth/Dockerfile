FROM golang:1.22.1-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://package-mirror.liara.ir/repository/go/
RUN go env -w GOSUMDB=off
RUN go build -o main cmd/main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY envs/dev.env ./envs/

CMD [ "/app/main" ]