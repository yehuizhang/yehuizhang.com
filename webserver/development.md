# Development Environment Setup

## Start Up

Containers

```shell
# Start
docker compose -p "zyh-gin-webapp" --profile dev up --build --detach
```

```shell
# Stop
docker compose -p "zyh-gin-webapp" down
```

Application

Dev: Local DB
```shell
go run main.go wire_gen.go
```

Dev: Remote DB
```shell
go run main.go wire_gen.go -configName=".env.dev"
```

## Code Generating

### Wire - Dependency Injection

```shell
# Generate wire; run in root level 
wire
```

### [Mockery](https://github.com/vektra/mockery)

To generate mocks based on the interface, update `--dir=` section and run the command below

```shell
mockery --all --output=./test --outpkg=test --dir=pkg/dao
```

## Tools

### [Swagger](http://localhost:8080/swagger/index.html)

```shell
# Generate/Update swagger doc
swag init
```

```shell
# Format comments
swag fmt
```

###  

### Development

[GoLand-Recommended](https://www.jetbrains.com/go/)

- Can be used to access Databases

### Database

#### Adminer(Postgres)

1. Go to http://localhost:8090/
2. Server: `{container_name}:{port}` in this case `postgres:5432`
3. Username/Password/Database: Check `.env`

#### [DBeaver](https://dbeaver.io/)

Used for Postgres

#### [Regis-Insight](https://redis.com/redis-enterprise/redis-insight/)

Used to access Redis container

[download Mac-M1 latest](https://download.redisinsight.redis.com/latest/RedisInsight-v2-mac-arm64.dmg)


### Docker
```shell
docker build -t zyh-webserver .
docker run -it --rm --name zyh-webserver-instance zyh-webserver
```