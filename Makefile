# 變量定義
APP_NAME=hr-system
DOCKER_COMPOSE=docker-compose
GO=go
GOTEST=$(GO) test
BINARY_NAME=main

# 顏色定義
COLOR_RESET=\033[0m
COLOR_GREEN=\033[32m
COLOR_YELLOW=\033[33m

# 默認目標
.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo "使用方法:"
	@echo "  make [target]"
	@echo ""
	@echo "目標列表:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run: ## 本地運行應用
	$(GO) run main.go

.PHONY: build
build: ## 構建應用
	@echo "${COLOR_GREEN}Building $(APP_NAME)...${COLOR_RESET}"
	$(GO) build -o $(BINARY_NAME) main.go

.PHONY: docker-build
docker-build: ## 構建 Docker 映像
	@echo "${COLOR_GREEN}Building Docker image...${COLOR_RESET}"
	$(DOCKER_COMPOSE) build

.PHONY: docker-up
docker-up: ## 啟動所有 Docker 服務
	@echo "${COLOR_GREEN}Starting Docker services...${COLOR_RESET}"
	$(DOCKER_COMPOSE) up -d

.PHONY: docker-down
docker-down: ## 停止所有 Docker 服務
	@echo "${COLOR_GREEN}Stopping Docker services...${COLOR_RESET}"
	$(DOCKER_COMPOSE) down

.PHONY: docker-logs
docker-logs: ## 查看 Docker 日誌
	@echo "${COLOR_GREEN}Showing Docker logs...${COLOR_RESET}"
	$(DOCKER_COMPOSE) logs -f

.PHONY: docker-restart
docker-restart: docker-down docker-up ## 重啟所有 Docker 服務

.PHONY: dev
dev: docker-build docker-up ## 啟動開發環境

.PHONY: clean
clean: ## 清理構建文件
	@echo "${COLOR_GREEN}Cleaning...${COLOR_RESET}"
	rm -f $(BINARY_NAME)
	rm -f coverage.out 