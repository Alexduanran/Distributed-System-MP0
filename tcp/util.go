package tcp

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"

	"github.com/Alexduanran/MP0/msg"
)

// checkError checks for error in err, exits
// and prints given error message errMSG to the console if err is not nil
func checkError(err error, errMsg string) {
	if err != nil {
		if errMsg == "" {
			errMsg = "Fatal error"
		}
		fmt.Println(errMsg, ":", err.Error())
		os.Exit(1)
	}
}

// Encode transmits the given message msg to the given connection conn
func Encode(conn net.Conn, msg msg.Msg) {
	encoder := gob.NewEncoder(conn)
	encoder.Encode(msg)
}

// Decode receives message from connection conn and stores it in msg
func Decode(conn net.Conn, msg *msg.Msg) {
	decoder := gob.NewDecoder(conn)
	decoder.Decode(msg)
}