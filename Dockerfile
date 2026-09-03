FROM golang:1.26.5

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app .
CMD ["./app"]
