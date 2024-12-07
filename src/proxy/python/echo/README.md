# Echo Server

## Build

```bash
uv sync
source .venv/bin/activate
```

## Run

### REST API

```bash
(.venv) fastapi run
```

```bash
curl localhost:8000

{"Hello":"World"}
```

```bash
curl 'localhost:8000/items/1?q=foo'

{"item_id":1,"q":"foo"}
```

