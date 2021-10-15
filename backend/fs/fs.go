package fs

import (
	"dbon/fs/aws"
	"dbon/fs/filesystem"
	"dbon/fs/sql"
)

var system = "sql"

func Write(name string, content []byte) error {
	switch system {
	case "filesystem":
		return filesystem.Write(name, content)
	case "aws":
		return aws.Write(name, content)
	case "sql":
		return sql.Write(name, content)
	}
	return filesystem.Write(name, content)
}

func Remove(name string) error {
	switch system {
	case "filesystem":
		return filesystem.Remove(name)
	case "aws":
		return aws.Remove(name)
	case "sql":
		return sql.Remove(name)
	}
	return filesystem.Remove(name)
}

func IsFileExist(name string) bool {
	switch system {
	case "filesystem":
		return filesystem.IsFileExist(name)
	case "aws":
		return aws.IsFileExist(name)
	case "sql":
		return sql.IsFileExist(name)
	}
	return filesystem.IsFileExist(name)
}

func GetAllFileNames() []string {
	switch system {
	case "filesystem":
		return filesystem.GetAllFileNames()
	case "aws":
		return aws.GetAllFileNames()
	case "sql":
		return sql.GetAllFileNames()
	}
	return filesystem.GetAllFileNames()
}

func Read(name string) ([]byte, error) {
	switch system {
	case "filesystem":
		return filesystem.Read(name)
	case "aws":
		return aws.Read(name)
	case "sql":
		return sql.Read(name)
	}
	return filesystem.Read(name)
}
