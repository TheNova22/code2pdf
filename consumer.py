import binascii
from kafka import KafkaConsumer

consumer = KafkaConsumer('url',
                         bootstrap_servers=['localhost:9093'])

for message in consumer:
    # message value and key are raw bytes -- decode if necessary!
    # e.g., for unicode: `message.value.decode('utf-8')`
    print ("%s:%d:%d: key=%s value=%s" % (message.topic, message.partition,
                                          message.offset, message.key,
                                          message.value))
    print(binascii.unhexlify(message.value).decode('utf-8'))