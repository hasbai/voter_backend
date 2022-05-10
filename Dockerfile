FROM golang:1.16 as builder

WORKDIR /app
ENV GO111MODULE=on

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o go

FROM alpine

WORKDIR /app

COPY --from=builder /app/go .

ENV GIN_MODE=release

EXPOSE 8000

ENTRYPOINT ["./go"]