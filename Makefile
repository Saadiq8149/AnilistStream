# Image
IMAGE ?= 12345saadiq/aniliststream
TAG ?= latest

# Compose files
COMPOSE=docker compose
COMPOSE_FILE=docker-compose.yml
COMPOSE_DEV_FILE=docker-compose.dev.yml

# Default target
.DEFAULT_GOAL := help

help:
	@echo "Available commands:"
	@echo " make build        Build docker image"
	@echo " make push         Push docker image"
	@echo " make run          Run container locally"
	@echo " make up           Start docker compose"
	@echo " make dev          Start dev compose"
	@echo " make down         Stop containers"
	@echo " make restart      Restart containers"
	@echo " make logs         Show logs"
	@echo " make clean        Remove containers + volumes"

build:
	docker build -t $(IMAGE):$(TAG) .

push:
	docker push $(IMAGE):$(TAG)

run:
	docker run -p 8080:8080 $(IMAGE):$(TAG)

up:
	$(COMPOSE) -f $(COMPOSE_FILE) up --build -d

dev:
	$(COMPOSE) -f $(COMPOSE_FILE) -f $(COMPOSE_DEV_FILE) up --build

down:
	$(COMPOSE) -f $(COMPOSE_FILE) down

restart:
	$(COMPOSE) -f $(COMPOSE_FILE) down
	$(COMPOSE) -f $(COMPOSE_FILE) up -d

logs:
	$(COMPOSE) logs -f

clean:
	$(COMPOSE) down -v --remove-orphans
