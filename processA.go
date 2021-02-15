package main

import (
	"fmt"

	"github.com/Alexduanran/MP0/email"
	"github.com/Alexduanran/MP0/tcp"
)

func main() {
	// prompt user to compose a new email
	// stores the composed email in newEmail
	newEmail := email.Compose()

	// establish connection with server
	conn := tcp.Connect("8888")

	// send email to the server
	tcp.Encode(conn, newEmail)

	// hold until ACK message "STOP" is received
	var msg string
	for msg != "STOP" {
		tcp.Decode(conn, &msg) //blocking
	}
	fmt.Println("ACK received")
	fmt.Println("Program exit")
}
