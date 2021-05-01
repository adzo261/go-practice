package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port to establish TCP connection")
	flag.Parse()
	address := ":" + strconv.Itoa(port)
	tcpAddress, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	tcpDialer, err := net.DialTCP("tcp", nil, tcpAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {

		fmt.Println("Waiting for user input...")
		userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Sent: " + userInput)
		_, err = fmt.Fprint(tcpDialer, userInput)
		if err != nil {
			fmt.Println(err)
		}

		if userInput == "stop" {
			return
		}
		message, _, err := bufio.NewReader(tcpDialer).ReadLine()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print("Received: " + string(message) + "\n")

	}

}
