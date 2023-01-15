# Go Webapp Using GIN

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

## Features

- Web Framework
  - [gin](https://github.com/gin-gonic/gin)
  - [validator](https://github.com/go-playground/validator)
    - [Doc](https://pkg.go.dev/github.com/go-playground/validator/v10)
  - [go-gin-boilerplate](https://github.com/vsouza/go-gin-boilerplate)
  - [viper-config](https://github.com/spf13/viper)
  - [go.uuid](https://github.com/satori/go.uuid)
- Database
  - [Go-Redis](https://redis.uptrace.dev/)
  - [Go-Redis-Repo](https://github.com/go-redis/redis)
- Authentication
- OAuth2

## Sample Curls

Start application

```sh
go run main.go
```

Register new user account

```sh
curl --location --request POST 'http://localhost:8080/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "valerie",
    "password": "1234567"
}'

{
    "user_id": "b8b0249c-8707-4b1b-b897-5739642fbc27",
    "username": "valerie",
    "password": "1234567",
    "active": true,
    "created_at": 1672741900752807800,
    "updated_at": 1672741900752807800
}
```

updateUserInfo

```sh
curl --location --request PUT 'http://localhost:8080/v1/user/info' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "1d541f1d-3e09-49d2-bcbc-ebbf38fb327c",
    "name": "Yehui Zhang",
    "birthday": "2021-01-01"
}'

{
    "name": "Yehui Zhang",
    "birthday": "2021-01-01",
    "updated_at": 1672742751180898700
}
```

login

```sh
curl --location --request POST 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "yehui",
    "password": "123456"
}'

{
    "error": "incorrect password",
    "message": "Error to retrieve userCredential"
}
```
