import binascii
import io
from kafka import KafkaConsumer

consumer = KafkaConsumer('url',
                         bootstrap_servers=['localhost:9093'])
consumer.subscribe(['url','file'])
for message in consumer:
    # message value and key are raw bytes -- decode if necessary!
    # e.g., for unicode: `message.value.decode('utf-8')`
    print ("%s:%d:%d: key=%s value=%s" % (message.topic, message.partition,
                                          message.offset, message.key,
                                          message.value))

    if message.topic == "url":
        print(binascii.hexlify(message.value).decode('utf-8'))
    
    elif message.topic == "file":
        fileBytes = bytes.fromhex(message.value.decode())
        f = open(message.key, 'wb')
        f.write(fileBytes)
        f.close()
