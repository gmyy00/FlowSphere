// Package logger 提供日志功能的封装
// 基于 zerolog 库，支持结构化日志输出
package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// log 是全局日志实例
var log zerolog.Logger

// init 初始化日志输出格式为控制台格式
func init() {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}
	log = zerolog.New(output).With().Timestamp().Logger()
}

// Init 设置全局日志级别
// 支持的级别: debug, info, warn, error
func Init(level string) {
	switch level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

// 全局日志
func Get() zerolog.Logger {
	return log
}

// debug 级别
func Debug() *zerolog.Event {
	return log.Debug()
}

// info 级别
func Info() *zerolog.Event {
	return log.Info()
}

// warn 级别
func Warn() *zerolog.Event {
	return log.Warn()
}

// error 级别
func Error() *zerolog.Event {
	return log.Error()
}

// fatal 级别
func Fatal() *zerolog.Event {
	return log.Fatal()
}

// With 返回带有上下文的日志构建器
func With() zerolog.Context {
	return log.With()
}
