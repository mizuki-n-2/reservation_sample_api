wire:
	wire di/wire.go

test:
	go test ./... -v

cover:
	mkdir -p coverage
	go test -cover ./... -coverprofile=coverage/cover.out
	go tool cover -html=coverage/cover.out -o coverage/cover.html
	open coverage/cover.html
	
lint:
	golangci-lint run
