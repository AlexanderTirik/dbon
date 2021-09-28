package table

import (
	"errors"
	"reflect"
	"testing"
)

func TestCreateTable(t *testing.T) {
	colTypes := map[string]string{
		"col1": "integer",
		"col2": "char",
		"col3": "color",
	}
	colNames := []string{"col1", "col2", "col3"}
	err := CreateTable("testCreateTable1", colTypes, colNames)

	if err != nil {
		t.Errorf("CreateTable(%q,%q,%q) == %q, want nil", "testCreateTable1", colTypes, colNames, err)
	}

	colTypes = map[string]string{
		"id": "integer",
	}
	colNames = []string{"id"}
	err = CreateTable("testCreateTable2", colTypes, colNames)

	if err == nil {
		t.Errorf("CreateTable(%q,%q,%q) == %q, want %q", "testCreateTable2", colTypes, colNames, err, errors.New("used fixed column name: id"))
	}
}

func TestPostData(t *testing.T) {
	colTypes := map[string]string{
		"col1": "integer",
		"col2": "char",
		"col3": "color",
	}
	colNames := []string{"col1", "col2", "col3"}
	CreateTable("testPostDataTable1", colTypes, colNames)
	data := map[string]string{
		"col1": "1",
		"col2": "q",
		"col3": "123,123,123",
	}
	err := PostData("testPostDataTable1", data)

	if err != nil {
		t.Errorf("PostData(%q,%q) == %q, want nil", "testPostDataTable1", data, err)
	}

	data = map[string]string{
		"col1": "1",
		"col":  "q",
		"col3": "123,123,123",
	}

	err = PostData("testPostDataTable2", data)

	if err == nil {
		t.Errorf("PostData(%q,%q) == %q, want nil", "testPostDataTable2", data, errors.New("wrong columns"))
	}
}

func TestJoinData(t *testing.T) {
	colTypes1 := map[string]string{
		"col1": "integer",
	}
	colNames1 := []string{"col1"}
	CreateTable("testJoinDataTable1", colTypes1, colNames1)
	data1 := map[string]string{
		"col1": "1",
		"id":   "t1",
	}
	PostData("testJoinDataTable1", data1)
	colTypes2 := map[string]string{
		"col2": "integer",
	}
	colNames2 := []string{"col2"}
	CreateTable("testJoinDataTable2", colTypes2, colNames2)
	data2 := map[string]string{
		"col2": "1",
		"id":   "t2",
	}
	PostData("testJoinDataTable2", data2)

	joinedData, err := JoinTables("testJoinDataTable1", "testJoinDataTable2", "col1", "col2")

	wantData := []map[string]string{{
		"testJoinDataTable1_id":   "t1",
		"testJoinDataTable2_id":   "t2",
		"testJoinDataTable1_col1": "1",
		"testJoinDataTable2_col2": "1",
	}}

	if err != nil || !reflect.DeepEqual(joinedData, wantData) {
		t.Errorf("JoinTables(%q,%q,%q,%q) == %q, %q, want %q, nil", "testJoinDataTable1", "testJoinDataTable2", "col1", "col2", joinedData, err, wantData)
	}

}
