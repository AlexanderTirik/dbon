package fs

import (
	"dbon/fs/aws"
	"dbon/fs/filesystem"
	"dbon/fs/sql"
)

var systemType = "sql"

type CRUD interface {
	Write(name string, content []byte) error
	Remove(name string) error
	IsFileExist(name string) bool
	GetAllFileNames() []string
	Read(name string) ([]byte, error)
}

func getDBSaver() CRUD {
	switch systemType {
	case "filesystem":
		return filesystem.Filesystem{}
	case "aws":
		return aws.Aws{}
	case "sql":
		return sql.Sql{}
	}
	return filesystem.Filesystem{}
}

func Write(name string, content []byte) error {
	return getDBSaver().Write(name, content)
}

func Remove(name string) error {
	return getDBSaver().Remove(name)
}

func IsFileExist(name string) bool {
	return getDBSaver().IsFileExist(name)
}

func GetAllFileNames() []string {
	return getDBSaver().GetAllFileNames()
}

func Read(name string) ([]byte, error) {
	return getDBSaver().Read(name)
}
