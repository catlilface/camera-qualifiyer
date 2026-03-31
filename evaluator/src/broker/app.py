from faststream import FastStream
from faststream.rabbit import RabbitBroker

broker = RabbitBroker()
app = FastStream(broker)


@broker.subscriber("routing_key")  # handle messages by routing key
async def handle(msg):
    print(msg)