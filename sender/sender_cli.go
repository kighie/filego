package sender

import (
	"flag"
	"log"
)


func ParseOptions() SenderInfo {
	var info SenderInfo
	var bSender bool
	
	flag.BoolVar(&bSender,"sender",false, "Sender mode.")
	flag.StringVar(&info.ReceiverAddr,"to","", "Receiver address. ex) 172.1.1.32:3000")
	flag.StringVar(&info.FilePath,"file","","File to send")

	flag.Parse()
	
	if info.ReceiverAddr == "" {
		flag.Usage()
		log.Fatalln("-to is not set.")
	} 
	
	if info.FilePath == "" {
		flag.Usage()
		log.Fatalln("-file is not set.")
	} 
	
	log.Println(info)
	return info
}
