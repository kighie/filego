package common

import (
	"os"
	"fmt"
	"log"
	"strconv"
	"encoding/json"
)

// load configuration json file to conf object
func LoadConfig(filePath string, conf interface{}) error {
	file, err := os.Open(filePath) 
	defer file.Close()
	
	if err != nil {
		log.Fatalln ("Cannot read ", filePath , err)
		return err
	}
	
	log.Println("Loading configuration file ", filePath)
	
	decoder := json.NewDecoder(file)
	
	err = decoder.Decode(&conf)
	
	if err != nil {
		log.Fatalln ("ERR decoding ", filePath , err)
	}
	
	return err
}

// int slice for options
type IntOptions []int


func (p *IntOptions) String() string {
    return fmt.Sprintf("%d", *p)
}

func (p *IntOptions) Set(value string) error {
    tmp, err := strconv.Atoi(value)
    if err != nil {
        *p = append(*p, -1)
    } else {
        *p = append(*p, tmp)
    }
    return nil
}

// string slice for options
type StringOptions []string


func (p *StringOptions) String() string {
    return fmt.Sprintf("%d", *p)
}

func (p *StringOptions) Set(value string) error {
    *p = append(*p, value)
    return nil
}
