package main


# A TCP server that handle incomming connections and process them using the chamego 
package main

import (
	"fmt"
	"net"
	"os"
	"github.com/fredcamps/chamego/internal/chame"
)

func main() {
	// Start a TCP server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	message := string(buffer[:n])
	fmt.Println("Received message:", message)

	// Process the message using chamego
}				