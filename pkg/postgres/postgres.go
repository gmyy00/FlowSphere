// 提供 PostgreSQL 数据库连接的封装

package postgres

import (
	"fmt"
	"time"

	"github.com/gmyy00/flowsphere/pkg/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PostgresSQL 连接配置
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string // SSL 模式，默认为 disable
}

func New(cfg Config) (*gorm.DB, error) {
	// 设置SSL模式
	sslMode := cfg.SSLMode
	if sslMode == "" {
		sslMode = "disable"
	}

	// 构建DSN(Data Source Name)连接字符串
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, sslMode,
	)

	// 打开数据库连接
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// logger.Default打开gorm的日志模式, 并且LogMode(logger.Info)设定此条日志等级为info
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// 获取底层数据库连接实例
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	// 配置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大存活时间

	return db, nil
}
