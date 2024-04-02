# Create Builder Image, to compile the source code into an executable
FROM golang:1.18.2-alpine3.16 as base
RUN apk update
WORKDIR /src/goppotunities
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o goppotunities ./cmd/api

# Create the final image, running the API and exposing port 8080
FROM alpine:3.16 as binary
EXPOSE $8080
CMD ["./goppotunities"]
