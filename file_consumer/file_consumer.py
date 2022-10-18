from datetime import datetime
import json
from  pygments import lexers
from cassandra.cluster import Cluster
from cassandra.util import max_uuid_from_time
from kafka import KafkaConsumer
import os

cluster = Cluster(['cassandra'], port = 9042, idle_heartbeat_interval=10 )
session = cluster.connect()
session.default_timeout = 100
session.set_keyspace('code2pdf')
consumer = KafkaConsumer('file',
                         bootstrap_servers=['kafka:29092'])

consumer.subscribe(['file'])

for message in consumer:
    m = message.value.decode('utf-8')
    d = json.loads(m)
    print(d)
    lex = lexers.guess_lexer_for_filename(d["file"],'int a = 3;')
    session.execute("""INSERT INTO codelang (usn, tid, lang) values (%(usn)s, %(tid)s, %(lang)s)""", {"usn" : d["usn"], "tid" : max_uuid_from_time(datetime.now()), "lang" : lex.name})
    os.remove("/volume/" + d["file"])