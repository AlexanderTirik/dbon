package middlewares

import (
	"dbon/db"
	"dbon/table"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetDB(c *gin.Context) {
	d := c.Param("db")
	err := db.FetchDB(d)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		c.Abort()
		return
	}
	c.Next()
}

func SaveDB(c *gin.Context) {
	d := c.Param("db")
	table.SaveCurrentTables(d)
	c.Next()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
