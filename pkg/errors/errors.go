// Package errors 定义了应用程序的错误类型
// 提供统一的错误格式和预定义的错误常量
package errors

import (
	"net/http"
)

type AppError struct {
	Code    string `json:"code"`    // 错误码，用于前端识别
	Message string `json:"message"` // 错误描述信息
	Status  int    `json:"-"`       // HTTP 状态码，不序列化
}

func (e *AppError) Error() string {
	return e.Message
}

// 创建AppError实例
func New(code string, message string, status int) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Status:  status,
	}
}

// ErrInvalidParameter 创建参数无效错误 (400)
func ErrInvalidParameter(msg string) *AppError {
	return New("INVALID_PARAMETER", msg, http.StatusBadRequest)
}

// ErrConflict 创建资源冲突错误 (400)
func ErrConflict(msg string) *AppError {
	return New("CONFLICT", msg, http.StatusBadRequest)
}

// ErrInvalidOperation 创建无效操作错误 (400)
func ErrInvalidOperation(msg string) *AppError {
	return New("INVALID_OPERATION", msg, http.StatusBadRequest)
}

// ErrUnauthorized 创建未认证错误 (401)
func ErrUnauthorized(msg string) *AppError {
	return New("UNAUTHORIZED", msg, http.StatusUnauthorized)
}

// ErrAuthFailed 创建认证失败错误 (401)
func ErrAuthFailed(msg string) *AppError {
	return New("AUTH_FAILED", msg, http.StatusUnauthorized)
}

// ErrInvalidToken 创建无效 Token 错误 (401)
func ErrInvalidToken(msg string) *AppError {
	return New("INVALID_TOKEN", msg, http.StatusUnauthorized)
}

// ErrForbidden 创建无权限错误 (403)
func ErrForbidden(msg string) *AppError {
	return New("FORBIDDEN", msg, http.StatusForbidden)
}

// ErrNotFound 创建资源不存在错误 (404)
func ErrNotFound(code string, msg string) *AppError {
	return New(code, msg, http.StatusNotFound)
}

// ErrInternal 创建服务器内部错误 (500)
func ErrInternal(msg string) *AppError {
	return New("INTERNAL_ERROR", msg, http.StatusInternalServerError)
}

// 预定义的常用错误
var (
	// ErrUserNotFound 用户不存在
	ErrUserNotFound = New("USER_NOT_FOUND", "user not found", http.StatusNotFound)
	// ErrProjectNotFound 项目不存在
	ErrProjectNotFound = New("PROJECT_NOT_FOUND", "project not found", http.StatusNotFound)
	// ErrTaskNotFound 任务不存在
	ErrTaskNotFound = New("TASK_NOT_FOUND", "task not found", http.StatusNotFound)
)
