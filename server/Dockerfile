# builder stage
FROM golang:1.25-alpine AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
WORKDIR /src/server
RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/admin-dashboard .

# final runtime image
FROM alpine:3.19
RUN apk add --no-cache ca-certificates mariadb-client
WORKDIR /app
COPY --from=builder /usr/local/bin/admin-dashboard /usr/local/bin/admin-dashboard
COPY server/configs ./configs
COPY server/sql ./sql

ENV ROOT_FOLDER=/app
ENV API_HOST=0.0.0.0
EXPOSE 8081

CMD ["admin-dashboard"]
