FROM golang:1.16-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

ADD . .

RUN go build -o ./bin/sgp

FROM alpine:3.12

COPY --from=build-env /src/bin/sgp /app/sgp

CMD ["/app/sgp"]
