package functions

import (
	"fmt"
	"net"
)

func CloseConnection(conn net.Conn, username string, timeNow string) {
	defer conn.Close()
	fmt.Println("username:", username)
	SendMessage(fmt.Sprintf("%s has left our chat...\n", username), conn, timeNow)
	mutex.Lock()
	delete(clients, conn)
	mutex.Unlock()
}
