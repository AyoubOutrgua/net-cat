package functions

import (
	"fmt"
	"net"
	"time"
)

// A function that send messages to all connections expect the sender's connection
func SendMessage(fullMsg string, sender net.Conn) {
	var errWrite error
	timeNow := time.Now().Format("\033[0m2006-01-02 15:04:05")
	mutexClient.Lock()
	for conn, username := range clients {
		if conn != sender {
			if fullMsg != "" {
				_, errWrite = conn.Write([]byte("\033[0m\n" + fullMsg))
				if errWrite != nil {
					fmt.Println(errWrite)
				}
				_, errWrite = conn.Write([]byte(fmt.Sprintf("\033[0m[%s][%s]:", timeNow, username)))
				if errWrite != nil {
					fmt.Println(errWrite)
				}
			}
		}
	}
	mutexClient.Unlock()
}
