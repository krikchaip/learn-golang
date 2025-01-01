.PHONY: setup
setup:
	@docker compose up -d

.PHONY: run
run:
	@go run ./cmd/web -addr=:3000

.PHONY: watch
watch:
	@gow run ./cmd/web -addr=:3000
