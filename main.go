//go run main.go
//nc localhost 4000

//Delete text files at start of program
//there are still panics if bad port number

package main

import(
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"
	"os"
	//"unicode/utf8"
	//"io/ioutil"
)

func main(){
	portPtr := flag.Int("port", 4000, "The port the server will run on.")
	flag.Parse()

	var portString string = ":" + strconv.Itoa(*portPtr)

	listener, err := net.Listen("tcp", portString)
	if err != nil{
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Server listening on localhost", portString)
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
	fileName := conn.RemoteAddr().String() + ".txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file; messages will not be logged", err)
	}else{
		defer file.Close()
	}

	defer fmt.Println(timestamp(), "Connection terminated with ", conn.RemoteAddr().String())
	defer conn.Close()
	buf := make([]byte, 1024)
	totalBytesRead := 0

	for{
		//timer := time.NewTimer(30 * time.Second)

		//<-timer.C
		//fmt.Println("Client disconnected due to inactivity")
		//conn.Close()

		conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		n, err := conn.Read(buf)
		//timer.Stop()
		if err != nil{
			fmt.Println("Client disconnected due to inactivity or other error")
			return
		}

		totalBytesRead += n
		if(totalBytesRead > 1024){
			n = 1024 - (totalBytesRead - n)
			buf = buf[:n]
			totalBytesRead = 1024
		}

		//Clean data
		stringData := string(buf[:n])
		runeData := []rune(stringData)
		var newData []rune
		count := 0
		lastDigit := ' '
		for i := 0; i < len(runeData); i++ {
			if lastDigit != ' ' || runeData[i] != lastDigit {
				//runeBites := utf8.RuneLen(runeData[i])
				//count += runeBites
				if count <= 1024 {
					newData = append(newData, runeData[i])
				}else{
					break;
				}
			}
			lastDigit = runeData[i]
		}

		_, err = conn.Write([]byte(string(newData)))
		if err != nil{
			fmt.Println("Error writing to client:", err)
		}
		_, err = file.WriteString(string(newData))
		if err != nil{
			fmt.Println("Error writing to client log file:", err)
		}
	}
}

func timestamp() string{
	t := time.FixedZone("America/Chicago (No DST)", -6*60*60)
	return time.Now().In(t).Format("[2006-01-02 15:04:05]")
}