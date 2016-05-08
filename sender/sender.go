package sender

import (
	"fmt"
	"log"
	"net"
	"github.com/kighie/filego/protocol"
	"github.com/kighie/filego/fio"
)


type SenderInfo struct {
	ReceiverAddr string
	FilePath	string
	FileName	string
	FileSize	int64
}


func (s SenderInfo) String() string {
	return fmt.Sprint("SenderInfo{ ReceiverAddr=", s.ReceiverAddr, 
		", FilePath=", s.FilePath,  
		", FileName=", s.FileName, ", FileSize=", s.FileSize, "}")
}

func Send(s SenderInfo) error {
	conn, err := net.Dial("tcp", s.ReceiverAddr) 
	defer conn.Close()
	
	if err != nil {
		return err
	}
	
	communicator, err := protocol.NewCommunicator(conn, conn)

	if err != nil {
		return err
	}	

	fileName, length, err := fio.GetFileInfo(s.FilePath)
	s.FileName = fileName
	s.FileSize = length
	
	fmt.Println (s)
	
	var header protocol.Header
	
	header, err = communicator.WriteHeader(protocol.MSG_BEFORE_SEND, length, fileName)
	
	if err != nil {
		return err
	}	
	
	log.Println(header)
//	communicator.
	return nil
}