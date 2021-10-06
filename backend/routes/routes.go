package routes

import (
	"dbon/db"
	"dbon/helpers"
	"dbon/table"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDBNames(c *gin.Context) {
	files := helpers.GetAllTxtFiles()
	c.JSON(http.StatusOK, files)
}

func GetTablesNames(c *gin.Context) {
	t := table.GetTablesNames()
	c.JSON(http.StatusOK, t)
}

type CreatedDB struct {
	Name string `json:"db"`
}

func PostDB(c *gin.Context) {
	var d CreatedDB
	if err := c.BindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	err := db.CreateDB(d.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, d)
}

func DeleteDB(c *gin.Context) {
	d := c.Param("db")
	if err := db.RemoveDB(d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	table.CleanTables()
	c.JSON(http.StatusOK, d)
}

func GetTableColNames(c *gin.Context) {
	t := c.Param("table")
	colNames, err := table.GetTable(t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, colNames)
}

type CreatedTable struct {
	TableName string            `json:"table"`
	ColTypes  map[string]string `json:"colTypes"`
	ColNames  []string          `json:"colNames"`
}

func CreateTable(c *gin.Context) {
	var t CreatedTable
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	if err := table.CreateTable(t.TableName, t.ColTypes, t.ColNames); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
}

func DeleteTable(c *gin.Context) {
	t := c.Param("table")
	if err := table.DeleteTable(t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

type DataResponse struct {
	ColNames []string            `json:"colNames"`
	Data     []map[string]string `json:"data"`
}

func GetAllTableData(c *gin.Context) {
	t := c.Param("table")
	colNames, err := table.GetTableColNames(t)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	data := table.GetAllTableData(t)
	c.JSON(http.StatusOK, DataResponse{colNames, data})
}

func Join(c *gin.Context) {
	t1 := c.Param("table1")
	t2 := c.Param("table2")
	on1 := c.Param("on1")
	on2 := c.Param("on2")
	joinedData, err := table.JoinTables(t1, t2, on1, on2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, joinedData)
}

func GetTableData(c *gin.Context) {
	id := c.Param("id")
	t := c.Param("table")
	data, err := table.GetData(t, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, data)
}

func PostData(c *gin.Context) {
	t := c.Param("table")
	var d map[string]string
	if err := c.BindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	if err := table.PostData(t, d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func DeleteTableData(c *gin.Context) {
	id := c.Param("id")
	t := c.Param("table")
	err := table.DeleteData(t, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
