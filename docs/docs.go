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
        "contact": {
            "name": "Maintainer Shi Yue",
            "email": "jsclndnz@gmail.com"
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
        "/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.MessageModel"
                        }
                    }
                }
            }
        },
        "/session": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Session"
                ],
                "summary": "Get The Last Session",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Session"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.MessageModel"
                        }
                    }
                }
            }
        },
        "/sessions": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Session"
                ],
                "summary": "List Sessions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Session"
                            }
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
                    "Session"
                ],
                "summary": "Add A Session",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.SessionAdd"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Session"
                        }
                    },
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.Session"
                        }
                    }
                }
            }
        },
        "/sessions/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Session"
                ],
                "summary": "Get Session",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Session"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.MessageModel"
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
                    "User"
                ],
                "summary": "List Users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.User"
                            }
                        }
                    }
                }
            }
        },
        "/users/{name}": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Add A User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.MessageModel": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "main.Motion": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "records": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Record"
                    }
                },
                "sessionID": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "main.Record": {
            "type": "object",
            "required": [
                "vote"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "motion": {
                    "$ref": "#/definitions/main.Motion"
                },
                "motionID": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/main.User"
                },
                "userID": {
                    "type": "integer"
                },
                "vote": {
                    "type": "integer"
                }
            }
        },
        "main.Session": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "motions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Motion"
                    }
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "main.SessionAdd": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "main.User": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Voter",
	Description:      "voter backend",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
