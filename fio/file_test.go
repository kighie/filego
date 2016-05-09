package fio

import (
	"fmt"
	"bytes"
	"bufio"
	"testing"
)

func TestMakeFile(t *testing.T) {
	file, _ := MakeFile("/tmp/filego", "test.txt")
	fmt.Println(file)
}

func TestReadString(t *testing.T) {
	var bs = []byte{'a','b','c',30,'d','e',30,'f','g','h',30}
	reader := bufio.NewReader(bytes.NewReader(bs))
	
	
	line, err := reader.ReadString(30)
	fmt.Println( "["+line+"]", len(line), err)
	
	line, err = reader.ReadString(30)
	fmt.Println( "["+line+"]", len(line), err)
	
	line, err = reader.ReadString(30)
	fmt.Println( "["+line+"]", len(line), err)
}