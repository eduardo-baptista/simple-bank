.PHONY: test
test:
	go test -v ./...

.PHONY: test.cover
test.cover:
	go test -coverpkg=./internal/domain/entity,./internal/usecase/...,./internal/infrastructure/http/handler/...,./internal/infrastructure/repository/... -coverprofile=./coverage.out ./...
	go tool cover -html=./coverage.out -o coverage.html

.PHONY: start.dev
start.dev:
	docker compose watch

.PHONY: start
start:
	docker compose up -d

.PHONY: stop
stop:
	docker compose down -v --remove-orphans

.PHONY: generate
generate:
	go generate ./...