COMPOSE = docker compose

build:
	${COMPOSE} -f docker-compose.yml up -d --build
exec:
	${COMPOSE} exec golang-srv /bin/sh

.PHONY:
cp:
	${COMPOSE} $(c)