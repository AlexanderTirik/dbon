package main

import (
	"io/ioutil"
	"strings"

	"github.com/manifoldco/promptui"
)

func remove(slice []map[string]string, s int) []map[string]string {
	return append(slice[:s], slice[s+1:]...)
}

func getConsoleSelectValue(label string, values []string) (*string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: values,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func readConsoleValue(label string, validator func(string) error) (*string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validator,
	}
	value, err := prompt.Run()
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func formatTableData(data []map[string]string, keys []string) []string {
	var formattedData []string
	for _, tableRow := range data {
		formattedRow := " | " + tableRow["id"] + " | "
		for _, v := range keys {
			if v != "id" {
				formattedRow = formattedRow + tableRow[v] + " | "
			}
		}
		formattedData = append(formattedData, formattedRow)
	}
	return formattedData
}

func getIdFromSelectedData(selected string) string {
	return strings.Split(selected, " | ")[1]
}

func removeTxtFromName(fileName string) string {
	return fileName[:(len(fileName) - 4)]
}

func getAllTxtFiles() []string {
	files, _ := ioutil.ReadDir("./")
	var fileNames []string
	for _, f := range files {
		fileName := f.Name()
		err := validateTxtFile(fileName)
		if err == nil {
			fileNames = append(fileNames, removeTxtFromName(fileName))
		}
	}
	return fileNames
}
