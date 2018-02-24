package sqlitedump

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCars(t *testing.T) {
	var b bytes.Buffer
	out := bufio.NewWriter(&b)
	err := Dump("testing/cars.db", out)
	assert.Nil(t, err)
	out.Flush()
	pythonOutput, _ := ioutil.ReadFile("testing/python.sql")
	assert.Equal(t, pythonOutput, b.Bytes())
}
