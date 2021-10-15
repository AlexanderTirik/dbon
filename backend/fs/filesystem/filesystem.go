package filesystem

import (
	"io/ioutil"
	"os"
)

func Write(name string, content []byte) error {
	return ioutil.WriteFile(name+".txt", []byte{}, 0600)
}

func Read(name string) ([]byte, error) {
	return ioutil.ReadFile(name + ".txt")
}

func GetAllFileNames() []string {
	files, _ := ioutil.ReadDir("./")
	fileNames := []string{}
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	return fileNames
}

func Remove(name string) error {
	return os.Remove(name + ".txt")
}

func IsFileExist(name string) bool {
	_, err := os.Stat(name + ".txt")
	return err == nil
}
