FROM golang:1.17.3-alpine

WORKDIR /app

COPY . .

RUN go build -o /app/cmd/server/apiserver -v ./cmd/server/main.go

RUN apk add --no-cache go

ENTRYPOINT ["tail", "-f", "/dev/null"]