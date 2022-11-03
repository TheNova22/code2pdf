import json
from flask import Flask, jsonify, send_file, url_for, redirect, request, Blueprint
from  pygments import lexers, formatters, styles, highlight
import pdfkit
from xvfbwrapper import Xvfb
import requests
import os
from langComment import langComment
from cassandra.cluster import Cluster
from cassandra.query import SimpleStatement

app = Flask(__name__)
vdisplay = Xvfb()
vdisplay.start()
cluster = Cluster(['cassandra'], port = 9042, idle_heartbeat_interval=10 )
session = cluster.connect()
session.default_timeout = 100
session.set_keyspace('code2pdf')

@app.route("/")
def hello():
    query = "SELECT * FROM codelang"
    stmnt = SimpleStatement(query)
    res = session.execute(stmnt)
    fin = []
    for r in res:
        fin.append(r.usn)
    return json.dumps(fin)

@app.route("/stats")
def stats():
    return "Hello Moto!"

@app.route("/makePdf",methods=['POST'])
def makepdf():

    details = request.json
    n_urls = details["urlLen"]
    n_files = details["fileLen"]
    usn = details["usn"]
    urlFiles = []
    finals = []

    for i in range(n_urls):
        u = details["urls"][i]
        if u[:11] == "https://raw":
            resp = requests.get(u)
            urlFiles.append(resp.text)

    for i in range(n_files):
        fname = os.path.join("/volume/", details["files"][i])
        lex = lexers.guess_lexer_for_filename(fname,'int a = 3;')
        with open(fname) as f:
            l = f.read()
            
            if usn:
                l = langComment[lex.name.lower()] + " Authored By: " + usn + "\n" + l
            
            x = highlight(l + '  \n' , lexers.get_lexer_by_name(lex.name), formatters.HtmlFormatter(linenos = 'inline', noclasses = True,style="bw"))
            
            finals.append(x)

    for i in range(len(urlFiles)):
        l = urlFiles[i]
        lex = lexers.guess_lexer_for_filename(details["urls"][i].split('/')[-1],'int a = 3;')
        i += 1
        if usn:
            l = langComment[lex.name.split()[0].lower()] + " Authored By: " + usn + "\n" + l

        x = highlight(l + '  \n' , lex, formatters.HtmlFormatter(linenos = 'inline', noclasses = True, style="bw"))

        finals.append(x)

    html = ("<br></br>" * 2).join(finals)

    pdfkit.from_string(html, os.path.join(os.getcwd(), '/volume/') + usn + '.pdf')

    return jsonify({ "pdf":usn + '.pdf'})

if __name__ == "__main__":
    app.run()