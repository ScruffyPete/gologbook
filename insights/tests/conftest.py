import pytest


def pytest_addoption(parser: pytest.Parser):
    parser.addoption(
        "--run-db", action="store_true", help="run tests that require a database"
    )
    parser.addoption(
        "--run-queue", action="store_true", help="run tests that require a queue"
    )


def pytest_configure(config: pytest.Config):
    config.addinivalue_line(
        "markers", "db: mark test as requiring a database connection"
    )
    config.addinivalue_line("markers", "queue: mark test as requiring a queue")


def pytest_collection_modifyitems(config: pytest.Config, items: list[pytest.Item]):
    skip_db = not config.getoption("--run-db")
    skip_queue = not config.getoption("--run-queue")

    for item in items:
        if skip_db and item.get_closest_marker("db") is not None:
            item.add_marker(pytest.mark.skip(reason="need --run-db option to run"))
        if skip_queue and item.get_closest_marker("queue") is not None:
            item.add_marker(pytest.mark.skip(reason="need --run-queue option to run"))
