FROM python:3.9-slim
# ADD ./pyscript/requirements.txt .
ADD ./pyscript .
RUN  apt-get update -y &&\
        apt install -y python3-pip &&\
        python3 -m pip install --no-cache-dir -r requirements.txt&&\
        apt-get -y install wkhtmltopdf && apt-get -y install xvfb
ENTRYPOINT FLASK_APP=app.py flask run --host=0.0.0.0