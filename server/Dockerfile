# builder stage
FROM golang:1.25-alpine AS builder
WORKDIR /src

# Copier go.mod et go.sum depuis server/ (contexte = racine du projet)
COPY server/go.mod server/go.sum ./
RUN go mod download

# Copier tout le code source de server/
COPY server/ .

# Compiler depuis la racine de /src (où se trouve main.go)
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