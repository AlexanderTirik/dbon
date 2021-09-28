package db

import (
	"errors"
	"io/ioutil"
	"os"
	"wiki/table"
)

var db string

func isDBExist(db string) bool {
	filename := db + ".txt"
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}

func CreateDB(name string) error {
	filename := name + ".txt"
	if isDBExist(name) {
		return errors.New("DB alreay exist")
	} else {
		ioutil.WriteFile(filename, []byte{}, 0600)
	}
	return nil
}

func FetchDB(name string) error {
	if isDBExist(name) {
		if db != name {
			if db != "" {
				table.SaveCurrentTables(name)
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
		filename := name + ".txt"
		os.Remove(filename)
		return nil
	}
	return errors.New("DB doesn't exist")
}
