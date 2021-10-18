package main

import (
	"context"
	"database/sql"
	fsmongo "dbon/fs/mongo"
	fssql "dbon/fs/sql"
	"dbon/middlewares"
	"dbon/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func connectSql() *sql.DB {
	connStr := "user=postgres password=1 dbname=dbon sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	fssql.Migrate(db)
	fssql.DB = db
	return db
}

func connectMongo() (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(`mongodb+srv://m001-student:password@dbon.n6u6v.mongodb.net/myFirstDatabase?retryWrites=true&w=majority`))
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.TODO()
	fsmongo.Context = ctx
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fsmongo.Collection = client.Database("dbon").Collection("dbs")
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
	}
	return client, ctx
}

func main() {
	mongoClient, ctx := connectMongo()
	defer mongoClient.Disconnect(ctx)

	sqlDb := connectSql()
	defer sqlDb.Close()

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
