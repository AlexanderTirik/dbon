package db

import (
	"dbon/fs"
	"dbon/helpers"
	"dbon/table"
	"errors"
)

var db string

func isDBExist(db string) bool {
	return fs.IsFileExist(db)
}

func CreateDB(name string) error {
	if err := helpers.ValidateName(name); err != nil {
		return err
	} else if isDBExist(name) {
		return errors.New("DB already exist")
	} else {
		db = name
		table.CleanTables()
		fs.Write(name, []byte{})
	}
	return nil
}

func FetchDB(name string) error {
	if isDBExist(name) {
		if db != name {
			if db != "" {
				table.SaveCurrentTables(db)
			}
			table.ReadTables(name)
			db = name
		}
	} else {
		return errors.New("DB doesn't exist")
	}
	return nil
}

func RemoveDB(name string) error {
	if isDBExist(name) {
		db = ""
		fs.Remove(name)
		return nil
	}
	return errors.New("DB doesn't exist")
}
