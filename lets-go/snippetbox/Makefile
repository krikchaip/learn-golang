KEY_FILE:=tls/key.pem
CERT_FILE:=tls/cert.pem

.PHONY: gen-cert
gen-cert:
	@openssl req -x509 -nodes -newkey rsa:2048 -days 365 -keyout ${KEY_FILE} -out ${CERT_FILE}

.PHONY: setup
setup:
	@[ ! -f "${KEY_FILE}" ] && [ ! -f "${CERT_FILE}" ] && $(MAKE) gen-cert || true
	@go mod download
	@docker compose up -d

.PHONY: run
run:
	@go run ./cmd/web -addr=:3000

.PHONY: watch
watch:
	@gow run ./cmd/web -addr=:3000

.PHONY: test
test:
	@gow test -v "./..."
