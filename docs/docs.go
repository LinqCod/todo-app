// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "linqcod",
            "email": "linqcod@yandex.ru"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/tasks": {
            "get": {
                "description": "get tasks with filters and pagination",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "get tasks with filters and pagination",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "tasks pagination offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "tasks pagination limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "task assigned date",
                        "name": "date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "status of task",
                        "name": "isCompleted",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task updated successfully",
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.SuccessWithDataResponse"
                        }
                    },
                    "400": {
                        "description": "error bad request data",
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "error while getting filtered tasks",
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "create task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "create task",
                "parameters": [
                    {
                        "description": "Create task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.Task"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Task created successfully",
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "error bad request data",
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "error while creating task",
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tasks/{id}": {
            "get": {
                "description": "get task by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "get task by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task got successfully",
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.SuccessWithDataResponse"
                        }
                    },
                    "400": {
                        "description": "error bad request data",
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "error while getting task by id",
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete task",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "delete task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "error bad request data",
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "error while deleting task",
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "update task",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "update task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.UpdateTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task updated successfully",
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.SuccessWithDataResponse"
                        }
                    },
                    "400": {
                        "description": "error bad request data",
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "error while updating task",
                        "schema": {
                            "$ref": "#/definitions/github_com_linqcod_todo-app_internal_domain.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_linqcod_todo-app_internal_domain.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "github_com_linqcod_todo-app_internal_domain.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "github_com_linqcod_todo-app_internal_domain.SuccessWithDataResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "github_com_linqcod_todo-app_internal_domain.Task": {
            "type": "object",
            "required": [
                "assigned_date",
                "description",
                "title"
            ],
            "properties": {
                "assigned_date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_completed": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "github_com_linqcod_todo-app_internal_domain.UpdateTaskRequest": {
            "type": "object",
            "properties": {
                "assigned_date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_completed": {
                    "type": "boolean"
                },
                "title": {
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
	Title:            "TODO API",
	Description:      "todo service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
