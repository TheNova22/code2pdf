from datetime import datetime
import json
from  pygments import lexers
from cassandra.cluster import Cluster
from cassandra.util import max_uuid_from_time
from kafka import KafkaConsumer


cluster = Cluster(['cassandra'], port = 9042)
session = cluster.connect()
session.set_keyspace('code2pdf')
consumer = KafkaConsumer('url',
                         bootstrap_servers=['kafka:29092'])

consumer.subscribe(['url'])

for message in consumer:
    m = message.value.decode('utf-8')
    print(m)
    d = json.loads(m)
    print(d)
    lex = lexers.guess_lexer_for_filename(d["url"].split('/')[-1],'int a = 3;')
    session.execute("""INSERT INTO codelang (usn, tid, lang) values (%(usn)s, %(tid)s, %(lang)s)""", {"usn" : d["usn"], "tid" : max_uuid_from_time(datetime.now()), "lang" : lex.name})
    # x = session.execute("INSERT INTO codelang (usn, tid, lang) values({0},{1},{2})".format(d["usn"],max_uuid_from_time(datetime.now()), lex.name))
# x = session.execute("DESC tables")

# for r in x:
#     print("Hello " + str(r))