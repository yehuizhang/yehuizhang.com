# Go Webapp Using GIN

[![GitHub version](https://badge.fury.io/gh/yehuizhang%2Fgo-zyh-webserver.svg)](https://badge.fury.io/gh/yehuizhang%2Fgo-zyh-webserver)
[![CodeFactor](https://www.codefactor.io/repository/github/yehuizhang/go-zyh-webserver/badge)](https://www.codefactor.io/repository/github/yehuizhang/go-zyh-webserver)
[![codecov](https://codecov.io/gh/yehuizhang/go-zyh-webserver/branch/main/graph/badge.svg?token=fQ74xxW1ez)](https://codecov.io/gh/yehuizhang/go-zyh-webserver)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyehuizhang%2Fgo-zyh-webserver.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyehuizhang%2Fgo-zyh-webserver?ref=badge_shield)

## Tasks

- Improve Registry to utilize db.pipelines/HSet

- Registry: Invitation code
  - only allows user to register with a valid invitation code.
  - Implement API to generate invitation code. Optional field: notes
  - When User registered, link invitation code with the notes and then delete the invitation record
  - Add expiration date on invitation code
- Live weather based on user's location

## Design

- [Password Authentication](https://www.sohamkamani.com/golang/password-authentication-and-storage/)

## Quality check

:white_check_mark: :heavy_check_mark: :x: :recycle:
| API | Checked |
| ----------- | ----------- |
| GET /health | :white_check_mark: |
| POST /register | :white_check_mark: |
| POST /login | :white_check_mark: |
| GET /v1/user/info | :white_check_mark: |
| PUT /v1/user/info | :white_check_mark: |

## Tech Stack

- Web Framework
  - [gin](https://github.com/gin-gonic/gin)
    - [Doc](https://gin-gonic.com/)
- Library
  - [validator](https://github.com/go-playground/validator)
    - [Doc](https://pkg.go.dev/github.com/go-playground/validator/v10)
  - Logger
    - [Zap](https://github.com/uber-go/zap)
  - Dependency Injection
    - [dig](https://github.com/uber-go/dig)
    - [wire](https://github.com/google/wire)
  - [go-gin-boilerplate](https://github.com/vsouza/go-gin-boilerplate)
  - [viper-config](https://github.com/spf13/viper)
  - [go.uuid](https://github.com/satori/go.uuid)
  - Testing
    - [Testify](https://github.com/stretchr/testify)
    - [Gin unit test example](https://github.com/yemiwebby/golang-company-api/blob/main/main_test.go)
    - [Github - Building and testing Go](https://docs.github.com/en/actions/automating-builds-and-tests/building-and-testing-go)
- Database
  - [Go-Redis](https://redis.uptrace.dev/)
  - [Go-Redis-Repo](https://github.com/go-redis/redis)
- Authentication
- OAuth2
- CI/CD
  - CircleCI [example](https://github.com/codecov/example-go/blob/main/.circleci/config.yml)

## Sample Curls

Start application

```sh
go run main.go
```

## CodeCov

![CodeCov Graph - Grid](https://codecov.io/gh/yehuizhang/go-zyh-webserver/branch/main/graphs/tree.svg?token=fQ74xxW1ez)

## License

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyehuizhang%2Fgo-zyh-webserver.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyehuizhang%2Fgo-zyh-webserver?ref=badge_large)
