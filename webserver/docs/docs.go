// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://yehuizhang.com/terms/",
        "contact": {
            "name": "Yehui Zhang",
            "url": "http://www.yehuizhang.com/support",
            "email": "yehuizhang@yehuizhang.com"
        },
        "license": {
            "name": "MIT License",
            "url": "https://github.com/yehuizhang/go-zyh-webserver/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "Get Health",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Get health status of the server",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register new account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register new account",
                "parameters": [
                    {
                        "description": "account credential",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/account.SignUpForm"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "description": "Get User Info",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User Info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/info.UserInfo"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "account.SignUpForm": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 6
                },
                "username": {
                    "description": "This should be improved by using custom validator",
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 3
                }
            }
        },
        "info.UserInfo": {
            "type": "object",
            "properties": {
                "birthday": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "photoURL": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Description for what is this security definition being used",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "zyh-go-webserver",
	Description:      "This is the backend server for yehuizhang.com",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
