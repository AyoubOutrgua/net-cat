package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	port := "8989"
	if len(args) == 1 {
		port = args[0]
	} else if len(args) > 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	fmt.Println(port)

	listener, errListen := net.Listen("tcp", ":"+port)
	if errListen != nil {
		log.Fatalln("Error :", errListen)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error :", err)
			continue
		}
		go HandleClient(conn)
	}
}

var welcomeMsg string = "Welcome to TCP-Chat!\n         _nnnn_\n        dGGGGMMb\n       @p~qp~~qMb\n       M|@||@) M|\n       @,----.JM|\n      JS^\\__/  qKL\n     dZP        qKRb\n    dZP          qKKb\n   fZP            SMMb\n   HZM            MMMM\n   FqM            MMMM\n __| \".        |\\dS\"qML\n |    `.       | `' \\Zq\n_)      \\.___.,|     .'\n\\____   )MMMMMP|   .'\n     `-'       `--'\n[ENTER YOUR NAME]:"

func HandleClient(conn net.Conn) {
	fmt.Println("connect seccess")
	defer conn.Close()
	conn.Write([]byte(welcomeMsg))
	fmt.Println("client jdid connecter:", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	name := ""
	msgMap := make(map[string][]string)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("client khraj:", conn.RemoteAddr())
			return
		}
		if name == "" {
			name = msg
			continue
		}
		if name != "" {
		}
		now := time.Now()
		timeNow := now.Format("2006-01-02 15:04:05")
		slice := []string{timeNow, name, msg}
		fmt.Println("slice :", slice)
		msgMap[name] = slice
		value := msgMap[name]
		fmt.Printf("[%s] [%s] [%s] :\n", value[1], value[1], value[2])

		conn.Write([]byte("wsl msg ok :" + msg))
	}
}
