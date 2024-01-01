IMAGE=jdwillmsen/go-discord-subcommand-bot
TAG=latest
VERSION=1.0.0

build:
	docker build . -t "${IMAGE}:${TAG}" -t "${IMAGE}:${VERSION}"

run: build
	docker run --env-file .env -d "${IMAGE}:${TAG}"

publish: build
	docker push -a "${IMAGE}"