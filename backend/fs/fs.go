package fs

var system = "aws"
var bucket = "bucketname"
var key = "/"

func Write(name string, content []byte) error {
	switch system {
	case "filesystem":
		return filesystemWrite(name, content)
	case "aws":
		return awsWrite(name, content)
	}
	return filesystemWrite(name, content)
}

func Remove(name string) error {
	switch system {
	case "filesystem":
		return filesystemRemove(name)
	case "aws":
		return awsRemove(name)
	}
	return filesystemRemove(name)
}

func IsFileExist(name string) bool {
	switch system {
	case "filesystem":
		return filesystemIsFileExist(name)
	case "aws":
		return awsIsFileExist(name)
	}
	return filesystemIsFileExist(name)
}

func GetAllFileNames() []string {
	switch system {
	case "filesystem":
		return filesystemGetAllFileNames()
	case "aws":
		return awsGetAllFileNames()
	}
	return filesystemGetAllFileNames()
}

func Read(name string) ([]byte, error) {
	switch system {
	case "filesystem":
		return filesystemRead(name)
	case "aws":
		return awsRead(name)
	}
	return filesystemRead(name)
}
