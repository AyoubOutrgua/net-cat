package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
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

var (
	welcomeMsg string = "Welcome to TCP-Chat!\n         _nnnn_\n        dGGGGMMb\n       @p~qp~~qMb\n       M|@||@) M|\n       @,----.JM|\n      JS^\\__/  qKL\n     dZP        qKRb\n    dZP          qKKb\n   fZP            SMMb\n   HZM            MMMM\n   FqM            MMMM\n __| \".        |\\dS\"qML\n |    `.       | `' \\Zq\n_)      \\.___.,|     .'\n\\____   )MMMMMP|   .'\n     `-'       `--'\n[ENTER YOUR NAME]:"
	clients           = make(map[net.Conn]string)
	messages   []string
	mutex      sync.Mutex
)

func HandleClient(conn net.Conn) {
	defer conn.Close()
	if len(clients) > 2 {
		conn.Write([]byte("Connections ghir 10 baraka"))
		return
	}
	conn.Write([]byte(welcomeMsg))
	reader := bufio.NewReader(conn)
	username := ""
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	for {
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Name read error:", err)
			return
		}
		name = strings.TrimSpace(name)
		if CheckUsername(name) {
			conn.Write([]byte("maymknch dir name khawi ....\n"))
			conn.Write([]byte("[ENTER YOUR NAME]:"))
		} else {
			SendMessage(fmt.Sprintf("\n🟢 %s has joined the chat\n", name), conn, timeNow)

			clients[conn] = name
			username = name
			break
		}
	}

	for {
		conn.Write([]byte(fmt.Sprintf("[%s][%s]:", timeNow, username)))

		msg, err := reader.ReadString('\n')
		if err != nil {
			SendMessage(fmt.Sprintf("\n🔴 %s disconnected\n", username), conn, timeNow)
			break
		}
		msg = strings.TrimSpace(msg)
		if msg == "" {
			continue
		}
		fullMsg := fmt.Sprintf("\n[%s][%s]: %s\n", timeNow, username, msg)
		SendMessage(fullMsg, conn, timeNow)

	}
}

func SendMessage(message string, sender net.Conn, timeNow string) {
	for conn, username := range clients {
		if conn != sender {
			_, err := conn.Write([]byte(message))
			_, err2 := conn.Write([]byte(fmt.Sprintf("[%s][%s]:", timeNow, username)))
			if err != nil || err2 != nil {
				fmt.Println("Failed to send message to", clients[conn])
			}
		}
	}
}

func CheckUsername(username string) bool {
	return username == ""
}
