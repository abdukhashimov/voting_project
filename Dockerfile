# Build stage
FROM golang:1.19.2-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o open_budget_core cmd/main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/open_budget_core .
COPY --from=builder /app/configs configs/

EXPOSE 8080
ENTRYPOINT [ "/app/open_budget_core" ]
