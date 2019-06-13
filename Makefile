
.PHONY: test
test:
	go test -v -race -cover -coverprofile=coverage.out  ./...

.PHONY: cover
cover: test
	go tool cover -html=coverage.out

