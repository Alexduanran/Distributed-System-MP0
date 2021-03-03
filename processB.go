package main

import (
	"fmt"
	"github.com/Alexduanran/MP0/msg"
	"github.com/Alexduanran/MP0/tcp"
	"net"
)

func main() {
	// build the server
	tcp.Server("8888", handleConnection)
}

// handleConnection prints out msg received from client
// server stops listening and closes when listen is set to false
func handleConnection(conn net.Conn, listen *bool) {
	// closing connection
	defer func() {
		conn.Close()
		fmt.Println("Connection closed")
	}()

	// receive new msg from the client
	newMsg := msg.Msg{}
	tcp.Decode(conn, &newMsg)
	newMsg.Email.Print()

	// send ACK message to client
	tcp.Encode(conn, msg.Msg{ACK: 1})
	fmt.Println("ACK sent")

	*listen = false // closing the server
}
