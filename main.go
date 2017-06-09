package main

import (
	"fmt"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}

	conns := connectedConns(server)
	for {
		go handleConn(<-conns)
	}
}

func connectedConns(l net.Listener) chan net.Conn {
	ch := make(chan net.Conn)

	go func() {
		client, err := l.Accept()

		if err != nil {
			fmt.Printf("%v", err)
			continue
		}
		ch <- client
	}()
	return ch
}

func handleConn(conn net.Conn) {
	//Buffer incoming data
	buffer := make([]byte, 1024)

	// Read the incoming message into the buffer.
	dataLen, err := conn.Read(buffer[:])
	defer conn.Close()
	// Check & return error if unable to read byte slice from io.ReadCloser
	if err != nil {
		fmt.Fprintln(conn, "ERROR")
		return
	}

	//Convert byte slice to string
	s := string(buffer[:dataLen])

	//Return OK due to successful command output
	fmt.Fprintln(conn, s)
	return
}
