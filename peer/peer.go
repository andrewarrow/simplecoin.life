package peer

import "github.com/andrewarrow/simplecoin.life/client"
import "fmt"
import "net"
import "io"
import "bytes"

func handleRequest(conn net.Conn) {
	buff := make([]byte, 1024)
	_, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	db := client.SqlInit()
	tl := client.TransactionsFrom(db)

	fmt.Println(string(buff))
	conn.Write([]byte(tl.Encode()))
	conn.Close()
}

func SayHello(peer string) {
	conn, err := net.DialTimeout("tcp", peer, 9000*9000)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Fprintf(conn, "Hello simplecoin.life/0.1\n")
	var buff bytes.Buffer
	io.Copy(&buff, conn)
	fmt.Println(string(buff.Bytes()))
}

func Listen(port string) {
	l, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer l.Close()

	fmt.Println("listening on 0.0.0.0:" + port)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}
		go handleRequest(conn)
	}
}
