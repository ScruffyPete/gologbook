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


def test_entry_flow(api_client):
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

    project = {"name": "integration test"}
    response = api_client.post("/api/projects/", json=project, headers=headers)
    assert response.status_code == 201
    project_id = response.json()["id"]

    entry = {"body": "integration test", "project_id": project_id}
    response = api_client.post("/api/entries/", json=entry, headers=headers)
    assert response.status_code == 201

    response = api_client.get(f"/api/entries?project_id={project_id}", headers=headers)
    entries = response.json()
    assert response.status_code == 200
    assert len(entries) == 1
    assert entries[0]["body"] == "integration test"
    assert entries[0]["projectId"] == project_id

    # TODO: test that the entry reached insights and was processed
