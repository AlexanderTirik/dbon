package fs

import (
	"io/ioutil"
	"os"
)

func filesystemWrite(name string, content []byte) error {
	return ioutil.WriteFile(name+".txt", []byte{}, 0600)
}

func filesystemRead(name string) ([]byte, error) {
	return ioutil.ReadFile(name + ".txt")
}

func filesystemGetAllFileNames() []string {
	files, _ := ioutil.ReadDir("./")
	fileNames := []string{}
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	return fileNames
}

func filesystemRemove(name string) error {
	return os.Remove(name + ".txt")
}

func filesystemIsFileExist(name string) bool {
	_, err := os.Stat(name + ".txt")
	return err == nil
}
