FROM python:3.9-slim
# ADD ./pyscript/requirements.txt .
ADD ./file_consumer .
RUN  apt-get update -y &&\
        apt install -y python3-pip &&\
        python3 -m pip install --no-cache-dir -r requirements.txt
CMD ["python", "file_consumer.py" ]