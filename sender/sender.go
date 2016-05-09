package sender

import (
	"fmt"
	"log"
	"net"
	"github.com/kighie/filego/protocol"
	"github.com/kighie/filego/fio"
	"github.com/kighie/filego/common"
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
	
	var stopWatch common.StopWatch
	
	fileName, length, err := fio.GetFileInfo(s.FilePath)
	s.FileName = fileName
	s.FileSize = length
	
	log.Println (s)
	
	conn, err := net.Dial("tcp", s.ReceiverAddr) 
	defer conn.Close()
	
	if err != nil {
		return err
	}
	
	communicator, err := protocol.NewCommunicator(conn)
	
	if err != nil {
		return err
	}	

	stopWatch.Start()
	
	var header protocol.Header
	
	//[1] Sending MSG_BEFORE_SEND
	header, err = communicator.WriteHeader(protocol.MSG_BEFORE_SEND, protocol.MSG_LENGTH_UNDEFINED, "")
	
	if err != nil {
		return err
	}	
	
	communicator.WriteInt64(length)
	communicator.WriteTextField(fileName)
	
	communicator.Flush()
	
	log.Println("Send [MSG_BEFORE_SEND]::",header, stopWatch.Mark())
	
	
	//[2] Receive MSG_PREPARE
	header, err = communicator.ReadHeader()
		
	if err != nil {
		log.Println("Read response header :: ", err)
		return err
	}
	
	var udpPort int16
	
	udpPort, err = communicator.ReadInt16()
	
	log.Println("Response[MSG_PREPARE]::", header, stopWatch.Mark(), " , udp port=" , udpPort)
	
	var sessionid = header.SessionId
	
	//[3] Sending MSG_SEND
	header, err = communicator.WriteHeader(protocol.MSG_SEND, length, sessionid)
	
	if err != nil {
		return err
	}	
	
	communicator.Flush()
	
	var wlen int64
	
	wlen, err = fio.LoadFileTo(s.FilePath, communicator.GetWriter())
	
	
	if err != nil {
		log.Fatal("send file to receiver ", err)
	}
	
	if wlen != s.FileSize {
		return common.NewError("File Length is not the same. meta=", s.FileSize, ", real=" , wlen)
	}
	
	communicator.Flush()
	
	log.Println("Send [MSG_SEND]::",header, stopWatch.Mark(), ", fileSize=" , wlen)
	
	
	//[4] Receive MSG_PREPARE
	header, err = communicator.ReadHeader()
		
	if err != nil {
		log.Println("Read response(MSG_PREPARE) header :: ", err)
		return err
	}
	
//	fmt.Println("** header:" , header )
	
	if header.Type != protocol.MSG_RECEIVED {
		return common.NewError("Unexpected response header type(MSG_RECEIVED):: ", header.Type)
	}
	
	var respMsg string
	respMsg, err = communicator.ReadTextField()
	
	if err != nil {
		log.Println("Read response(MSG_PREPARE) header :: ",header , err)
		return err
	}
	
//	fmt.Println("** respMag:" , respMsg )
	
	wlen, err = communicator.ReadInt64()
	
	if err != nil {
		log.Println("Read response(MSG_PREPARE) header :: ",header, err)
		return err
	}
	
//	fmt.Println("** wlen:" , wlen )
	
	log.Println("Response[MSG_PREPARE]::", header, stopWatch.Mark(), " , response=" , respMsg, wlen)
	
	
	log.Println("Ellapsed time:", stopWatch.Stop())
	
	return nil
}

func SendFile(s SenderInfo) {
	
}


func SendFileUdp(s SenderInfo) {
	
}

