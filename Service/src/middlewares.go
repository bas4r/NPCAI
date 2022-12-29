package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type NewContext struct {
	*gin.Context
	DBConn *pgx.Conn
}

func InjectDBToContextMiddleware(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DBConn", conn)
		c.Next()
	}
}
