package functions

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// Handle the entered connection to join in the room and be able to send and recieve messages
func HandleClient(conn net.Conn) {
	defer conn.Close()

	var errWrite error
	_, errWrite = conn.Write([]byte("\033[33m" + welcomeMsg))
	if errWrite != nil {
		fmt.Println(errWrite)
		return
	}

	reader := bufio.NewReader(conn)
	username := ""
	timeNow := ""
	checkConnection := false
	// checking the name of user is valid
	for {
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Name read error:", err)
			return
		}

		username = strings.TrimSpace(name)

		if !IsPrintableRange(username) {
			_, errWrite = conn.Write([]byte("\033[31m❌​ Invalid Username\n"))
			if errWrite != nil {
				fmt.Println(errWrite)
				return
			}
			_, errWrite = conn.Write([]byte("\033[34m[ENTER YOUR NAME]: "))
			if errWrite != nil {
				fmt.Println(errWrite)
				return
			}
		} else if !IsValidUsername(username) {
			_, errWrite = conn.Write([]byte("\033[31m❌​ Invalid Username\n"))
			if errWrite != nil {
				fmt.Println(errWrite)
				return
			}
			_, errWrite = conn.Write([]byte("\033[0m[ENTER YOUR NAME]: "))
			if errWrite != nil {
				fmt.Println(errWrite)
				return
			}
		} else {

			mutexClient.Lock()
			clients[conn] = username
			mutexClient.Unlock()

			if len(clients) > 10 {
				_, errWrite = conn.Write([]byte("\033[31mThe room is full"))
				if errWrite != nil {
					fmt.Println(errWrite)
					return
				}
				mutexClient.Lock()
				delete(clients, conn)
				mutexClient.Unlock()
				conn.Close()
				checkConnection = true
			}

			if !checkConnection {
				SendMessage(fmt.Sprintf("\033[32m🟢 %s has joined our chat...\n", username), conn)
				mutexMessage.Lock()
				for _, msg := range messages {
					_, errWrite = conn.Write([]byte("\033[0m" + msg))
					if errWrite != nil {
						fmt.Println(errWrite)
						return
					}
				}
				mutexMessage.Unlock()
			}
			break
		}
	}

	for {
		timeNow = time.Now().Format("2006-01-02 15:04:05")
		_, errWrite = conn.Write([]byte(fmt.Sprintf("\033[0m[%s][%s]:", timeNow, username)))
		if errWrite != nil {
			fmt.Println(errWrite)
			return
		}
		msg, err := reader.ReadString('\n')
		if err != nil {
			SendMessage(fmt.Sprintf("​\033[31m🔴 %s has left our chat...\n", username), conn)
			mutexClient.Lock()
			delete(clients, conn)
			mutexClient.Unlock()
			break
		}
		msg = strings.TrimSpace(msg)
		if len(msg) > 2000 {
			_, errWrite = conn.Write([]byte("\033[31m❌​ You can't write a message over 2000 letters\n"))
			if errWrite != nil {
				fmt.Println(errWrite)
				return
			}
		} else if !IsPrintableRange(msg) {
			SendMessage("", conn)
		} else {
			timeNow = time.Now().Format("2006-01-02 15:04:05")
			fullMsg := fmt.Sprintf("\033[0m[%s][%s]:%s\n", timeNow, username, msg)
			SendMessage(fullMsg, conn)
			mutexMessage.Lock()
			messages = append(messages, fullMsg)
			mutexMessage.Unlock()
		}
	}
}
