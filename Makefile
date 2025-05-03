test-integration:
	docker compose build test-integration
	docker compose run --rm test-integration
