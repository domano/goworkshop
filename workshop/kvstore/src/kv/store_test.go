package main

import (
	"testing"
	"os"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	"fmt"
)

var filename string
var testmap map[string]string
var kvstore *Store
var key1 string
var key2 string
var value1 string
var value2 string

func TestWriting(t *testing.T) {
	testmap=map[string]string{
		key1:value1,
		key2:value2,
	}


	kvstore.WriteStore(testmap, filename)
	content,_ := ioutil.ReadFile(filename)
	contentstr := string(content)
	assert.Contains(t,contentstr,key1+"="+value1)
	assert.Contains(t,contentstr,key2+"="+value2)
}

func TestReading(t *testing.T) {
	file, _ := os.Create(filename)
	fmt.Fprintf(file, "%v=%v\n%v=%v", key1, value1, key2, value2)
	testmap,_ := kvstore.ReadStore(filename)
	assert.Equal(t, value1, testmap[key1])
	assert.Equal(t, value2, testmap[key2])
}

func TestMain(m *testing.M) {
	filename="testfile.file"
	key1="key1"
	key2="key2"
	value1="value1"
	value2="value2"
	os.Remove(filename)
	kvstore=NewStore()
	os.Exit(m.Run())
}