FROM python:3.13-slim

# Keeps Python from generating .pyc files in the container
ENV PYTHONDONTWRITEBYTECODE=1
# Turns off buffering for easier container logging
ENV PYTHONUNBUFFERED=1

COPY requirements.txt .
RUN python -m pip install -r requirements.txt

COPY . /app
WORKDIR /app


ENV PYTHONPATH=/app

RUN apt-get update && apt-get install -y curl
HEALTHCHECK --interval=5s --timeout=5s --start-period=5s --retries=3 CMD [ "curl", "-f", "http://localhost:8081/healthz" ]

CMD ["python", "main.py"]
