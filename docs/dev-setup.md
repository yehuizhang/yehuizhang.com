# Development Environment Setup

## Containers

```shell
# Start
docker-compose -p "zyh-gin-webapp" up --build --detach
```

```shell
# Stop
docker-compose -p "zyh-gin-webapp" down
```

## Application

```shell
go run main.go wire_gen.go
```

[Swagger](http://localhost:8080/swagger/index.html)

Tools to access DB:

- [GoLand-Recommended](https://www.jetbrains.com/go/)
- If you don't have GoLand
  - [DBeaver](https://dbeaver.io/)
  - [Regis-Insight](https://redis.com/redis-enterprise/redis-insight/)
      - [download Mac-M1 latest](https://download.redisinsight.redis.com/latest/RedisInsight-v2-mac-arm64.dmg)

## Testing

### [Mockery](https://github.com/vektra/mockery)
To generate mocks based on the interface, update `--dir=` section and run the command below
```shell
mockery --all --output=./test --outpkg=test --dir=pkg/dao
```

## Swagger

```shell
# Generate/Update swagger doc
swag init
```
```shell
# Format comments
swag fmt
```

## Wire
```shell
# Generate wire; run in root level 
wire
```