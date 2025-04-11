package main

import(
	"fmt"
	"net"
	"time"
)

func main(){
	listener, err := net.Listen("tcp", ":4000")
	if err != nil{
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Server listening on :4000")
	for{
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("Error accepting:", err)
			continue
		}
		fmt.Println(timestamp(), "Connection established with ", conn.RemoteAddr().String())
		go handleConnection(conn);
	}
}

func handleConnection(conn net.Conn){
	defer fmt.Println(timestamp(), "Connection terminated with ", conn.RemoteAddr().String())
	defer conn.Close()
	buf := make([]byte, 1024)

	for{
		n, err := conn.Read(buf)
		if err != nil{
			fmt.Println("Error reading from client:", err)
			return
		}
		_, err = conn.Write(buf[:n])
		if err != nil{
			fmt.Println("Error writing to client:", err)
		}
	}

	
}

func timestamp() string{
	t := time.FixedZone("America/Chicago (No DST)", -6*60*60)
	return time.Now().In(t).Format("[2006-01-02 15:04:05]")
}