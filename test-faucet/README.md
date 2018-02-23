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
    go run app.go
```
