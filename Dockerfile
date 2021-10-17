FROM golang:1.17.2-alpine3.14
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

ENV CGO_ENABLED=0
RUN chmod +x scripts/start.sh
CMD ./scripts/start.sh && tail -f /dev/null