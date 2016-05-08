package main

import (
	"fmt"
	"os"
	"log"
	"github.com/kighie/filego/receiver"
	"github.com/kighie/filego/sender"
)

func main(){
	bRec, bSend := parseMode()
	
	if bRec {
		startReceiver()
	} else if bSend {
		startSender()
	}
}

func parseMode() (bRec, bSend bool){
	if len(os.Args) > 1 {
		fmt.Println (os.Args[1] )
		
		if os.Args[1] == "-receiver" {
			bRec = true
		} else if os.Args[1] == "-sender" {
			bSend = true
		} else {
			fmt.Println ("1 Options must start with -receiver or -sender option")
			os.Exit(0)
		}
	} else {
		fmt.Println ("2 Options must start with -receiver or -sender option")
		os.Exit(0)
	}
	
	return bRec, bSend
}

func startReceiver() {
	recConf := receiver.ParseOptions()
	receiver.StartUp(recConf)
}

func startSender() {
	sendInfo := sender.ParseOptions()
	err := sender.Send(sendInfo)
	
	if err != nil {
		log.Fatalln(sendInfo, err)
	}
}