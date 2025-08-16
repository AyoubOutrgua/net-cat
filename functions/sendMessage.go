package functions

import (
	"fmt"
	"net"
)

func SendMessage(fullMsg string, sender net.Conn, timeNow string) {
	var errWrite error
	mutex.Lock()
	for conn, username := range clients {
		if conn != sender {
			if fullMsg != "" {
				_, errWrite = conn.Write([]byte("\n" + fullMsg))
				if errWrite != nil {
					fmt.Println("Failed to send message to", clients[conn])
				}
				_, errWrite = conn.Write([]byte(fmt.Sprintf("[%s][%s]: ", timeNow, username)))
				if errWrite != nil {
					fmt.Println("Failed to send message to", clients[conn])
				}
			}
		}
	}
	mutex.Unlock()
}
