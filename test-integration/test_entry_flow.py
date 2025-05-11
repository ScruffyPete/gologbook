import json
import pytest
import requests  # type: ignore


@pytest.fixture
def api_client():
    class TestClient:
        def __init__(self, base_url):
            self.base_url = base_url

        def post(self, url, json=None, headers=None):
            return requests.post(f"{self.base_url}{url}", json=json, headers=headers)

        def get(self, url, headers=None):
            return requests.get(f"{self.base_url}{url}", headers=headers)

    return TestClient("http://test-api:8080")


@pytest.fixture
def insights_client():
    class TestClient:
        def __init__(self, base_url):
            self.base_url = base_url

        def get(self, url, headers=None):
            return requests.get(f"{self.base_url}{url}", headers=headers)

    return TestClient("http://test-insights:8081")


def test_entry_flow(api_client, insights_client):
    email = "test-integration@test.com"
    password = "testpassword"

    signup_response = api_client.post(
        "/api/signup",
        json={
            "email": email,
            "password": password,
            "confirmPassword": password,
        },
    )
    assert signup_response.status_code == 201

    login_response = api_client.post(
        "/api/login", json={"email": email, "password": password}
    )
    assert login_response.status_code == 200
    auth_token = login_response.json()["token"]
    assert auth_token is not None

    headers = {"Authorization": f"Bearer {auth_token}"}

    project = {"name": "test-project"}
    response = api_client.post("/api/projects/", json=project, headers=headers)
    assert response.status_code == 201
    project_id = response.json()["id"]

    entry = {"body": "test-entry", "project_id": project_id}
    response = api_client.post("/api/entries/", json=entry, headers=headers)
    assert response.status_code == 201
    entry_id = response.json()["id"]

    response = api_client.get(f"/api/entries?project_id={project_id}", headers=headers)
    entries = response.json()
    assert response.status_code == 200
    assert len(entries) == 1
    assert entries[0]["body"] == "test-entry"
    assert entries[0]["projectId"] == project_id

    # Check that the insights service has processed the entry
    insights_response = insights_client.get("/status")
    assert insights_response.status_code == 200
    insights_data = insights_response.json()
    assert insights_data["has_started"]
    assert insights_data["last_processed"] is not None
    assert insights_data["last_processed"] > 0

    last_processed_message = json.loads(insights_data["last_processed_message"])
    assert last_processed_message is not None
    assert last_processed_message["type"] == "new_entry"
    assert last_processed_message["payload"]["entry_id"] == entries[0]["id"]

    # Check that the insights service has added new insights
    insights_response = api_client.get(
        f"/api/insights?project_id={project_id}", headers=headers
    )
    assert insights_response.status_code == 200
    insights = insights_response.json()
    assert len(insights) == 1
    assert insights[0]["body"] == f"Insight for entry {entry_id}: {entry['body'][:100]}"
    assert insights[0]["projectId"] == project_id
