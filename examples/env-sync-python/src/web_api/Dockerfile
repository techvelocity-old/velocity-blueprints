FROM python:3.10-alpine
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY main.py .
CMD ["sh", "-c", "uvicorn main:app --host 0.0.0.0 --port 8000 --reload"]
