# Setup Application

Start Docker containers for ZooKeeper and Kafka
```
make start
```

Create required topics: balance-table, deposit, flag_wallet, flagger-table, threshold-table
```
make topics
```

Start endpoint handlers, emitter and view
The service offers only two endpoints
```
make dev
```

Start collector processor, detector processor, and flagger processor
```
make processor
```

# Test Service
Check wallet deposit
```
curl localhost:8080/check/0x19AbNdzoUYXV6kpJwNbFXneY9QtsTJjtTm
```

Deposit money to wallet
```
curl -X POST -d '{"wallet_id": "0x19AbNdzoUYXV6kpJwNbFXneY9QtsTJjtTm", "amount": 2000}' http://localhost:8080/deposit
```

To add or remove a wallet flag manually from the flagger table, we can use the command line tool cmd/flag-wallet.

Use this to enable wallet flag
```
go run cmd/flag-wallet/main.go -wallet 0x19AbNdzoUYXV6kpJwNbFXneY9QtsTJjtTm
```
Use this to remove wallet flag
```
go run cmd/flag-wallet/main.go -wallet 0x19AbNdzoUYXV6kpJwNbFXneY9QtsTJjtTm -remove
```

# Optional Command
View list of created topics
```
docker-compose exec broker kafka-topics --list --bootstrap-server localhost:9092
```

Delete existing topic
```
docker-compose exec broker kafka-topics --bootstrap-server localhost:9092 --delete --topic [topic name]
```