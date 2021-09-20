package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

var tables = make(map[string]map[string]string)
var tablesColNames = make(map[string][]string)
var tablesData = make(map[string][]map[string]string)

func addTable(name string, columns map[string]string, keys []string) {
	tables[name] = columns
	tablesColNames[name] = keys
}

func deleteTable(tableName string) {
	delete(tables, tableName)
	delete(tablesColNames, tableName)
	delete(tablesData, tableName)
}

func isTableNameExist(tableName string) bool {
	return len(tables[tableName]) != 0
}

func getTableColNames(tableName string) []string {
	return tablesColNames[tableName]
}

func getTablesNames() []string {
	keys := make([]string, len(tables))
	i := 0
	for k := range tables {
		keys[i] = k
		i++
	}
	return keys
}

func getTable(tableName string) map[string]string {
	return tables[tableName]
}

func addData(tableName string, data map[string]string) error {
	tablesData[tableName] = append(tablesData[tableName], data)
	return nil
}

func getData(tableName string) []map[string]string {
	return tablesData[tableName]
}

func findElementIndex(tableName string, id string) (*int, error) {
	data := tablesData[tableName]
	for i, v := range data {
		if v["id"] == id {
			return &i, nil
		}
	}
	return nil, errors.New("Invalid id")
}

func deleteData(tableName string, id string) error {
	index, err := findElementIndex(tableName, id)
	if err != nil {
		return err
	}
	tablesData[tableName] = remove(tablesData[tableName], *index)
	return nil
}

type TableMap struct {
	Tables         map[string]map[string]string   `json:"tables"`
	TablesColNames map[string][]string            `json:"tablesColNames"`
	TablesData     map[string][]map[string]string `json:"tablesData"`
}

func saveDB(dbName string) {
	tableMap := TableMap{tables, tablesColNames, tablesData}
	filename := dbName + ".txt"
	json, _ := json.Marshal(tableMap)
	ioutil.WriteFile(filename, json, 0600)
}

func readDB(dbName string) {
	body, _ := ioutil.ReadFile(dbName + ".txt")
	var tableMap TableMap
	json.Unmarshal(body, &tableMap)
	tables = tableMap.Tables
	tablesColNames = tableMap.TablesColNames
	tablesData = tableMap.TablesData
}
