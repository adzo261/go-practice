package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"net"
	"strconv"
)

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}

	return value > 1
}
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

	tcpListener, err := net.ListenTCP("tcp", tcpAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {

		conn, err := tcpListener.Accept()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("New Connection Established")
		go handleClient(conn)

	}

}

func handleClient(conn net.Conn) {
	for {
		fmt.Println("Waiting for query...")
		rawQuery, _, err := bufio.NewReader(conn).ReadLine()
		if err != nil {
			fmt.Println(err)
		}
		queryString := string(rawQuery)

		fmt.Print("Received: " + queryString + "\n")
		if queryString == "stop" {
			conn.Close()
			return
		}
		query, err := strconv.Atoi(queryString)
		if err != nil {
			fmt.Println(err)
		}
		result := isPrime(query)
		var resultString string = queryString
		if result {
			resultString += " is a prime number"
		} else {
			resultString += " is not a prime number"
		}
		fmt.Fprintln(conn, resultString)

	}
}
