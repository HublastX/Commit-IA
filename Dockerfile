FROM python:3.10-slim
WORKDIR /app
COPY . .
RUN pip install -r requirements.txt
EXPOSE 3004
ENV FLASK_APP=main.py

CMD ["gunicorn", "-w", "4", "-b", "0.0.0.0:3004", "main:app"]