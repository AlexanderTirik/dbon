package filesystem

import (
	"io/ioutil"
	"os"
)

type Filesystem struct{}

func (s Filesystem) Write(name string, content []byte) error {
	return ioutil.WriteFile(name+".txt", []byte{}, 0600)
}

func (s Filesystem) Read(name string) ([]byte, error) {
	return ioutil.ReadFile(name + ".txt")
}

func (s Filesystem) GetAllFileNames() []string {
	files, _ := ioutil.ReadDir("./")
	fileNames := []string{}
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	return fileNames
}

func (s Filesystem) Remove(name string) error {
	return os.Remove(name + ".txt")
}

func (s Filesystem) IsFileExist(name string) bool {
	_, err := os.Stat(name + ".txt")
	return err == nil
}
