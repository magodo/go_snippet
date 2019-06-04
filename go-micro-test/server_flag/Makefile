all:
	protoc -I . --go_out=. \
		--micro_out=. \
		proto/greeter/greeter.proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server cmd/server/*.go

