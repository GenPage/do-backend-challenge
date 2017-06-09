package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/genpage/do-backend-challenge/lib/manager"
)

func main() {
	server, err := net.Listen("tcp", "localhost:8080")
	manager := manager.NewManager()

	if err != nil {
		panic(err)
	}

	conns := connectedConns(server)
	for {
		go handleConn(<-conns, manager)
	}
}

func connectedConns(l net.Listener) chan net.Conn {
	ch := make(chan net.Conn)

	go func() {
		for {
			client, err := l.Accept()

			if err != nil {
				fmt.Printf("%v", err)
				continue
			}
			ch <- client
		}
	}()
	return ch
}

func handleConn(conn net.Conn, manager manager.Manager) {
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

	//Make sure we always have two pipe literals to ensure the right number of parameters
	if !(strings.Count(s, "|") == 2) {
		fmt.Fprintln(conn, "ERROR")
		return
	}

	err = nil
	switch strings.Split(s, "|")[0] {
	case "INDEX":
		//err = manager.Index()
	case "REMOVE":
		//err = manager.Remove()
	case "QUERY":
		//err = manager.Query()
	default:
		fmt.Fprintln(conn, "ERROR")
		return
	}

	if err != nil {
		fmt.Fprintln(conn, "FAIL")
		return
	}

	//Return OK due to successful command output
	fmt.Fprintln(conn, "OK")
	return
}
