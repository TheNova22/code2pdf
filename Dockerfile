FROM ubuntu:18.04
ADD requirements.txt .
RUN  apt-get update -y &&\
        apt install -y python3-pip &&\
        python3 -m pip install --no-cache-dir -r requirements.txt&&\
        apt-get -y install wkhtmltopdf && apt-get -y install xvfb
ADD langComment.py .
ADD parser.py .
ADD code2pdf.py .
ENTRYPOINT [ "python3", "./code2pdf.py"]