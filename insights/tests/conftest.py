import pytest


def pytest_addoption(parser: pytest.Parser):
    parser.addoption(
        "--run-db", action="store_true", help="run tests that require a database"
    )


def pytest_configure(config: pytest.Config):
    config.addinivalue_line(
        "markers", "db: mark test as requiring a database connection"
    )


def pytest_collection_modifyitems(config: pytest.Config, items: list[pytest.Item]):
    if config.getoption("--run-db"):
        return

    skip_db = pytest.mark.skip(reason="need --run-db option to run")
    for item in items:
        if item.get_closest_marker("db") is not None:
            item.add_marker(skip_db)
