package tcp

import (
	"fmt"
	"net"
)

// Server creates a server running on the given port
// runs the given function handleConnection when connection with a client is established
func Server(port string, handleConnection func(net.Conn, *bool)) {
	ln, err := net.Listen("tcp", "127.0.0.1:"+port)
	checkError(err, "Listening error")

	// closing the server
	defer func() {
		ln.Close()
		fmt.Println("Server closed")
	}()

	// listen determines whether the server will keep on listening
	// listen is passed into handleConnection so that handleConnection can close the server
	listen := true
	fmt.Println("Server listening...")

	for listen {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		handleConnection(conn, &listen)
	}
}