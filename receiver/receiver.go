package receiver

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
//	"container/list"
	"github.com/kighie/filego/common"
	"github.com/kighie/filego/protocol"
	"github.com/kighie/filego/fio"
)


type ReceiverConfig struct {
	ListenPort int
	UdpPorts common.IntOptions
	SenderIps common.StringOptions
	FileDir	string
}

func (rc ReceiverConfig) String() string {
	return fmt.Sprint("ReceiverConfig{ listenPort=", rc.ListenPort, 
		", udpPorts=", rc.UdpPorts, ", senderIps=", rc.SenderIps,
		", FileDir=", rc.FileDir,
		"}")
}

func StartUp(config ReceiverConfig){
	sessionMap := make(map[string]Session)
	
	for _, p := range config.UdpPorts {
		go startUpUdpDaemon(sessionMap, config, p)
	}
	
	
	startUpControlDaemon(sessionMap, config)
}

func startUpControlDaemon(sessionMap  map[string]Session, config ReceiverConfig){
	listener, err := net.Listen("tcp", ":" + strconv.Itoa(config.ListenPort) )
	
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("Receiver starts up on ", listener.Addr())
	
	defer listener.Close()
	
	for {
		conn, err := listener.Accept()
		
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Accept from ", conn.RemoteAddr())
			go Handle(config, sessionMap, conn)
		}
	}
}

func startUpUdpDaemon(sessionMap  map[string]Session, config ReceiverConfig, port int){
	fmt.Println("PORT" , port)
}

func Handle (config ReceiverConfig, sessionMap map[string]Session, conn net.Conn) {
	defer conn.Close()
	session := newSession(config, conn)
	
	sessionMap[session.uid] = session
	defer delete(sessionMap,session.uid)
	
	log.Println("Create:", session)
	
	communicator, err := protocol.NewCommunicator(conn)
	
	if err != nil {
		log.Println("Terminated: Cannot create Communicator :: ", err)
		return
	}
	
	var quit = false
	
	for !quit {
		header, err := communicator.ReadHeader()
		
		if err != nil {
			log.Println("Terminated: Read Header :: ", err)
			break
		}
		
		log.Println("Receive::", header)
		
		switch header.Type {
			case protocol.MSG_BEFORE_SEND :
				if err := session.doBeforeSend(&communicator) ; err != nil {
					log.Println("Proc. MSG_BEFORE_SEND :: ", err)
					break
				}
			case protocol.MSG_SEND :
				if err := session.doSend(&communicator) ; err != nil {
					log.Println("Proc. MSG_SEND :: ", err)
					break
				}
			case protocol.MSG_RECEIVED :
			case protocol.MSG_RELAY :
			case protocol.MSG_QUIT :
				quit = true
			case protocol.MSG_ERROR :
			case protocol.MSG_COMPLETE :
			default :
				log.Println("Unknown message type ", header.Type)
		}
	}
	
	log.Println("Terminated:", session)
}

type Session struct {
	remoteAddr string
	uid string
	time time.Time
	config ReceiverConfig
	fileLen int64
	fileName string
}

func (s Session) String() string {
	return fmt.Sprint("Session{ remoteAddr=", s.remoteAddr, 
		", uid=", s.uid, ", time=", s.time,"}")
}

// Create session 
func newSession(conf ReceiverConfig, conn net.Conn) Session{
	session := Session{ remoteAddr:conn.RemoteAddr().String(), 
			uid : common.MakeUid(conn),
			time : time.Now(),
			config : conf,
		 }
	
	return session
}

func (s *Session)doBeforeSend(communicator *protocol.Communicator) error {
	// BODY 
	var len int64
	var fileName string
	var err error 
	
	if len, err = communicator.ReadInt64() ; err != nil {
		return err
	}
	
	
	if fileName, err = communicator.ReadTextField() ; err != nil {
		return err
	}
	
	s.fileLen = len
	s.fileName = fileName
	
	log.Println("Receive Body :: file length=",len, ", file name=", fileName)
	
	// Response MSG_PREPARE
	var header protocol.Header
	
	header, err = communicator.WriteHeader(protocol.MSG_PREPARE, protocol.MSG_LENGTH_UNDEFINED, s.uid)
	
	if err != nil {
		return err
	}	
	

	communicator.WriteInt16(int16(s.config.UdpPorts[0]))
	
	communicator.Flush()
	
	
	log.Println("Send[MSG_PREPARE]::" , header, ", udpPort=" , s.config.UdpPorts[0])
	
	return nil
}

func (s *Session)doSend(communicator *protocol.Communicator) error {
	file ,err := fio.MakeFile(s.config.FileDir, s.fileName)
	if err != nil {
		return err
	}
	
	length,err = fio.SaveFileTo(file, communicator.GetReader())
	
	log.Println("File Received::", s.fileName ,", size:", s.fileLen, ", file:" , file.Name() )
	
	
//	if err != nil {
//		return err
//	}
	
	return err
}


