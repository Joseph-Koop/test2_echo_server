package main

import(
	"net"
	"fmt"
	"time"
)

func main(){
	conn, err := net.Dial("tcp", "localhost:4000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprint(conn, "      Hello     \n jiejfeijf    \n   je")
	time.Sleep(1 * time.Second)

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println("Received:", string(buf[:n]))
}