package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	// (1024 to 49151) ports
	args := os.Args[1:]
	port := "8989"
	if len(args) == 1 {
		prt, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if prt < 1024 || prt > 49151 {
			fmt.Println("Error: Port Range Not Valid!\nUse Port Between 1024 and 49515")
			return
		}
		port = args[0]
	} else if len(args) > 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	fmt.Println("chat rah bda f port:", port)

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

	// check := false
	if len(clients) > 0 {
		conn.Write([]byte("Connections ghir 10 baraka"))
		conn.Close()
		// check = true
		return
	}

	mutex.Lock()
	conn.Write([]byte(welcomeMsg))
	mutex.Unlock()

	reader := bufio.NewReader(conn)
	username := ""
	timeNow := ""
	// check := false
	for {
		timeNow = time.Now().Format("2006-01-02 15:04:05")
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Name read error:", err)
			break
		}

		mutex.Lock()
		name = strings.TrimSpace(name)
		mutex.Unlock()

		if !IsPrintableRange(name) {
			conn.Write([]byte("had smya mchi valid ....\n"))
			conn.Write([]byte("[ENTER YOUR NAME]: "))
		} else if !IsValidUsername(name) {
			conn.Write([]byte("had smya deja kayna \n"))
			conn.Write([]byte("[ENTER YOUR NAME]: "))
		} else {

			mutex.Lock()
			clients[conn] = name

			// if len(clients) > 10 {
			// 	conn.Write([]byte("Connections ghir 10 baraka"))
			// 	delete(clients, conn)
			// 	conn.Close()
			// 	check = true
			// }

			// if !check {
			SendMessage(fmt.Sprintf("🟢 %s has joined the chat\n", name), conn, timeNow)
			for _, msg := range messages {
				conn.Write([]byte(msg))
			}
			// }
			username = name
			mutex.Unlock()
			break
		}
	}
	for {
		mutex.Lock()
		// if !check {

		timeNow = time.Now().Format("2006-01-02 15:04:05")
		conn.Write([]byte(fmt.Sprintf("[%s][%s]: ", timeNow, username)))
		mutex.Unlock()

		msg, err := reader.ReadString('\n')
		if err != nil {
			mutex.Lock()
			SendMessage(fmt.Sprintf("🔴 %s disconnected\n", username), conn, timeNow)
			delete(clients, conn)
			mutex.Unlock()
			break
		}
		mutex.Lock()
		msg = strings.TrimSpace(msg)
		mutex.Unlock()
		if msg == "" {
			continue
		}
		mutex.Lock()
		if !IsPrintableRange(msg) {
			SendMessage("", conn, timeNow)
		} else {
			fullMsg := fmt.Sprintf("[%s][%s]: %s\n", timeNow, username, msg)
			SendMessage(fullMsg, conn, timeNow)
			messages = append(messages, fullMsg)
		}
		// }
		mutex.Unlock()
	}
}

func SendMessage(fullMsg string, sender net.Conn, timeNow string) {
	for conn, username := range clients {
		if conn != sender {
			if fullMsg != "" {
				_, err := conn.Write([]byte("\n" + fullMsg))
				_, err2 := conn.Write([]byte(fmt.Sprintf("[%s][%s]: ", timeNow, username)))
				if err != nil || err2 != nil {
					fmt.Println("Failed to send message to", clients[conn])
				}
			}
		}
	}
}

func IsPrintableRange(nameORmsg string) bool {
	if nameORmsg == "" {
		return false
	}
	for _, v := range nameORmsg {
		if v < 32 || v > 126 {
			return false
		}
	}
	return true
}

func IsValidUsername(username string) bool {
	if len(username) > 15 {
		return false
	}
	if len(clients) == 0 {
		return true
	}
	for _, client := range clients {
		if username == client {
			return false
		}
	}
	return true
}
