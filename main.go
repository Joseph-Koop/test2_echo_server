//main.go

package main

import(
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
	"os"
)

func main(){
	//Port command line flag
	portPtr := flag.Int("port", 4000, "The port the server will run on.")
	personalityPtr := flag.Bool("personality", true, "Enables customized responses to specific requests.")
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
		//Concurrency
		go handleConnection(conn, *personalityPtr);
	}
}

func handleConnection(conn net.Conn, personality bool){
	fileName := conn.RemoteAddr().String() + ".txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file; messages will not be logged", err)
	}else{
		defer file.Close()
	}

	//Logging connections/disconnections
	defer fmt.Println(timestamp(), "Connection terminated with ", conn.RemoteAddr().String())
	defer conn.Close()

	buf := make([]byte, 1024)
	totalBytesRead := 0

	for{
		//Inactivity timeout
		conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		n, err := conn.Read(buf)
		if err != nil{
			fmt.Println("Client disconnected due to inactivity or other error")
			return
		}

		//Overflow protection
		totalBytesRead += n
		if(totalBytesRead > 1024){
			n = 1024 - (totalBytesRead - n)
			buf = buf[:n]
			totalBytesRead = 1024
		}

		//Trim and clean
		stringData := string(buf[:n])
		runeData := []rune(stringData)
		var newData []rune
		lastDigit := ' '
		for i := 0; i < len(runeData); i++ {
			if lastDigit != ' ' || runeData[i] != lastDigit {
				newData = append(newData, runeData[i])
			}
			lastDigit = runeData[i]
		}

		dataS := string(newData)

		checkData := strings.TrimSpace(dataS)
		checkData = strings.ToLower(checkData)

		//Command protocols
		if(strings.HasPrefix(checkData, "/")){
			parts := strings.SplitN(checkData, " ", 2)
			switch parts[0]{
				case "/time":
					dataS = string(timestamp()) + "\n"
				case "/quit":
					fmt.Println("Client terminated the connection")
					conn.Close()
					return
				case "/echo":
					if len(parts) > 1{
						dataS = parts[1] + "\n"
					}else{
						dataS = "\n"
					}
				default:
					dataS = parts[0] + " is not recognized as a command\n"
			}
		//Specialized words
		}else if(personality == true){
			switch checkData{
				case " ", "":
					dataS = "Start yapping . . .  \n"
				case "exam", "test", "quiz", "homework", "assignment", "project":
					dataS = "####\n"
				case "exit", "close", "bye":
					_, err = conn.Write([]byte("Farewell!"))
					if err != nil{
						fmt.Println("Error writing to client:", err)
					}

					fmt.Println("Client terminated the connection")
					conn.Close()
					return
				default:
					dataS = dataS
			}
		}

		//Proper error handling: no panics or crashes
		fmt.Println("Message: ", strings.TrimSpace(dataS))
		_, err = conn.Write([]byte(dataS))
		if err != nil{
			fmt.Println("Error writing to client:", err)
		}

		//Logging and saving messages
		_, err = file.WriteString(string(dataS))
		if err != nil{
			fmt.Println("Error writing to client log file:", err)
		}
	}
}

func timestamp() string{
	t := time.FixedZone("America/Chicago (No DST)", -6*60*60)
	return time.Now().In(t).Format("[2006-01-02 15:04:05]")
}