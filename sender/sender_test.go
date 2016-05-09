package sender

import (
	"fmt"
	"math/rand"
	"testing"
	"strconv"
)

func TestSend(t *testing.T) {
	sendTest(1, "/tmp/filego/test_data.txt", "test_data.txt")
	sendTest(2, "/tmp/filego/test_10M.zip", "test_10M.zip")
	sendTest(3, "/tmp/filego/test_110M.zip", "test_110M.zip")
}

func sendTest(count int, filePath string, fileName string){
	
	if fileName == "" {
		fileName = "test_data_" + strconv.Itoa(rand.Int())+"-" +strconv.Itoa(count) + ".dat"
	}
	
	var si = SenderInfo {
		ReceiverAddr:"192.168.219.115:3000",
		FileName: fileName,
		FilePath:filePath,
	}
	
	Send(si)
}

//go test github.com/kighie/filego/sender -bench=BenchmarkSend10M -benchmem   -benchtime 30s
func BenchmarkSend10M(b *testing.B){
	for i:=0; i<b.N ;i++ {
		sendTest(i, "/tmp/filego/test_10M.zip","")
		fmt.Println("***** " , i , "****************")
	}
}

//go test github.com/kighie/filego/sender -bench=BenchmarkSend90 -benchmem  -benchtime 30s
func BenchmarkSend90K(b *testing.B){
	for i:=0; i<b.N ;i++ {
		sendTest(i*10 , "/tmp/filego/test_data.txt","")
		fmt.Println("***** " , i , "****************")
	}
}