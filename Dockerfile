# Dockerfile (root of project)
FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/app ./cmd || \
    GO111MODULE=on go build -o /app/bin/app ./...

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /app/bin/app /usr/local/bin/app
EXPOSE 8989
CMD ["/usr/local/bin/app"]
