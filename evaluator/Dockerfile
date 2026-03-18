FROM ghcr.io/astral-sh/uv:python3.13-alpine

WORKDIR /app

COPY ./pyproject.toml ./uv.lock ./main.py /app/
COPY ./src /app/src

RUN uv sync --no-dev

CMD ["uv", "run", "main.py"]
