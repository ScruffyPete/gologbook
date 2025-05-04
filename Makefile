test-integration:
	docker compose build test-integration
	docker compose run --rm --build test-integration

api:
	docker compose run --rm --build migrate
	docker compose up --build api

