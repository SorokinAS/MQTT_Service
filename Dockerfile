FROM golang:1.20.12-alpine

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/main.go

CMD [ "./main" ]