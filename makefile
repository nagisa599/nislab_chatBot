build-local:
	cd ./build && docker compose build --no-cache

up-local:
	cd ./build && docker compose up

backend:
	cd ./build && docker compose exec backend bash

