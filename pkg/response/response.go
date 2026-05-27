// Package response 提供 HTTP 响应的统一格式封装
// 基于 gin 框架，提供成功和错误响应的标准化输出
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gmyy00/flowsphere/pkg/errors"
)

// Response 是统一的 HTTP 响应结构
type Response struct {
	Data  any        `json:"data,omitempty"`  // 响应数据，成功时返回
	Error *ErrorBody `json:"error,omitempty"` // 错误信息，失败时返回
}

// ErrorBody 是错误响应的详细信息
type ErrorBody struct {
	Code    string `json:"code"`    // 错误码
	Message string `json:"message"` // 错误描述
}

// Success 返回 200 成功响应
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Data: data,
	})
}

// Created 返回 201 创建成功响应
func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, Response{
		Data: data,
	})
}

// Error 返回错误响应，状态码由 AppError 决定
func Error(c *gin.Context, err *errors.AppError) {
	c.JSON(err.Status, Response{
		Error: &ErrorBody{
			Code:    err.Code,
			Message: err.Message,
		},
	})
}

// InternalError 返回 500 内部服务器错误
func InternalError(c *gin.Context, msg string) {
	Error(c, errors.ErrInternal(msg))
}

// BadRequest 返回 400 参数错误
func BadRequest(c *gin.Context, msg string) {
	Error(c, errors.ErrInvalidParameter(msg))
}

// Unauthorized 返回 401 未认证错误
func Unauthorized(c *gin.Context, msg string) {
	Error(c, errors.ErrUnauthorized(msg))
}

// Forbidden 返回 403 无权限错误
func Forbidden(c *gin.Context, msg string) {
	Error(c, errors.ErrForbidden(msg))
}

// NotFound 返回 404 资源不存在错误
func NotFound(c *gin.Context, code string, msg string) {
	Error(c, errors.ErrNotFound(code, msg))
}
