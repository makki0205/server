product-build:
	rm -rf ./cmd/main
	env GOOS=linux env GOARCH=amd64 env CGO_ENABLED=0 go build -o ./cmd/main product/main.go

schedule-build:
	rm -rf ./cmd/main
	env GOOS=linux env GOARCH=amd64 env CGO_ENABLED=0 go build -o ./cmd/main schedule/main.go

batch-build:
	rm -rf ./cmd/main
	env GOOS=linux env GOARCH=amd64 env CGO_ENABLED=0 go build -o ./cmd/main batch/main.go

log-build:
	rm -rf ./cmd/main
	env GOOS=linux env GOARCH=amd64 env CGO_ENABLED=0 go build -o ./cmd/main log/main.go