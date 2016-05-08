package convert

import (
	"fmt"
	"testing"
	"bytes"
)

func TestInt16(t *testing.T) {
	var buffer bytes.Buffer
	convertTest(buffer, true)
}

//go test -bench=. -benchmem -benchtime 2s
func BenchmarkInt64(b *testing.B) {
	var buffer bytes.Buffer
	
	for i := 0; i < b.N; i++ {
		convertTest(buffer, false)
    }
}

func convertTest(buffer bytes.Buffer, prt bool){
	WriteInt64(&buffer, 294803885324)
	i := ReadInt64(&buffer)
	
	if prt { fmt.Println(i) }
	
	WriteInt64(&buffer, -294803885324)
	i = ReadInt64(&buffer)
	
	if prt { fmt.Println(i) }
	
	WriteInt16(&buffer, -12345)
	i16 := ReadInt16(&buffer)
	
	if prt { fmt.Println(i16) }
	
	WriteInt64(&buffer, -1234567890)
	i16 = ReadInt16(&buffer)
	if prt { fmt.Println(i16) }
	i16 = ReadInt16(&buffer)
	if prt { fmt.Println(i16) }
	i16 = ReadInt16(&buffer)
	if prt { fmt.Println(i16) }
	i16 = ReadInt16(&buffer)
	if prt { fmt.Println(i16) }
}
	