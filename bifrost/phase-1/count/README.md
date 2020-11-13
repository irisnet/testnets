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

Run those cmds to count tasks

```bash
cd <your-path>/testnets/bifrost/phase-1/count

go mod tidy

go run main.go
```

You can see the result in `result.json`

If you have completed a task, it will be `true`

```json
[
    ...
    {
        "github": "dolphintwo",
        "pgp": "7A38C8CD6C0FA32A",
        "addr": "iaa18p5jgdfw6dh0ftwg72ezxj4qfffn6l4jnzqrnj",
        "serial": 1,
        "tasks": {
            "service_tasks": {
                "task1": true,
                "task2": true,
                "task3": true,
                "task4": true,
                "task5": true,
                "task6": true,
                "task7": true
            },
            "record_tasks": {
                "task1": true
            },
            "nft_tasks": {
                "task1": true,
                "task2": true,
                "task3": true,
                "task4": true,
                "task5": true
            },
            "random_tasks": {
                "task1": true,
                "task2": true
            }
        }
    },
    ...
]
```
