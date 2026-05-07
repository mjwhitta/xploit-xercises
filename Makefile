DATE := $(shell date +%Y%m%d)
NAME := ghcr.io/mjwhitta/xploit-xercises

all: build-alpine build-debian build-rolling build-slim build-ubuntu;

build: build-rolling;

build-alpine:
	@mkdir -p ./local
	@sed "s/#alpine# //g" ./Dockerfile >./local/Dockerfile
	@docker build --pull -f ./local/Dockerfile -t $(NAME):alpine .
	@docker buildx prune -f &>/dev/null
	@rm -f ./local/Dockerfile

build-debian:
	@mkdir -p ./local
	@sed "s/#debian# //g" ./Dockerfile >./local/Dockerfile
	@docker build --pull -f ./local/Dockerfile -t $(NAME):debian .
	@docker buildx prune -f &>/dev/null
	@rm -f ./local/Dockerfile

build-rolling:
	@mkdir -p ./local
	@cp -f ./Dockerfile ./local/Dockerfile
	@sed -i -e "s/#rolling# //g" ./local/Dockerfile
	@docker build --pull -f ./local/Dockerfile -t $(NAME):rolling .
	@docker buildx prune -f &>/dev/null
	@rm -f ./local/Dockerfile

build-slim:
	@mkdir -p ./local
	@sed "s/#slim# //g" ./Dockerfile >./local/Dockerfile
	@docker build --pull -f ./local/Dockerfile -t $(NAME):slim .
	@docker buildx prune -f &>/dev/null
	@rm -f ./local/Dockerfile

build-ubuntu:
	@mkdir -p ./local
	@cp -f ./Dockerfile ./local/Dockerfile
	@sed -i -e "s/#ubuntu# //g" ./local/Dockerfile
	@docker build --pull -f ./local/Dockerfile -t $(NAME):ubuntu .
	@docker buildx prune -f &>/dev/null
	@rm -f ./local/Dockerfile

clean: down
	@if [[ -f ./local/Dockerfile ]]; then \
	    rm -f ./local/Dockerfile; \
	fi
	@docker rmi \
		$(NAME):alpine $(NAME):debian $(NAME):rolling $(NAME):slim \
		$(NAME):ubuntu \
	    $$(docker images | awk '/<none>/ {print $$3}') \
	    >/dev/null 2>&1 || true

clena: clean;

down:
	@if [[ -f ./local/compose.yaml ]]; then \
	    docker compose -f ./local/compose.yaml down -v || true; \
	    rm -f ./local/compose.yaml; \
	else \
	    docker compose down -v || true; \
	fi

release:
	@git tag -f v1.0.$(DATE)
	@git push
	@git push --tags

up: up-rolling;

up-alpine: down build-alpine
	@sed "s/rolling/alpine/" ./compose.yaml >./local/compose.yaml
	@docker compose -f ./local/compose.yaml up -d
	@docker compose logs -f || true

up-debian: down build-debian
	@sed "s/rolling/debian/" ./compose.yaml >./local/compose.yaml
	@docker compose -f ./local/compose.yaml up -d
	@docker compose logs -f || true

up-rolling: down build-rolling
	@cp -f ./compose.yaml ./local/compose.yaml
	@docker compose -f ./local/compose.yaml up -d
	@docker compose logs -f || true

up-slim: down build-slim
	@sed "s/rolling/slim/" ./compose.yaml >./local/compose.yaml
	@docker compose -f ./local/compose.yaml up -d
	@docker compose logs -f || true

up-ubuntu: down build-ubuntu
	@sed "s/rolling/ubuntu/" ./compose.yaml >./local/compose.yaml
	@docker compose -f ./local/compose.yaml up -d
	@docker compose logs -f || true
