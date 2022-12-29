package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Context struct {
	Context gin.Context
	DBConn  *gorm.DB
}

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	DBOptions      string `mapstructure:"POSTGRE_OPTIONS"`
	ServerPort     string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`

	GoogleApiClientID     string `mapstructure:"GOOGLE_API_CLIENT_ID"`
	GoogleApiClientSecret string `mapstructure:"GOOGLE_API_CLIENT_SECRET"`
	OAuthStateString      string `mapstructure:"OAUTH_STATE_STRING"`
}
