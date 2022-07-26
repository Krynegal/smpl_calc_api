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
        "contact": {},
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/add": {
            "post": {
                "description": "return sum of two numbers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "addition",
                "parameters": [
                    {
                        "description": "comment",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/smpl_calc_api.Data"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/div": {
            "post": {
                "description": "return division of two numbers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "division",
                "parameters": [
                    {
                        "description": "comment",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/smpl_calc_api.Data"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/mul": {
            "post": {
                "description": "return multiplication of two numbers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "multiplication",
                "parameters": [
                    {
                        "description": "comment",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/smpl_calc_api.Data"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/sub": {
            "post": {
                "description": "return subtraction of two numbers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "subtraction",
                "parameters": [
                    {
                        "description": "comment",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/smpl_calc_api.Data"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "handlers.Response": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "integer"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "smpl_calc_api.Data": {
            "type": "object",
            "properties": {
                "operand1": {
                    "$ref": "#/definitions/smpl_calc_api.Operand"
                },
                "operand2": {
                    "$ref": "#/definitions/smpl_calc_api.Operand"
                },
                "toBase": {
                    "type": "integer"
                }
            }
        },
        "smpl_calc_api.Operand": {
            "type": "object",
            "properties": {
                "base": {
                    "type": "integer",
                    "example": 10
                },
                "value": {
                    "type": "string"
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
	Schemes:          []string{},
	Title:            "Simple Calculator API",
	Description:      "Description",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
