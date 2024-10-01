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
        "/customers": {
            "post": {
                "description": "Creates customers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Creates new customers",
                "parameters": [
                    {
                        "description": "Creates customer",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CustomerDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/http.HttpCustomerCreated"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.HttpResponse"
                        }
                    }
                }
            }
        },
        "/customers/{dni}": {
            "get": {
                "description": "Through a get request the customer dni is sent to http client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Get customers from db using the system service layer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer DNI",
                        "name": "dni",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Customer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/http.HttpResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Customer": {
            "type": "object",
            "properties": {
                "cashbackCollected": {
                    "type": "number"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "pointsCollected": {
                    "type": "number"
                },
                "userDNI": {
                    "type": "string"
                }
            }
        },
        "dto.CustomerDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "user_dni": {
                    "type": "string"
                }
            }
        },
        "http.HttpCustomerCreated": {
            "type": "object",
            "properties": {
                "customer": {
                    "$ref": "#/definitions/dto.CustomerDTO"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "http.HttpResponse": {
            "type": "object",
            "properties": {
                "response": {
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
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "HTTP Client for Loyalty System",
	Description:      "Client for handling http requests for loyalty system",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
