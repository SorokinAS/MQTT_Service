build-container:
	docker build -t gateway .
	docker run --name gateway gateway

run-container:
	docker run gateway

build-local:
	go build -o .\bin\ .\cmd\main.go

run-local:
	go run .\cmd\main.go	