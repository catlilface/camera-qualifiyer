import asyncio

from faststream import FastStream
from faststream.rabbit import RabbitBroker

from src.env import PENDING_QUEUE, PROCESS_QUEUE, RABBITMQ_PORT

broker = RabbitBroker(f"amqp://guest:guest@rabbitmq:{RABBITMQ_PORT}")
app = FastStream(broker)


@broker.subscriber(PENDING_QUEUE)
async def pending(): ...


@broker.subscriber(PROCESS_QUEUE)
async def process(): ...


if __name__ == "__main__":
    asyncio.run(app.run())
