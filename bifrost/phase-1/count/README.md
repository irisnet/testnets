# Bifrost-1 Tasks Count

Add your account address and serial number to `participants.json`

```json
[
    ...
    {
        "github": "dolphintwo",
        "pgp": "7A38C8CD6C0FA32A",
        "addr": "iaa18p5jgdfw6dh0ftwg72ezxj4qfffn6l4jnzqrnj",
        "serial": 1
    },
    ...
]
```

Run this cmd in project root dir

```bash
go run main.go
```

You can see the result in `result.json`

```json
[
    ...
    {
        "github": "dolphintwo",
        "pgp": "7A38C8CD6C0FA32A",
        "addr": "iaa18p5jgdfw6dh0ftwg72ezxj4qfffn6l4jnzqrnj",
        "serial": 1,
        "tasks": [
            true,
            true,
            true,
            true,
            true,
            true,
            true,
            true,
            true,
            true,
            true,
            true,
            true,
            true,
            true
        ]
    },
    ...
]
```
