package main

import (
	"database/sql"
	fssql "dbon/fs/sql"
	"dbon/middlewares"
	"dbon/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	connStr := "user=postgres password=12345 dbname=dbon sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	fssql.Migrate(db)
	fssql.DB = db

	router := gin.Default()
	setDB := middlewares.SetDB
	saveDB := middlewares.SaveDB
	cors := middlewares.CORSMiddleware()
	router.Use(cors)
	router.GET("/", routes.GetAllDBNames)
	router.GET("/:db", setDB, routes.GetTablesNames, saveDB)
	router.POST("/", routes.PostDB)
	router.DELETE("/:db", setDB, routes.DeleteDB)
	router.GET("/:db/table/:table", setDB, routes.GetTableColNames)
	router.POST("/:db/table", setDB, routes.CreateTable, saveDB)
	router.DELETE("/:db/table/:table", setDB, routes.DeleteTable, saveDB)
	router.GET("/:db/table/:table/data", setDB, routes.GetAllTableData)
	router.GET("/:db/join/:table1/on/:on1/with/:table2/on/:on2", setDB, routes.Join)
	router.GET("/:db/table/:table/data/:id", setDB, routes.GetTableData)
	router.POST("/:db/table/:table/data", setDB, routes.PostData, saveDB)
	router.DELETE("/:db/table/:table/data/:id", setDB, routes.DeleteTableData, saveDB)

	router.Run("localhost:8080")
}
