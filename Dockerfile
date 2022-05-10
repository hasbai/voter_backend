FROM golang:1.16 as builder

WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go build -o go

EXPOSE 8080

FROM alpine

WORKDIR /app

COPY --from=builder /app/go .

ENV GIN_MODE=release

EXPOSE 8000

ENTRYPOINT ["./go"]