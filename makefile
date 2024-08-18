make build-local:
	cd ./build && docker compose build --no-cache

make up-local:
	cd ./build && docker compose up

make backend:
	cd ./build && docker compose exec backend bash

