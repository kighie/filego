package fio

import (
	"os"
	"io"
	"log"
	"bufio"
//	"fmt"
	"path/filepath"
)

const (
	KB1		= 1024
	KB10 	= 10 * KB1
	KB100 	= 100 * KB1
	MB1 	= 1024 * KB1
	MB10	= 10 * MB1
	MB100	= 100 * MB1
	GB1		= 1024 * MB1
	GB10	= 10 * GB1
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
	
//	fmt.Println ("fullPath:" , fullPath, " , dir=", dir, " , fileName=", fileName)
	
	if err := os.MkdirAll(dir,os.ModePerm) ; err != nil {
		return nil, err
	}
	return os.Create(fullPath)
}

func guessBufferSize(length int64) int{
	if length > MB10 {
		return MB1 * 2
	} else if length > MB1 {
		return KB10 * 2
	} else {
		return KB1 * 2
	}
}

func SaveFileTo(file *os.File, in *bufio.Reader, length int64) (int64, error) {
	fb,_ := in.ReadByte()
	
	defer file.Close()
	
	if fb > 31 {
		file.Write([]byte{fb})
	}
	
	buffer := make([]byte, guessBufferSize(length) )
	
	log.Println ("Before Saving file ", file.Name(), ", length=", length, ",buffer=", len(buffer) )
	
	var n int
	var total int64
	var err error
	
	for {
		n, err = in.Read(buffer)
		if err != nil {
            if err != io.EOF {
                log.Println("read error:", file.Name() , err)
            }
            break
        }
		
		total = total + int64(n)
		
		_, err = file.Write(buffer[:n])
		
		if err != nil {
            log.Println("File Write error:", file.Name() , err)
            break
        }
		
//		fmt.Println("* read", n, ", total:" , total)	
		
		if total == length {
			break
		}
    }
	
	return total,err
}
