package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type Sql struct{}

func (s Sql) Write(name string, content []byte) error {
	name = name + ".txt"
	script := `INSERT INTO dbs (name,data)
	VALUES ($1, $2)
	ON CONFLICT (name)
	DO UPDATE SET data = $2;`
	_, err := DB.Exec(script, name, string(content))
	return err
}

func (s Sql) Remove(name string) error {
	name = name + ".txt"
	script := `DELETE FROM dbs WHERE name=$1;`
	_, err := DB.Exec(script, name)
	return err
}

func (s Sql) Read(name string) ([]byte, error) {
	name = name + ".txt"
	script := `SELECT data FROM dbs WHERE name=$1;`
	row := DB.QueryRow(script, name)
	var data string
	if err := row.Scan(&data); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("read db %d: no such db", name)
		}
		return nil, fmt.Errorf("read db %d: %v", name, err)
	}
	return []byte(data), nil
}

func (s Sql) GetAllFileNames() []string {
	var fileNames []string

	rows, _ := DB.Query("SELECT name FROM dbs")
	defer rows.Close()
	for rows.Next() {
		var filename string
		rows.Scan(&filename)
		fileNames = append(fileNames, filename)
	}
	return fileNames
}

func (s Sql) IsFileExist(name string) bool {
	_, err := Sql{}.Read(name)
	return err == nil
}

func Migrate(db *sql.DB) {
	createTable := `CREATE TABLE Dbs (
		name    text PRIMARY KEY,
		data  text NOT NULL
	);`
	db.Exec(createTable)
}
