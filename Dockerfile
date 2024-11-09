FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o upassed-form-service ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir -p /upassed-form-service/config
RUN mkdir -p /upassed-form-service/migration/scripts
COPY --from=builder /app/upassed-form-service /upassed-form-service/upassed-form-service
COPY --from=builder /app/config/* /upassed-form-service/config
COPY --from=builder /app/migration/scripts/* /upassed-form-service/migration/scripts
RUN chmod +x /upassed-form-service/upassed-form-service
ENV APP_CONFIG_PATH="/upassed-form-service/config/local.yml"
EXPOSE 44044
CMD ["/upassed-form-service/upassed-form-service"]
