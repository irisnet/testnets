# Bifrost-2 Tasks Count

## Verify

Add `github` and `pgp` to  `participants_raw.json`.

```json
[
    ...
    {
        "github": "dolphintwo",
        "pgp": "7A38C8CD6C0FA32A"
    },
    ...
]
```

There are restrictions on github's api access, please keep only your own information.

Verify GitHub and Keybase.

```bash
go run ./verify/main.go
```

If you are eligible, `verified` will be `true` in `participants_verified.json`.

```json
[
    ...
    {
        "github": "dolphintwo",
        "pgp": "7A38C8CD6C0FA32A",
        "verified": true,
        "reason": "",
        ...
    },
    ...
]
```

## Get Addresses

Get `addr` and `val_addr`.

```bash
go run ./tools/main.go
```

Addresses will be set in `participants.json`.

```json
[
    ...    
    {
        "github": "dolphintwo",
        "pgp": "7A38C8CD6C0FA32A",
        "verified": true,
        "reason": "",
        "addr": "iaa18p5jgdfw6dh0ftwg72ezxj4qfffn6l4jnzqrnj",
        "val_addr": "740A34638BF6603FC6B7F8AE46FD8F4FE6C782E5",
        ...
    },
    ...
]
```

## Task statistics

Count task results.

```bash
go run ./mian.go
```

Task will be true if you complete successfully.

```json
[
    ...
    {
        "github": "dolphintwo",
        "pgp": "7A38C8CD6C0FA32A",
        "verified": true,
        "reason": "",
        "addr": "iaa18p5jgdfw6dh0ftwg72ezxj4qfffn6l4jnzqrnj",
        "val_addr": "740A34638BF6603FC6B7F8AE46FD8F4FE6C782E5",
        "tasks": {
            "task1": true,
            "task2": true,
            "task3": true
        },
        "badge": {
            "sliver": 1,
            "bronze": 2
        }
    },
    ...
]
```
