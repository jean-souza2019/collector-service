FROM golang:1.22-alpine AS builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -o collector-service cmd/main.go

FROM alpine:latest AS runtime
COPY --from=builder /app/collector-service /server
EXPOSE 3000
RUN apk add --no-cache libgcc libstdc++ librdkafka-dev
ENTRYPOINT ["/server"]

FROM golang:1.22-alpine AS dev
RUN apk add --no-cache gcc musl-dev libgcc libstdc++ librdkafka-dev
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
EXPOSE 3000
CMD ["go", "run", "cmd/collector-service/main.go"]