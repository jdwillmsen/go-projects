IMAGE=jdwillmsen/go-jwt-project
TAG=latest
VERSION=1.0.0

mongodb:
	docker run \
    -d \
    --name mongodb \
    -p 27017:27017 \
    -e MONGO_INITDB_ROOT_USERNAME=user \
    -e MONGO_INITDB_ROOT_PASSWORD=secret \
    mongo
build:
	go build main.go
execute: build
	./main.exe
run:
	go run main.go

docker-build:
	docker build . -t "${IMAGE}:${TAG}" -t "${IMAGE}:${VERSION}"
docker-run: docker-build
	docker run --env-file .env -d "${IMAGE}:${TAG}"
docker-compose:
	docker compose up -d
docker-publish: docker-build
	docker push -a "${IMAGE}"

github-build:
	docker build . -t "ghcr.io/${IMAGE}:${VERSION}" -t "ghcr.io/${IMAGE}:${VERSION}"
github-publish: github-build
	docker push -a "ghcr.io/${IMAGE}"