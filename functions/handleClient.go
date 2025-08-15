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

	conn.Write([]byte(welcomeMsg))

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

		name = strings.TrimSpace(name)

		if !IsPrintableRange(name) {
			conn.Write([]byte("Invalid Username\n"))
			conn.Write([]byte("[ENTER YOUR NAME]: "))
		} else if !IsValidUsername(name) {
			conn.Write([]byte("Invalid Username\n"))
			conn.Write([]byte("[ENTER YOUR NAME]: "))
		} else {

			mutex.Lock()
			clients[conn] = name
			mutex.Unlock()

			if len(clients) > 10 {
				conn.Write([]byte("The room is full"))
				mutex.Lock()
				delete(clients, conn)
				mutex.Unlock()
				conn.Close()
				check = true
			}

			if !check {
				SendMessage(fmt.Sprintf("%s has joined our chat...\n", name), conn, timeNow)
				for _, msg := range messages {
					conn.Write([]byte(msg))
				}
			}
			username = name
			break
		}
	}
	for {
		if !check {
			timeNow = time.Now().Format("2006-01-02 15:04:05")
			conn.Write([]byte(fmt.Sprintf("[%s][%s]: ", timeNow, username)))
			msg, err := reader.ReadString('\n')
			if err != nil {
				SendMessage(fmt.Sprintf("%s has left our chat...\n", username), conn, timeNow)
				mutex.Lock()
				delete(clients, conn)
				mutex.Unlock()
				break
			}
			msg = strings.TrimSpace(msg)
			if msg == "" {
				continue
			}
			if !IsPrintableRange(msg) {
				SendMessage("", conn, timeNow)
			} else {
				fullMsg := fmt.Sprintf("[%s][%s]: %s\n", timeNow, username, msg)
				SendMessage(fullMsg, conn, timeNow)
				mutex.Lock()
				messages = append(messages, fullMsg)
				mutex.Unlock()
			}
		}
	}
}
