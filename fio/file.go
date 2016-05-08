package fio

import (
	"os"
//	"io"
)

func GetFileInfo(path string) (name string, length int64, err error) {
	file, err := os.Open(path)
	
	if err != nil {
	    return "", 0, err
	}
	
	fi, err := file.Stat()
	
	if err != nil {
	    return file.Name(), 0, err
	}
	
	return fi.Name(), fi.Size(), err
}
