FROM golang:1.21.6-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -C /app/cmd/storage-api

EXPOSE 8080

ENTRYPOINT exec go run cmd/storage-api/main.go --config=deployment/local/storage-api.yaml