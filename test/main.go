package main

import "github.com/gin-gonic/gin"

func main() {
	g := gin.Default()
	g.GET("/api/doc.json", func(ctx *gin.Context) {
		ctx.String(200, `{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {},
  "paths": {
    "/api/v1/addressbook/dept/list": {
      "post": {
        "description": "获取部门树形结构",
        "tags": [
          "addressbook"
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/ListReq"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "json response",
            "schema": {
              "$ref": "#/definitions/DepartmentTree"
            },
            "examples": {
              "application/json": {
                "children": null
              }
            }
          },
          "default": {
            "description": "json response",
            "schema": {
              "$ref": "#/definitions/DepartmentTree"
            },
            "examples": {
              "application/json": {
                "children": null
              }
            }
          }
        }
      }
    },
    "/api/v1/addressbook/sync": {
      "post": {
        "description": "同步组织架构,从企微员工部门信息",
        "tags": [
          "addressbook"
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "json response",
            "schema": {
              "$ref": "#/definitions/Response"
            },
            "examples": {
              "application/json": {
                "data": {},
                "code": 0,
                "message": "",
                "stack": null
              }
            }
          },
          "default": {
            "description": "json response",
            "schema": {
              "$ref": "#/definitions/Response"
            },
            "examples": {
              "application/json": {
                "data": {},
                "code": 0,
                "message": "",
                "stack": null
              }
            }
          }
        }
      }
    },
    "/api/v1/auth/info": {
      "get": {
        "description": "获取当前账号登录信息",
        "tags": [
          "auth"
        ],
        "responses": {
          "200": {
            "description": "json response",
            "schema": {
              "$ref": "#/definitions/Response"
            },
            "examples": {
              "application/json": {
                "data": {
                  "id": "",
                  "username": "",
                  "avatar": "",
                  "nickname": "",
                  "permissions": null
                },
                "code": 0,
                "message": "",
                "stack": null
              }
            }
          },
          "default": {
            "description": "json response",
            "schema": {
              "$ref": "#/definitions/Response"
            },
            "examples": {
              "application/json": {
                "data": {
                  "id": "",
                  "username": "",
                  "avatar": "",
                  "nickname": "",
                  "permissions": null
                },
                "code": 0,
                "message": "",
                "stack": null
              }
            }
          }
        }
      }
    },
    "/api/v1/auth/login": {
      "post": {
        "description": "登录接口",
        "tags": [
          "auth"
        ],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/LoginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "json response",
            "schema": {
              "$ref": "#/definitions/Response"
            },
            "examples": {
              "application/json": {
                "data": "session id",
                "code": 0,
                "message": "",
                "stack": null
              }
            }
          },
          "default": {
            "description": "json response",
            "schema": {
              "$ref": "#/definitions/Response"
            },
            "examples": {
              "application/json": {
                "data": "session id",
                "code": 0,
                "message": "",
                "stack": null
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "": {
      "type": "object"
    },
    "DepartmentTree": {
      "type": "object",
      "properties": {
        "children": {
          "type": "object",
          "$ref": "#/definitions/DepartmentTree"
        },
        "id": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "nameEn": {
          "type": "string"
        },
        "order": {
          "type": "integer"
        },
        "parentId": {
          "type": "integer"
        }
      }
    },
    "ListReq": {
      "type": "object",
      "properties": {
        "asc": {
          "type": "boolean"
        },
        "keyword": {
          "type": "string"
        },
        "order": {
          "type": "string"
        },
        "page": {
          "type": "integer"
        },
        "pageSize": {
          "type": "integer"
        }
      }
    },
    "LoginRequest": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "password": {
          "description": "密码: 开发环境是 admin",
          "type": "string"
        },
        "username": {
          "description": "账号，开发环境是 admin",
          "type": "string"
        }
      }
    },
    "Response": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "data": {
          "type": "object"
        },
        "message": {
          "type": "string"
        },
        "stack": {
          "type": "object"
        }
      }
    }
  },
  "securityDefinitions": {
    "ApiKeyAuth": {
      "description": "apiKey auth",
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "ApiKeyAuth": []
    }
  ]
}`)
	})
	g.Run(":8080")
}
