package functions

import (
	"fmt"
	"net"
)

func SendMessage(fullMsg string, sender net.Conn, timeNow string) {
	var errWrite error
	mutexClient.Lock()
	for conn, username := range clients {
		if conn != sender {
			if fullMsg != "" {
				_, errWrite = conn.Write([]byte("\n" + fullMsg))
				if errWrite != nil {
					fmt.Println(errWrite)
				}
				_, errWrite = conn.Write([]byte(fmt.Sprintf("[%s][%s]: ", timeNow, username)))
				if errWrite != nil {
					fmt.Println(errWrite)
				}
			}
		}
	}
	mutexClient.Unlock()
}
