// Code generated by swaggo/swag. DO NOT EDIT
package docs

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
        "/users/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Query users list",
                "parameters": [
                    {
                        "enum": [
                            "Reader",
                            "Author",
                            "Admin"
                        ],
                        "type": "string",
                        "description": "user kind, default is all",
                        "name": "kind",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserListResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ResponseMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ResponseMessage": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "200,0,1: success; others: failed",
                    "type": "integer"
                },
                "data": {
                    "description": "omit in failed response"
                },
                "message": {
                    "description": "omit in success response",
                    "type": "string"
                }
            }
        },
        "dto.UserListResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object",
                    "properties": {
                        "list": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.User"
                            }
                        },
                        "total": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "kind": {
                    "type": "string",
                    "enum": [
                        "Reader",
                        "Author",
                        "Admin"
                    ]
                },
                "nickname": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
