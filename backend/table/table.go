package table

import (
	"dbon/helpers"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/google/uuid"
)

var tables = make(map[string]map[string]string)
var tablesColNames = make(map[string][]string)
var tablesData = make(map[string][]map[string]string)

func CreateTable(name string, colTypes map[string]string, colNames []string) error {
	if _, ok := tables[name]; ok {
		return errors.New("table already exist")
	}
	colTypesArr := helpers.GetMapKeys(colTypes)
	colNamesArr := colNames
	sort.Strings(colTypesArr)
	sort.Strings(colNamesArr)
	for i, v := range colNamesArr {
		if v == "id" {
			return errors.New("used fixed column name: id")
		}
		if i+1 != len(colNamesArr) && colNamesArr[i] == colNamesArr[i+1] {
			return fmt.Errorf("key repeats: %v", colNamesArr[i])
		}
	}
	if !helpers.ArrEq(colTypesArr, colNamesArr) {
		return errors.New("wrong column types or column names")
	}
	for _, v := range colTypes {
		if !helpers.IsValidType(v) {
			return fmt.Errorf("wrong type: %v", v)
		}
	}
	colTypes["id"] = "string"
	colNames = append(colNames, "id")
	tables[name] = colTypes
	tablesColNames[name] = colNames
	return nil
}

func DeleteTable(tableName string) error {
	if _, ok := tables[tableName]; !ok {
		return errors.New("table doesn't exist")
	}
	delete(tables, tableName)
	delete(tablesColNames, tableName)
	delete(tablesData, tableName)
	return nil
}

func GetTableColNames(tableName string) ([]string, error) {
	if _, ok := tablesColNames[tableName]; !ok {
		return nil, errors.New("table doesn't exist")
	}
	return tablesColNames[tableName], nil
}

func GetTablesNames() []string {
	keys := make([]string, len(tables))
	i := 0
	for k := range tables {
		keys[i] = k
		i++
	}
	return keys
}

func PostData(tableName string, data map[string]string) error {
	if _, ok := tables[tableName]; !ok {
		return errors.New("table doesn't exist")
	}
	if !isProvidedDataColsValid(tableName, helpers.GetMapKeys(data)) {
		return errors.New("wrong columns")
	}
	for k, v := range data {
		isValid := validateData(tableName, k, v)
		if !isValid {
			return errors.New("not valid data")
		}
	}
	if _, ok := data["id"]; !ok {
		data["id"] = uuid.NewString()
	}
	err := addData(tableName, data)
	if err != nil {
		return err
	}
	return nil
}

func addData(tableName string, data map[string]string) error {
	id := data["id"]
	for _, v := range tablesData[tableName] {
		if v["id"] == id {
			return errors.New("data with such id already exists")
		}
	}
	tablesData[tableName] = append(tablesData[tableName], data)
	return nil
}

func validateData(table string, colName string, value string) bool {
	colType := tables[table][colName]
	return helpers.Validate(value, colType)
}

func GetAllTableData(tableName string) []map[string]string {
	return tablesData[tableName]
}

func GetData(tableName string, id string) (map[string]string, error) {
	dataIndex, err := findElementIndex(tableName, id)
	if err != nil {
		return nil, err
	}
	return tablesData[tableName][*dataIndex], nil
}

func findElementIndex(tableName string, id string) (*int, error) {
	data := tablesData[tableName]
	for i, v := range data {
		if v["id"] == id {
			return &i, nil
		}
	}
	return nil, errors.New("invalid id")
}

func DeleteData(tableName string, id string) error {
	index, err := findElementIndex(tableName, id)
	if err != nil {
		return err
	}
	tablesData[tableName] = helpers.Remove(tablesData[tableName], *index)
	return nil
}

type TableMap struct {
	Tables         map[string]map[string]string   `json:"tables"`
	TablesColNames map[string][]string            `json:"tablesColNames"`
	TablesData     map[string][]map[string]string `json:"tablesData"`
}

func setReadedTable(tableMap TableMap) {
	tables = tableMap.Tables
	tablesColNames = tableMap.TablesColNames
	tablesData = tableMap.TablesData
}

func fetchTablesFromFile(dbName string) TableMap {
	body, _ := ioutil.ReadFile(dbName + ".txt")
	var tableMap TableMap
	json.Unmarshal(body, &tableMap)
	return tableMap
}

func ReadTables(dbName string) {
	tableMap := fetchTablesFromFile(dbName)
	setReadedTable(tableMap)
}

func SaveCurrentTables(dbName string) {
	tableMap := TableMap{tables, tablesColNames, tablesData}
	filename := dbName + ".txt"
	json, _ := json.Marshal(tableMap)
	ioutil.WriteFile(filename, json, 0600)
}

func isProvidedDataColsValid(t string, cols []string) bool {
	colNamesArr, _ := GetTableColNames(t)

	haveId := false
	for _, c := range cols {
		if c == "id" {
			haveId = true
		}
	}
	if !haveId {
		cols = append(cols, "id")
	}

	sort.Strings(cols)
	sort.Strings(colNamesArr)

	fmt.Print(cols, colNamesArr)

	return helpers.ArrEq(cols, colNamesArr)
}

func JoinTables(t1 string, t2 string, on1 string, on2 string) ([]map[string]string, error) {
	if _, ok1 := tables[t1]; !ok1 {
		return nil, errors.New("first table doesn't exist")
	}
	if _, ok2 := tables[t2]; !ok2 {
		return nil, errors.New("second table doesn't exist")
	}
	d1 := tablesData[t1]
	d2 := tablesData[t2]
	var result []map[string]string
	for _, v1 := range d1 {
		for _, v2 := range d2 {
			if v1[on1] == v2[on2] {
				fmt.Print(v1, v2)
				result = append(result, helpers.JoinMaps(t1, v1, t2, v2))
			}
		}
	}
	return result, nil
}
