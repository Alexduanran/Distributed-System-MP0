package main

import (
	"fmt"
	"net"

	"github.com/Alexduanran/MP0/email"
	"github.com/Alexduanran/MP0/tcp"
)

func main() {
	// build the server
	tcp.Server("8888", handleConnection)
}

// handleConnection prints out email received from client
// server stops listening and closes when listen is set to false
func handleConnection(conn net.Conn, listen *bool) {
	// closing connection
	defer func() {
		conn.Close()
		fmt.Println("Connection closed")
	}()

	// receive new email from the client
	newEmail := new(email.Email)
	tcp.Decode(conn, newEmail)
	newEmail.Print()

	// send ACK message to client
	tcp.Encode(conn, "STOP")
	fmt.Println("ACK sent")

	*listen = false // closing the server
}
