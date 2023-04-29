package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	// TODO: write server program to handle concurrent client connections.

}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			return
		}
		fmt.Println("accepted")
		time.Sleep(time.Second)
	}
}
