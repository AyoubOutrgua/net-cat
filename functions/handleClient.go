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
			break
		}

		username = strings.TrimSpace(name)

		if !IsPrintableRange(username) {
			_, errWrite = conn.Write([]byte("Invalid Username\n"))
			if errWrite != nil {
				fmt.Println(errWrite)
				return
			}
			_, errWrite = conn.Write([]byte("[ENTER YOUR NAME]: "))
			if errWrite != nil {
				fmt.Println(errWrite)
				return
			}
		} else if !IsValidUsername(username) {
			_, errWrite = conn.Write([]byte("Invalid Username\n"))
			if errWrite != nil {
				fmt.Println(errWrite)
				return
			}
			_, errWrite = conn.Write([]byte("[ENTER YOUR NAME]: "))
			if errWrite != nil {
				fmt.Println(errWrite)
				return
			}
		} else {

			mutexClient.Lock()
			clients[conn] = username
			mutexClient.Unlock()

			if len(clients) > 2 {
				_, errWrite = conn.Write([]byte("The room is full"))
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
				SendMessage(fmt.Sprintf("%s has joined our chat...\n", username), conn)
				mutexMessage.Lock()
				for _, msg := range messages {
					_, errWrite = conn.Write([]byte(msg))
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
		if !checkConnection {
			timeNow = time.Now().Format("2006-01-02 15:04:05")
			_, errWrite = conn.Write([]byte(fmt.Sprintf("[%s][%s]:", timeNow, username)))
			if errWrite != nil {
				fmt.Println(errWrite)
				return
			}
			msg, err := reader.ReadString('\n')
			if err != nil {
				SendMessage(fmt.Sprintf("%s has left our chat...\n", username), conn)
				mutexClient.Lock()
				delete(clients, conn)
				mutexClient.Unlock()
				break
			}
			msg = strings.TrimSpace(msg)
			if !IsPrintableRange(msg) {
				SendMessage("", conn)
			} else {
				timeNow = time.Now().Format("2006-01-02 15:04:05")
				fullMsg := fmt.Sprintf("[%s][%s]:%s\n", timeNow, username, msg)
				SendMessage(fullMsg, conn)
				mutexMessage.Lock()
				messages = append(messages, fullMsg)
				mutexMessage.Unlock()
			}
		}
	}
}
