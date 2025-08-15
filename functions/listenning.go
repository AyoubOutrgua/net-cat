package functions

import (
	"fmt"
	"net"
	"sync"
)

var (
	welcomeMsg string = "Welcome to TCP-Chat!\n         _nnnn_\n        dGGGGMMb\n       @p~qp~~qMb\n       M|@||@) M|\n       @,----.JM|\n      JS^\\__/  qKL\n     dZP        qKRb\n    dZP          qKKb\n   fZP            SMMb\n   HZM            MMMM\n   FqM            MMMM\n __| \".        |\\dS\"qML\n |    `.       | `' \\Zq\n_)      \\.___.,|     .'\n\\____   )MMMMMP|   .'\n     `-'       `--'\n[ENTER YOUR NAME]:"
	clients           = make(map[net.Conn]string)
	messages   []string
	mutex      sync.Mutex
)

func Listenning(listener net.Listener) {
	defer listener.Close()

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error :", err)
			continue
		}
		go HandleClient(conn)
	}
}
