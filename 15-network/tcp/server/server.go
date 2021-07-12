package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:4000")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting incoming connections", err)
			return
		}
		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn) {
	var response string
	inputReader := bufio.NewReader(os.Stdin)
	for {
		buff := make([]byte, 512)
		_, err := conn.Read(buff)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("Received data : %v", string(buff))
		fmt.Println("Enter the response")
		response, _ = inputReader.ReadString('\n')
		conn.Write([]byte(response))
	}
}
