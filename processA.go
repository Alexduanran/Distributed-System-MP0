package main

import (
	"fmt"

	"github.com/Alexduanran/MP0/msg"
	"github.com/Alexduanran/MP0/tcp"
)

func main() {
	// prompt user to compose a new msg
	// stores the composed msg in newEmail
	newEmail := msg.ComposeEmail()
	newMsg := msg.Msg{newEmail, 0}

	// establish connection with server
	conn := tcp.Connect("8888")

	// send msg to the server
	tcp.Encode(conn, newMsg)

	// hold until ACK message "STOP" is received
	msg := msg.Msg{}
	for msg.ACK == 1 {
		tcp.Decode(conn, &msg) //blocking
	}
	fmt.Println("ACK received")
	fmt.Println("Program exit")
}
