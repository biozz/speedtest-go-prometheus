FROM golang:1.16-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /src

ADD . .

RUN go build -o ./bin/sgp

FROM alpine:3.12

COPY --from=builder /src/bin/sgp /app/sgp

WORKDIR /app

CMD ["./sgp"]
