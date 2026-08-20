FROM golang:1.24

WORKDIR /app

COPY go.mod ./
COPY . .

RUN go build -o app .

EXPOSE 8080

CMD ["./app"]