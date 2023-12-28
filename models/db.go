package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbEnv struct {
	Host     string
	User     string
	Pass     string
	DB       string
	Port     string
	SslMode  string
	TimeZone string
	Dsn      string
}

func (env *DbEnv) InitDB() error {
	if env.Host == "" {
		if os.Getenv("DATABASE_HOST") != "" {
			env.Host = os.Getenv("DATABASE_HOST") // host=コンテナ名
		} else {
			return errors.New("host is not found")
		}
	}
	if env.User == "" {
		if os.Getenv("DATABASE_USER") != "" {
			env.User = os.Getenv("DATABASE_USER")
		} else {
			return errors.New("POSTGRES_USER is not found")
		}
	}
	if env.Pass == "" {
		if os.Getenv("DATABASE_PASSWORD") != "" {
			env.Pass = os.Getenv("DATABASE_PASSWORD")
		} else {
			return errors.New("POSTGRES_PASSWORD is not found")
		}
	}
	if env.DB == "" {
		if os.Getenv("DATABASE_DB") != "" {
			env.DB = os.Getenv("DATABASE_DB")
		} else {
			return errors.New("POSTGRES_DB is not found")
		}
	}
	if env.Port == "" {
		if os.Getenv("HOST_MACHINE_DATABASE_PORT") != "" {
			env.Port = os.Getenv("HOST_MACHINE_DATABASE_PORT")
		} else {
			env.Port = "5432"
		}
	}
	if env.SslMode == "" {
		if os.Getenv("SSL_MODE") != "" {
			env.SslMode = os.Getenv("SSL_MODE")
		} else {
			env.SslMode = "disable"
		}
	}
	if env.TimeZone == "" {
		if os.Getenv("TIME_ZONE") != "" {
			env.TimeZone = os.Getenv("TIME_ZONE")
		} else {
			env.TimeZone = "Asia/Tokyo"
		}
	}
	env.Dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		env.Host, env.User, env.Pass, env.DB, env.Port, env.SslMode, env.TimeZone)
	return nil
}

// Connect はデータベースに接続します。
func (env *DbEnv) Connect() (*gorm.DB, error) {
  log.Printf("Current DSN:%v", env.Dsn)

	if env.Dsn == "" {
		return nil, errors.New("database DSN is empty. Call InitDB first")
	}

	db, err := gorm.Open(postgres.Open(env.Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB error(Connect): ", err)
		return nil, err
	}

	return db, nil
}

// Close はデータベース接続をクローズします。
func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error getting SQL DB: %v", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}
