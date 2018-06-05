# test-faucet
faucet for iris testnet

### RUN
install PostgreSQL or by docker,
```
    docker-compose docker/postgre.yml up -d
```
And initialize database with config/faucet.sql.

Modify config.yml for yourself, then:

```
    # if window
    go build -o build/iris-faucet.exe app.go
    #else 
    go build -o build/iris-faucet app.go
```

if you will run it in docker,you will
```
    docker build -t faucet ./
```

docker/docker-compose.yml is an example for docker-compose