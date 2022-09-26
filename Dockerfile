FROM python:3.9
ADD requirements.txt .
ADD app.py .
RUN pip3 install --no-cache-dir -r requirements.txt
ENTRYPOINT FLASK_APP=app.py flask run --host=0.0.0.0