# Stage 1: Build binary
FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server main.go

# Stage 2: Minimal runtime image
FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Taipei

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
