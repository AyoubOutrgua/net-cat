package functions

import (
	"fmt"
	"net"
)

func SendMessage(fullMsg string, sender net.Conn, timeNow string) {
	mutex.Lock()
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
	mutex.Unlock()
}
