# FROM python:3.10-slim
FROM python:3.7.5-slim

COPY docker/app/requirements.txt .
RUN pip install -r requirements.txt

WORKDIR /app
COPY . .
WORKDIR /app/src

ENTRYPOINT ["python", "app.py"]
# ENTRYPOINT ["flask", "run", "--host", "0.0.0.0", "--port", "8080"]