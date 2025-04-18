FROM python:3.13-slim

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt


COPY ./app /app



ENV PYTHONPATH=/

EXPOSE 6013


CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "6013"]
