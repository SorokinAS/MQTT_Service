FROM golang:1.20.12-alpine

WORKDIR /app

COPY . .

RUN go build -o ./bin/app ./cmd/main.go

CMD [ "./bin/app" ]