test-integration:
	docker compose build test-integration
	docker compose run --rm test-integration

api:
	docker compose run --rm migrate
	docker compose up --build api

