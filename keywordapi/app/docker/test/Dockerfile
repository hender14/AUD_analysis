# FROM python:3.10-slim
FROM python:3.7.5-slim

COPY docker/test/requirements.txt .
RUN pip install -r requirements.txt

WORKDIR /app
COPY . .
WORKDIR /app/src

# RUN pip install --upgrade pip

ENV FLASK_APP test_app.py

# CMD ["tail", "-f", "/dev/null"]
# ENTRYPOINT ["flask", "run", "--host", "0.0.0.0", "--port", "8080"]