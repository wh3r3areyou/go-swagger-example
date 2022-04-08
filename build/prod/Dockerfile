FROM golang:1.17.3-alpine

WORKDIR /app

COPY . .

RUN go build -o /app/cmd/server/apiserver -v ./cmd/server/main.go

CMD ["/app/cmd/server/apiserver"]