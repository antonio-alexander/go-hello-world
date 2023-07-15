## ----------------------------------------------------------------------
## This makefile can be used to execute common functions to interact with
## the source code, these functions ease local development and can also be
## used in CI/CD pipelines.
## ----------------------------------------------------------------------

golangcilint_version=v1.44.2
swagger_version=v0.29.0
godoc_version=v0.1.10

rsa_bits=4096
ssl_subj="//C=US"

swagger_port=8000
godoc_port=8001

# REFERENCE: https://stackoverflow.com/questions/16931770/makefile4-missing-separator-stop
help: ## - Show this help.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

check-lint: ## - validate/install golangci-lint installation
	@which golangci-lint || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@${golangcilint_version})

lint: check-lint ## - lint the source with verbose output
	@golangci-lint run --verbose

# Reference: https://medium.com/@pedram.esmaeeli/generate-swagger-specification-from-go-source-code-648615f7b9d9
check-swagger: ## - validate/install swagger (v0.29.0)
	@which swagger || (go install github.com/go-swagger/go-swagger/cmd/swagger@${swagger_version})

swagger: check-swagger ## - generate the swagger.json
	@swagger generate spec --work-dir=./internal/swagger --output ./tmp/swagger.json --scan-models

validate-swagger: swagger ## - validate the swagger.json
	@swagger validate ./tmp/swagger.json

serve-swagger: swagger ## - serve (web) the swagger.json
	@swagger serve -F=swagger ./tmp/swagger.json -p ${swagger_port} --no-open

check-godoc: ## - validate/install godoc
	@which godoc || (go install golang.org/x/tools/cmd/godoc@${godoc_version})

serve-godoc: check-godoc ## - serve (web) the godocs
	@godoc -http :${godoc_port}

build: ## - build the source (latest)
	@docker compose build --build-arg GIT_COMMIT=`git rev-parse HEAD` \
	--build-arg GIT_BRANCH=`git rev-parse --abbrev-ref HEAD`
	@docker image prune -f

run: ## - run the service and its dependencies (docker) detached
	@docker compose up -d

check-openssl: ## Check if openssl is installed
	@which openssl || echo "openssl not found"

gen-certificates: check-openssl ## Generate public/private SSL certificates
	@openssl req -x509 -sha256 -nodes -days 1 -newkey rsa:${rsa_bits} -keyout ./certs/ssl.key -out ./certs/ssl.crt -subj ${ssl_subj}

stop:
	@docker compose down
