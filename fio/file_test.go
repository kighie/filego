package fio

import (
	"fmt"
	"testing"
)

func TestMakeFile(t *testing.T) {
	file, _ := MakeFile("/tmp/filego", "test.txt")
	fmt.Println(file)
}