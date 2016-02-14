package main

import (
	"testing"
	"os"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
	"fmt"
	"io"
	"bytes"
	"strconv"
	"sync"
)

var filename string
var testmap map[string]string
var kvstore *Store
var key1 string
var key2 string
var value1 string
var value2 string

func TestWriting(t *testing.T) {
	os.Remove(filename)
	kvstore.WriteStore(testmap, filename)
	content,_ := ioutil.ReadFile(filename)
	contentstr := string(content)
	assert.Contains(t,contentstr,key1+"="+value1)
	assert.Contains(t,contentstr,key2+"="+value2)
}

func TestReading(t *testing.T) {
	os.Remove(filename)
	file, _ := os.Create(filename)
	fmt.Fprintf(file, "%v=%v\n%v=%v", key1, value1, key2, value2)
	testmap,_ := kvstore.ReadStore(filename)
	assert.Equal(t, value1, testmap[key1])
	assert.Equal(t, value2, testmap[key2])
}

func TestConcurrency(t *testing.T) {
	os.Remove(filename)
	waitGroup := &sync.WaitGroup{}
	i:=0
	for i<1000 {
		go mainRoutineWithWaitGroup([]string{"kv","a"+strconv.Itoa(i)+"=b"+strconv.Itoa(i)}, waitGroup, filename)
		i++
	}
	waitGroup.Wait()
	file, _ := os.Open(filename)
	j,_:=lineCounter(file)
	assert.Equal(t, i, j)

}

func BenchmarkConcurrency(b *testing.B) {
	b.StopTimer()
	os.Remove(filename)
	b.StartTimer()

	for i:=0;i<b.N;i++  {
		kvstore.WriteStore(testmap, filename)
		kvstore.ReadStore(filename)
	}

}

func TestMain(m *testing.M) {
	filename="testfile.file"
	key1="key1"
	key2="key2"
	value1="value1"
	value2="value2"
	testmap=map[string]string{
		key1:value1,
		key2:value2,
	}
	kvstore=NewStore()
	os.Exit(m.Run())
}


func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 8196)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return count, err
		}

		count += bytes.Count(buf[:c], lineSep)

		if err == io.EOF {
			break
		}
	}

	return count, nil
}