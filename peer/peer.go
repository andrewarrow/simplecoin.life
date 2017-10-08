package peer

import "fmt"
import "net"

func handleRequest(conn net.Conn) {
	buff := make([]byte, 1024)
	_, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	conn.Write([]byte("received."))
	conn.Close()
}

func Listen() {
	l, err := net.Listen("tcp", "0.0.0.0:9376")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer l.Close()

	fmt.Println("listening on 0.0.0.0:9376")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}
		go handleRequest(conn)
	}
}
