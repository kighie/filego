package fio

import (
	"os"
	"io"
	"fmt"
	"path/filepath"
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

func LoadFileTo(path string, out io.Writer) (int64, error) {
	file, err := os.Open(path)
	
	if err != nil {
	    return 0, err
	}
	
	return io.Copy(out, file)
}

func MakeFile(dir string, fileName string) (*os.File, error) {
	fullPath := filepath.Join(dir, fileName)
	
	fmt.Println ("fullPath:" , fullPath, " , dir=", dir, " , fileName=", fileName)
	
	if err := os.MkdirAll(dir,os.ModePerm) ; err != nil {
		return nil, err
	}
	return os.Create(fullPath)
}

func SaveFileTo(file *os.File, in io.Reader) (int64, error) {
	return io.Copy(file, in)
}
