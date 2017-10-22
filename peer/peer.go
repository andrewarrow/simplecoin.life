package peer

import "github.com/andrewarrow/simplecoin.life/sql"
import "github.com/andrewarrow/simplecoin.life/crypto"
import "fmt"
import "net"
import "io"
import "bytes"
import "strings"

var myPort = ""
var myPeers = map[string]string{"": ""}

func handleRequest80(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().(*net.TCPAddr)
	s := remoteAddr.IP.String()
	fmt.Println("Connection from: " + s)
	buff := make([]byte, 1024)
	_, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	data := string(buff)
	fmt.Println(string(buff))
	if strings.HasPrefix(data, "IAM ") {
		tokens := strings.Split(data, " ")
		myPeers[tokens[1]] = "hi"
		for k, _ := range myPeers {
			conn.Write([]byte(k + "\n"))
		}
	} else {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
		conn.Write([]byte("Date: Tue, 17 Oct 2017 01:53:16 GMT\r\n"))
		conn.Write([]byte("\r\n"))
		conn.Write([]byte("<html>"))
		conn.Write([]byte("<head>"))
		conn.Write([]byte("<title>the simple coin life</title>"))
		conn.Write([]byte("</head>"))
		conn.Write([]byte("<body>"))
		conn.Write([]byte("<h1>the simple coin life</h1>"))
		conn.Write([]byte("</body>"))
		conn.Write([]byte("</html>"))
	}
	conn.Close()
}

func SendUsername(username string) {
	conn, err := net.DialTimeout("tcp", "simplecoin.life:80", 9000*9000)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Fprintf(conn, "IAM "+username+"")
	var buff bytes.Buffer
	io.Copy(&buff, conn)
	data := string(buff.Bytes())
	fmt.Println("got: ", data)
}

func handleRequest(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().(*net.TCPAddr)
	s := remoteAddr.IP.String()
	fmt.Println("Connection from: " + s)

	buff := make([]byte, 1024)
	_, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	raw := string(buff)
	raw = strings.TrimSpace(raw)
	fmt.Println("|" + raw + "|")
	tokens := strings.Split(raw, " ")
	command := tokens[0]
	if command == "HELLO" {
		conn.Write([]byte("hi"))
	} else if command == "TX" {
		db := sql.SqlInit()
		tl := sql.TransactionsFrom(db)
		conn.Write([]byte(tl.Encode()))
	} else {
		conn.Write([]byte("?"))
	}
	conn.Close()
}

func Handshake(peer string, version string) string {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:8666", peer), 9000*9000)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer conn.Close()
	fmt.Fprintf(conn, "HELLO "+version+"\n")
	var buff bytes.Buffer
	io.Copy(&buff, conn)
	res := string(buff.Bytes())
	return res
}

func SayHello(version string) {
	for _, p := range FindPeers() {
		res := Handshake(p, version)
		fmt.Println(res)
	}
}

func OldSayHello(peer string) {
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

	fmt.Println("listening at " + ":" + port)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}
		go handleRequest(conn)
	}
}
