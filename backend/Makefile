##@ 🎯 The commands are:

GO_ENV := GO111MODULE=on GOPROXY=https://goproxy.cn,direct
IMG ?= backend:latest

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<command>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: run
run: ## Running backend.
	@go run main.go

.PHONY: build-linux
build-linux: ## Build backend Linux binary.
	GOOS=linux GOARCH=amd64 $(GO_ENV) go build -v -o backend

.PHONY: build-docker
build-docker: ## Build docker image with the backend.
	@echo "📝 Generating temporary Dockerfile..."
	@printf "\
FROM nginx:latest \n\
WORKDIR / \n\
COPY ./main.go . \n\
RUN apt update && apt install golang -y && go mod init backend && go mod tidy && $(GO_ENV) go build -o backend\n\
ENTRYPOINT [\"backend\"] \n\
" > Dockerfile
	@echo "✅ Dockerfile created."

	@echo "🐳 Building Docker image: ${IMG}..."
	@echo "Builder image"
	@docker build -t ${IMG} .