package receiver

import (
	"os"
	"fmt"
	"flag"
	"log"
	"file_transfer/common"
)


func ParseOptions() ReceiverConfig {
	var conf ReceiverConfig
	var confFile string
	var bRec bool
	
	flag.BoolVar(&bRec,"receiver",false, "Starts up Receiver mode.")
	flag.StringVar(&confFile,"config","", "Configuration file")
	flag.IntVar(&conf.ListenPort,"listen",0,"server listen port")
	flag.Var(&conf.UdpPorts,"udp" ,"udp ports : ex) -udp=123 -udp=456 ...")
	flag.Var(&conf.SenderIps,"serderIp" ,"sender ips : ex) -serderIp=172.1.1.1 ...")
	
	flag.Parse()
	
	if confFile != "" {
		common.LoadConfig(confFile, &conf)
	} else if conf.ListenPort == 0 {
		fmt.Println("Receiver mode needs -config={config file path} or -listen={listen port}")
		os.Exit(0)
	}
	
	log.Println(conf)
	return conf
}

