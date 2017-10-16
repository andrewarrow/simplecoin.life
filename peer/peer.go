package peer

import "github.com/andrewarrow/simplecoin.life/sql"
import "github.com/andrewarrow/simplecoin.life/crypto"
import "fmt"
import "net"
import "io"
import "bytes"

var myPort = ""
var myPeers = []string{}

func handleRequest(conn net.Conn) {
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	s := localAddr.IP.String()
	fmt.Println("Connection from: " + s)

	buff := make([]byte, 1024)
	_, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	db := sql.SqlInit()
	tl := sql.TransactionsFrom(db)

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

	fmt.Println("got: ", len(buff.Bytes()))
	tl := crypto.DataToTransactionList(buff.Bytes())
	db := sql.SqlInit()
	for _, t := range tl.Items {
		//insert each t
		//unique index
		fmt.Println("adding tx", t.Id)
		sql.InsertTransactionFromPeer(t.Id, t.Owner, t.Previous, t.Transfered, t.Created, db)
	}
}

// https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func Listen(port string) {
	myPort = port
	l, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer l.Close()

	obi := GetOutboundIP()
	fmt.Println("listening at " + obi + ":" + port)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}
		go handleRequest(conn)
	}
}
