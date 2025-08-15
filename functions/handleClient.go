package functions

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func HandleClient(conn net.Conn) {
	defer conn.Close()

	mutex.Lock()
	conn.Write([]byte(welcomeMsg))
	mutex.Unlock()

	reader := bufio.NewReader(conn)
	username := ""
	timeNow := ""
	check := false
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

			if len(clients) > 10 {
				conn.Write([]byte("Connections ghir 10 baraka"))
				delete(clients, conn)
				conn.Close()
				check = true
			}

			if !check {
				SendMessage(fmt.Sprintf("🟢 %s has joined the chat\n", name), conn, timeNow)
				for _, msg := range messages {
					conn.Write([]byte(msg))
				}
			}
			username = name
			mutex.Unlock()
			break
		}
	}
	for {
		mutex.Lock()
		if !check {

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
		}
		mutex.Unlock()
	}
}
