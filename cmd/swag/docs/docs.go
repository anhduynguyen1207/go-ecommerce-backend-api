// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "github.com/anhduynguyen1207/go-ecommerce-backend-go",
        "contact": {
            "name": "TEAM TIPSGO",
            "url": "github.com/anhduynguyen1207/go-ecommerce-backend-go",
            "email": "anhduynguyen1@gmail.com"
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
        "/user/login": {
            "post": {
                "description": "User Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account management"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "When user is register send otp to email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account management"
                ],
                "summary": "User Register",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/user/two-factor/setup": {
            "post": {
                "description": "User Setup Two Factor Authentication",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account 2fa"
                ],
                "summary": "User Setup Two Factor Authentication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SetupTwoFactorAuthInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/user/two-factor/verify": {
            "post": {
                "description": "User verify Two Factor Authentication",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account 2fa"
                ],
                "summary": "User verify Two Factor Authentication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TwoFactorVerificationInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/user/verify_account": {
            "post": {
                "description": "Verify OTP Login By User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account management"
                ],
                "summary": "Verify OTP Login By User",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.VerifyInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        },
        "/user/verify_update_pass_register": {
            "post": {
                "description": "UpdatePasswordRegister",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account management"
                ],
                "summary": "UpdatePasswordRegister",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdatePasswordRegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponseData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.LoginInput": {
            "type": "object",
            "properties": {
                "user_account": {
                    "type": "string"
                },
                "user_password": {
                    "type": "string"
                }
            }
        },
        "model.RegisterInput": {
            "type": "object",
            "properties": {
                "verify_key": {
                    "type": "string"
                },
                "verify_purpose": {
                    "type": "string"
                },
                "verify_type": {
                    "type": "integer"
                }
            }
        },
        "model.SetupTwoFactorAuthInput": {
            "type": "object",
            "properties": {
                "two_factor_auth_type": {
                    "type": "string"
                },
                "two_factor_email": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.TwoFactorVerificationInput": {
            "type": "object",
            "properties": {
                "two_factor_code": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.UpdatePasswordRegisterInput": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "user_token": {
                    "type": "string"
                }
            }
        },
        "model.VerifyInput": {
            "type": "object",
            "properties": {
                "verify_code": {
                    "type": "string"
                },
                "verify_key": {
                    "type": "string"
                }
            }
        },
        "response.ErrorResponseData": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "mã trạng thái trả về",
                    "type": "integer"
                },
                "detail": {
                    "description": "dữ liệu trả về"
                },
                "message": {
                    "description": "thông điệp trả về",
                    "type": "string"
                }
            }
        },
        "response.ResponseData": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "mã trạng thái trả về",
                    "type": "integer"
                },
                "data": {
                    "description": "dữ liệu trả về"
                },
                "message": {
                    "description": "thông điệp trả về",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8002",
	BasePath:         "/v1/2024",
	Schemes:          []string{},
	Title:            "API Documentation Ecommerce Backend SHOPDEVGO",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
