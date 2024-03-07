// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/currency/save/{date}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kmf"
                ],
                "summary": "По указанной дате сохранять курсы валют в БД",
                "parameters": [
                    {
                        "type": "string",
                        "description": "date",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "$ref": "#/definitions/handlers.saveResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/currency/save/{date}/{code}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "kmf"
                ],
                "summary": "По указанной дате показать курсы валют из БД",
                "parameters": [
                    {
                        "type": "string",
                        "description": "date",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "response",
                        "schema": {
                            "$ref": "#/definitions/handlers.rateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.rateItem": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "handlers.rateResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.rateItem"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "handlers.saveResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "Nat Service",
	Description:      "Description of the service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
