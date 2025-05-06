test-integration:
	docker compose build test-integration
	docker compose up -d redis
	docker compose run --rm --build test-integration
	docker compose down -v

api:
	docker compose run --rm --build migrate
	docker compose up --build api
