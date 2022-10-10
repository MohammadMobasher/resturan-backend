// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/foodgroup": {
            "post": {
                "description": "create a food group",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "food group"
                ],
                "summary": "create a food group",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/user": {
            "get": {
                "description": "get a user",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "get a user",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "put": {
                "description": "put a user",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "put a user",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "description": "create a user",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "create a user",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/user/:userId": {
            "delete": {
                "description": "delete a user",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "delete a user",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/users": {
            "get": {
                "description": "Get all users",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get all users.",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v2/foodgroup": {
            "post": {
                "description": "create a food group",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "food group"
                ],
                "summary": "create a food group",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v2/foodgroup/:foodgroupId": {
            "delete": {
                "description": "delete a food group",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "food group"
                ],
                "summary": "delete a food group",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Gin Swagger Example API",
	Description:      "resturan",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
