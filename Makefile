.PHONY: test-backend-integration test-insights-integration test-integration api

test-backend-integration:
	docker compose run --rm --build test-backend-integration || true
	docker compose down -v

test-insights-integration:
	docker compose run --rm --build test-insights-integration || true
	docker compose down -v

test-integration:
	docker compose run --rm --build test-integration || true
	docker compose down -v

api:
	docker compose run --rm --build migrate
	docker compose up --build api
