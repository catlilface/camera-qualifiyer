import os

from dotenv import load_dotenv

load_dotenv()


PENDING_QUEUE = os.environ.get("PENDING_QUEUE", "evaluation.pending")
PROCESS_QUEUE = os.environ.get("PROCESS_QUEUE", "evaluation.process")
RABBITMQ_PORT = int(os.environ.get("RABBITMQ_PORT_AMQP", 5672))
