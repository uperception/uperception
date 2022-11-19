DOCKER_IMAGE ?= mmonitoring/mmonitoring
DOCKER_TAG ?= latest
DOCKER_REF := $(DOCKER_IMAGE):$(DOCKER_TAG)

DOCKER_COLLECTOR_IMAGE ?= mmonitoring/collector
DOCKER_COLLECTOR_TAG ?= latest
DOCKER_COLLECTOR_REF := $(DOCKER_COLLECTOR_IMAGE):$(DOCKER_COLLECTOR_TAG)

.PHONY: dev
dev:
	docker compose --profile database --env-file="./.env" up

.PHONY: build 
build:
	go build -v ./...

.PHONY: docker-build
docker-build:
	docker build docker/api -t $(DOCKER_REF)

.PHONY: build-collector
docker-build-collector:
	docker build docker/lighthouse -t $(DOCKER_COLLECTOR_REF)

.PHONY: test
test:
	go test -v ./...

.PHONY: create-env
create-env:
	./scripts/create-sample-env.sh


