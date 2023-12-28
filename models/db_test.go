package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DbInitCase struct {
	Case         string
	Host         string
	User         string
	Pass         string
	DB           string
	Port         string
	SslMode      string
	TimeZone     string
	IsError      bool
	ErrorMessage string
}

func TestDbInit(t *testing.T) {

	cases := []DbInitCase{
		{
			Case:     "No Error",
			Host:     os.Getenv("DATABASE_HOST"),
			User:     os.Getenv("DATABASE_USER"),
			Pass:     os.Getenv("DATABASE_PASSWORD"),
			DB:       os.Getenv("DATABASE_DB"),
			Port:     os.Getenv("HOST_MACHINE_DATABASE_PORT"),
			SslMode:  os.Getenv("SSL_MODE"),
			TimeZone: os.Getenv("TIME_ZONE"),
			IsError:  false,
		},
		{
			Case:         "Void HOST",
			Host:         "",
			User:         os.Getenv("DATABASE_USER"),
			Pass:         os.Getenv("DATABASE_PASSWORD"),
			DB:           os.Getenv("POSTGRES_DB"),
			Port:         os.Getenv("HOST_MACHINE_DATABASE_PORT"),
			SslMode:      os.Getenv("SSL_MODE"),
			TimeZone:     os.Getenv("TIME_ZONE"),
			IsError:      true,
			ErrorMessage: "host is not found",
		},
		{
			Case:         "Void POSTGRES_USER",
			Host:         os.Getenv("DATABASE_HOST"),
			User:         "",
			Pass:         os.Getenv("DATABASE_PASSWORD"),
			DB:           os.Getenv("DATABASE_DB"),
			Port:         os.Getenv("HOST_MACHINE_DATABASE_PORT"),
			SslMode:      os.Getenv("SSL_MODE"),
			TimeZone:     os.Getenv("TIME_ZONE"),
			IsError:      true,
			ErrorMessage: "POSTGRES_USER is not found",
		},
		{
			Case:         "Void POSTGRES_PASSWORD",
			Host:         os.Getenv("DATABASE_HOST"),
			User:         os.Getenv("DATABASE_USER"),
			Pass:         "",
			DB:           os.Getenv("DATABASE_DB"),
			Port:         os.Getenv("HOST_MACHINE_DATABASE_PORT"),
			SslMode:      os.Getenv("SSL_MODE"),
			TimeZone:     os.Getenv("TIME_ZONE"),
			IsError:      true,
			ErrorMessage: "POSTGRES_PASSWORD is not found",
		},
		{
			Case:         "Void POSTGRES_DB",
			Host:         os.Getenv("DATABASE_HOST"),
			User:         os.Getenv("DATABASE_USER"),
			Pass:         os.Getenv("DATABASE_PASSWORD"),
			DB:           "",
			Port:         os.Getenv("PORT"),
			SslMode:      os.Getenv("SSL_MODE"),
			TimeZone:     os.Getenv("TIME_ZONE"),
			IsError:      true,
			ErrorMessage: "POSTGRES_DB is not found",
		},
	}

	for _, tc := range cases {
		t.Run(tc.Case, func(t *testing.T) {
			dbEnv := DbEnv{
				Host:     tc.Host,
				User:     tc.User,
				Pass:     tc.Pass,
				DB:       tc.DB,
				Port:     tc.Port,
				SslMode:  tc.SslMode,
				TimeZone: tc.TimeZone,
			}
			if tc.Host == "" {
				os.Setenv("DATABASE_HOST", "")
			}
			if tc.User == "" {
				os.Setenv("DATABASE_USER", "")
			}
			if tc.Pass == "" {
				os.Setenv("DATABASE_PASSWORD", "")
			}
			if tc.DB == "" {
				os.Setenv("DATABASE_DB", "")
			}
			err := dbEnv.InitDB()
			if tc.IsError {
				assert.Equal(t, err.Error(), tc.ErrorMessage)
			} else {
				assert.Equal(t, err, nil)
				assert.NotEqual(t, dbEnv.Host, "")
				assert.NotEqual(t, dbEnv.User, "")
				assert.NotEqual(t, dbEnv.Pass, "")
				assert.NotEqual(t, dbEnv.DB, "")
				assert.NotEqual(t, dbEnv.Port, "")
				assert.NotEqual(t, dbEnv.SslMode, "")
				assert.NotEqual(t, dbEnv.TimeZone, "")
				assert.NotEqual(t, dbEnv.Dsn, "")
			}
		})
	}
}
