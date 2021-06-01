package sqlite3dump

import (
	"bufio"
	"bytes"
	"database/sql"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCars(t *testing.T) {
	var b bytes.Buffer
	out := bufio.NewWriter(&b)
	err := Dump("testdata/cars.db", out)
	assert.Nil(t, err)
	out.Flush()
	pythonOutput, _ := ioutil.ReadFile("testdata/python.sql")
	assert.Equal(t, pythonOutput, b.Bytes())
	ioutil.WriteFile("out.sql", b.Bytes(), 0644)
}

func TestMigrate(t *testing.T) {
	var b bytes.Buffer
	out := bufio.NewWriter(&b)

	db, err := sql.Open("sqlite3", "testdata/cars.db")
	assert.Nil(t, err)
	defer db.Close()

	err = DumpMigration(db, out)
	assert.Nil(t, err)

	out.Flush()
	pythonOutput, _ := ioutil.ReadFile("testdata/migrate.sql")
	assert.Equal(t, pythonOutput, b.Bytes())
}
