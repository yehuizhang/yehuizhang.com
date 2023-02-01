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
mockery --dir=src/dao/user/account --all --output=./test --outpkg=test
```
