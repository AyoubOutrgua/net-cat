package functions

import (
	"fmt"
	"net"
	"time"
)

func SendMessage(fullMsg string, sender net.Conn) {
	var errWrite error
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	mutexClient.Lock()
	for conn, username := range clients {
		if conn != sender {
			if fullMsg != "" {
				_, errWrite = conn.Write([]byte("\n" + fullMsg))
				if errWrite != nil {
					fmt.Println(errWrite)
				}
				_, errWrite = conn.Write([]byte(fmt.Sprintf("[%s][%s]:", timeNow, username)))
				if errWrite != nil {
					fmt.Println(errWrite)
				}
			}
		}
	}
	mutexClient.Unlock()
}
