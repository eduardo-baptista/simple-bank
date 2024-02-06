.PHONY: test
test:
	go test -v ./...

.PHONY: start.dev
start.dev:
	docker compose watch

.PHONY: generate
generate:
	go generate ./...