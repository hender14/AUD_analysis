# FROM python:3.10-slim
FROM python:3.7.5-slim

COPY docker/test/requirements.txt .
RUN pip install -r requirements.txt

WORKDIR /app
COPY . .
WORKDIR /app/src

# RUN pip install --upgrade pip

ENV FLASK_APP test_app.py