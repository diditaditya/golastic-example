FROM golang:1-buster

WORKDIR /app
COPY . .

CMD ["go", "run", "main.go"]