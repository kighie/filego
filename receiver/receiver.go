package receiver

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
	"container/list"
	"github.com/kighie/filego/common"
)


type ReceiverConfig struct {
	ListenPort int
	UdpPorts common.IntOptions
	SenderIps common.StringOptions
}

func (rc ReceiverConfig) String() string {
	return fmt.Sprint("ReceiverConfig{ listenPort=", rc.ListenPort, 
		", udpPorts=", rc.UdpPorts, ", senderIps=", rc.SenderIps,"}")
}

func StartUp(config ReceiverConfig){
	listener, err := net.Listen("tcp", ":" + strconv.Itoa(config.ListenPort) )
	
	if err != nil {
		log.Fatal(err)
	}
	
	log.Println("Receiver starts up on ", listener.Addr())
	
	defer listener.Close()
	
	sessionList := list.New()
	
	for {
		conn, err := listener.Accept()
		
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Accept from ", conn.RemoteAddr())
			go Handle(sessionList, conn)
		}
	}
}

func Handle (sessionList *list.List, conn net.Conn) {
	defer conn.Close()
	session := newSession(conn)
	se := sessionList.PushBack(session)
	
	defer sessionList.Remove(se)
	
	log.Println("Create:", session)
	
	
	log.Println("Terminated:", session)
}

type Session struct {
	remoteAddr string
	uid string
	time time.Time
}

func (s Session) String() string {
	return fmt.Sprint("Session{ remoteAddr=", s.remoteAddr, 
		", uid=", s.uid, ", time=", s.time,"}")
}

// Create session 
func newSession(conn net.Conn) Session{
	session := Session{ remoteAddr:conn.RemoteAddr().String(), 
			uid : common.MakeUid(conn),
			time : time.Now(),
		 }
	
	return session
}



