FROM golang:1.22-bookworm as builder
WORKDIR /app/
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./bin/main.go

FROM gcr.io/distroless/base-debian12
WORKDIR /app/
COPY ./.config /app
COPY ./files /app
COPY --from=builder /app/main /app
CMD ["./main", "-config", "./.config/config.toml"]
