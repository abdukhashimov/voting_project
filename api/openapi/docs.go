// Package openapi GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package openapi

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/boards": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "boards"
                ],
                "summary": "API to create boards",
                "parameters": [
                    {
                        "description": "payload-body",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.BoardCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SuccessResp"
                        }
                    }
                }
            }
        },
        "/boards/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "boards"
                ],
                "summary": "API to get a single board",
                "parameters": [
                    {
                        "type": "string",
                        "description": "key-id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Board"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "API to get list of users",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "district_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "region_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.UserAllResp"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "API to create users",
                "parameters": [
                    {
                        "description": "users-body",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SuccessResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResp"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "API to get a single user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "key-id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "API to update users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "paylaod",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SuccessResp"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "API to delete a single user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "key-id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SuccessResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Board": {
            "type": "object",
            "properties": {
                "accept_end_date": {
                    "type": "string"
                },
                "accept_start_date": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "review_end_date": {
                    "type": "string"
                },
                "review_start_date": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "total_amount": {
                    "type": "integer"
                },
                "voting_end_date": {
                    "type": "string"
                },
                "voting_start_date": {
                    "type": "string"
                }
            }
        },
        "domain.BoardCreate": {
            "type": "object",
            "properties": {
                "accept_end_date": {
                    "type": "string"
                },
                "accept_start_date": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "review_end_date": {
                    "type": "string"
                },
                "review_start_date": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "total_amount": {
                    "type": "integer"
                },
                "voting_end_date": {
                    "type": "string"
                },
                "voting_start_date": {
                    "type": "string"
                }
            }
        },
        "domain.ErrorResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.SuccessResp": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "domain.User": {
            "type": "object",
            "required": [
                "district_id",
                "fullname",
                "phone_number",
                "region_id"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "district_id": {
                    "type": "integer"
                },
                "fullname": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "region_id": {
                    "type": "integer"
                },
                "role": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.UserAllResp": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.User"
                    }
                }
            }
        },
        "domain.UserCreate": {
            "type": "object",
            "required": [
                "district_id",
                "fullname",
                "phone_number",
                "region_id"
            ],
            "properties": {
                "district_id": {
                    "type": "integer"
                },
                "fullname": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "region_id": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "UsersAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Open budget API",
	Description:      "This API contains the source for the Open budget app",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
