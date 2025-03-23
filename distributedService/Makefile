proto:
	rm -f api/v1/*.go
	protoc --proto_path=. --go_out=. --go_opt=paths=source_relative \
    api/v1/*.proto


test:
	go test -race ./...

.PHONY: protoc test
