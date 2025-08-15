package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"netcat/functions"
)

func main() {
	args := os.Args[1:]
	portDefault := "8989"
	if len(args) == 1 {
		portDefault = args[0]
		port := functions.Atoi(args[0])
		if port < 1024 || port > 49151 {
			fmt.Println("Error: Port Range Not Valid!\nUse Port Between 1024 and 49515")
			return
		}
	} else if len(args) > 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	listener, errListen := net.Listen("tcp", ":"+portDefault)
	if errListen != nil {
		log.Fatalln("Error :", errListen)
	}

	functions.Listenning(listener)
}
