package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type Context struct {
	Context gin.Context
	DBConn  *pgx.Conn
}

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}
