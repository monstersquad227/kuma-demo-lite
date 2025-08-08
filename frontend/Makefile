##@ 🎯 The commands are:

IMG ?= frontend:latest

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<command>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: run
run: ## Running backend.
	@npm install
	@npm run server


.PHONY: build-docker
build-docker: ## Build docker image with the backend.
	@echo "📝 Preparing dist..."
	@npm install
	@npm run build
	@echo "✅ dist created."

	@echo "📝 Generating temporary nginx.conf file..."
	@printf "\
server {\n\
    listen       80;\n\
    server_name  _;\n\
    root         /opt/;\n\
	\n\
    location / {\n\
        index index.html;\n\
    }\n\
	\n\
    location /healthz {\n\
        return 200 \"ok\";\n\
    }\n\
}\n\
" > nginx.conf
	@echo "✅ nginx.conf file created."

	@echo "📝 Generating temporary Dockerfile..."
	@echo "\
FROM nginx:latest \n\
WORKDIR /opt \n\
COPY nginx.conf /etc/nginx/conf.d/ \n\
ADD ./dist . \n\
" > Dockerfile
	@echo "✅ Dockerfile created."

	@echo "🐳 Building Docker image: ${IMG}..."
	@docker build -t ${IMG} .
	@echo "✅ Build complete."
	