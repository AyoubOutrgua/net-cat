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

	var errWrite error
	_, errWrite = conn.Write([]byte(welcomeMsg))
	if errWrite != nil {
		fmt.Println("Failed to write message to", clients[conn])
		return
	}

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

		username = strings.TrimSpace(name)

		if !IsPrintableRange(username) {
			_, errWrite = conn.Write([]byte("Invalid Username\n"))
			if errWrite != nil {
				fmt.Println("Failed to write message to", clients[conn])
				return
			}
			_, errWrite = conn.Write([]byte("[ENTER YOUR NAME]: "))
			if errWrite != nil {
				fmt.Println("Failed to write message to", clients[conn])
				return
			}
		} else if !IsValidUsername(username) {
			_, errWrite = conn.Write([]byte("Invalid Username\n"))
			if errWrite != nil {
				fmt.Println("Failed to write message to", clients[conn])
				return
			}
			_, errWrite = conn.Write([]byte("[ENTER YOUR NAME]: "))
			if errWrite != nil {
				fmt.Println("Failed to write message to", clients[conn])
				return
			}
		} else {

			mutex.Lock()
			clients[conn] = username
			mutex.Unlock()

			if len(clients) > 2 {
				_, errWrite = conn.Write([]byte("The room is full"))
				if errWrite != nil {
					fmt.Println("Failed to write message to", clients[conn])
					return
				}
				mutex.Lock()
				delete(clients, conn)
				mutex.Unlock()
				conn.Close()
				check = true
			}

			if !check {
				SendMessage(fmt.Sprintf("%s has joined our chat...\n", username), conn, timeNow)
				for _, msg := range messages {
					_, errWrite = conn.Write([]byte(msg))
					if errWrite != nil {
						fmt.Println("Failed to write message to", clients[conn])
						return
					}
				}
			}
			break
		}
	}
	for {
		timeNow = time.Now().Format("2006-01-02 15:04:05")
		_, errWrite = conn.Write([]byte(fmt.Sprintf("[%s][%s]: ", timeNow, username)))
		if errWrite != nil {
			fmt.Println("Failed to write message to", clients[conn])
			return
		}
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
