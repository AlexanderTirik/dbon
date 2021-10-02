package helpers

import (
	"io/ioutil"
)

func ArrEq(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Remove(slice []map[string]string, s int) []map[string]string {
	return append(slice[:s], slice[s+1:]...)
}

func GetMapKeys(m map[string]string) []string {
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func removeTxtFromName(fileName string) string {
	return fileName[:(len(fileName) - 4)]
}

func GetAllTxtFiles() []string {
	files, _ := ioutil.ReadDir("./")
	fileNames := []string{}
	for _, f := range files {
		fileName := f.Name()
		err := validateTxtFile(fileName)
		if err == nil {
			fileNames = append(fileNames, removeTxtFromName(fileName))
		}
	}
	return fileNames
}

func Validate(v string, t string) bool {
	var err error
	switch t {
	case "integer":
		err = validateInteger(v)
	case "char":
		err = validateChar(v)
	case "real":
		err = validateReal(v)
	case "color":
		err = validateColor(v, [][]int{{0, 256}, {0, 256}, {0, 256}})
	case "colorInvl":
		err = validateColor(v, intervals)
	}
	return err == nil
}

func JoinMaps(aName string, a map[string]string, bName string, b map[string]string) map[string]string {
	result := make(map[string]string)
	for k, v := range a {
		result[aName+"_"+k] = v
	}
	for k, v := range b {
		result[bName+"_"+k] = v
	}
	return result
}
