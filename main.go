package main

import (
	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
)

var tablesValues = make(map[string][]map[string]string)

func readColumn(key string, typeValue string) string {
	validator := func(input string) error { return nil }
	switch typeValue {
	case "real":
		validator = validateReal
	case "char":
		validator = validateChar
	case "integer":
		validator = validateInteger
	}
	label := "Print " + key + ". It should be type: " + typeValue
	val, _ := readConsoleValue(label, validator)
	return *val
}

func getMapKeys(mapValue map[string]string) []string {
	keys := make([]string, len(mapValue))
	i := 0
	for k := range mapValue {
		keys[i] = k
		i++
	}
	return keys
}

func createTable() error {
	tableName, _ := readConsoleValue("Table name: ", validateName)
	columns := make(map[string]string)
	columns["id"] = "integer"
	columnsKeys := []string{"id"}
createTableLoop:
	for {
		prompt := promptui.Select{
			Label: "Creating columns",
			Items: []string{"Add column", "Back"},
		}

		_, result, _ := prompt.Run()

		switch result {
		case "Add column":
			{
				newColumn := getNewColumn()
				columns[newColumn.Name] = newColumn.Type
				columnsKeys = append(columnsKeys, newColumn.Name)
			}
		case "Back":
			break createTableLoop
		}
	}
	addTable(*tableName, columns, columnsKeys)
	return nil
}

type Column struct {
	Name string
	Type string
}

func getNewColumn() *Column {
	colName, _ := readConsoleValue("Column name: ", validateColName)
	typeValue, _ := getConsoleSelectValue("Select column type: ", []string{"integer", "real", "char", "string"})
	return &Column{Name: *colName, Type: *typeValue}
}

func addDataToTable(tableName string) error {
	tableColNames := getTableColNames(tableName)
	table := getTable(tableName)
	data := make(map[string]string)
	for _, k := range tableColNames {
		if k == "id" {
			data[k] = uuid.NewString()
		} else {
			value := readColumn(k, table[k])
			data[k] = value
		}
	}
	addData(tableName, data)
	return nil
}

func showTableData(tableName string) string {
	tableData := getData(tableName)
	tableColNames := getTableColNames(tableName)
	formattedTableData := formatTableData(tableData, tableColNames)
	selected, _ := getConsoleSelectValue("Data", formattedTableData)
	return *selected
}

func deleteTableData(tableName string) error {
	selectedData := showTableData(tableName)
	selectedId := getIdFromSelectedData(selectedData)
	deleteData(tableName, selectedId)
	return nil
}

func removeTable() {
	tableNames := getTablesNames()
	tableName, _ := getConsoleSelectValue("Select table", append(tableNames, "@Back"))
	if *tableName == "@Back" {
		return
	}
	deleteTable(*tableName)
}

func manipulateWithData() {
	tableNames := getTablesNames()
	for {
		table, _ := getConsoleSelectValue("Select table", append(tableNames, "@Back"))
		if *table == "@Back" {
			return
		}
	columnLoop:
		for {
			act, _ := getConsoleSelectValue("Select action", []string{"Add data", "Check data", "Delete data", "Back"})
			switch *act {
			case "Add data":
				addDataToTable(*table)
			case "Check data":
				showTableData(*table)
			case "Delete data":
				deleteTableData(*table)
			case "Back":
				break columnLoop
			}
		}
	}
}

func saveDBToFile() {
	dbName, _ := readConsoleValue("DB name: ", validateName)
	saveDB(*dbName)
}

func readDBFromFile() {
	txtFiles := getAllTxtFiles()
	txtFiles = append(txtFiles, "@Back")
	dbName, _ := getConsoleSelectValue("Select DB", txtFiles)
	if *dbName == "@Back" {
		return
	}
	readDB(*dbName)
}

func main() {
mainLoop:
	for {
		act, _ := getConsoleSelectValue("Table manipilating", []string{"Create table", "Manipulate with data", "Delete table", "Save DB", "Read DB", "Exit"})

		switch *act {
		case "Create table":
			createTable()
		case "Manipulate with data":
			manipulateWithData()
		case "Delete table":
			removeTable()
		case "Save DB":
			saveDBToFile()
		case "Read DB":
			readDBFromFile()
		case "Exit":
			break mainLoop
		}
	}

}
