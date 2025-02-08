FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY ../go.mod ../go.sum ./
RUN go mod tidy
COPY ../ ./
RUN go mod tidy && go build -o backend .
CMD ["./backend"]