package middlewares

import (
	"net/http"
	"wiki/db"
	"wiki/table"

	"github.com/gin-gonic/gin"
)

func SetDB(c *gin.Context) {
	d := c.Param("db")
	err := db.FetchDB(d)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}
}

func SaveDB(c *gin.Context) {
	d := c.Param("db")
	table.SaveCurrentTables(d)
}
