from cassandra.cluster import Cluster

cluster = Cluster(['cassandra'], port = 9042)

session = cluster.connect()
session.set_keyspace('code2pdf')
x = session.execute("DESC tables")

for r in x:
    print("Hello " + str(r))