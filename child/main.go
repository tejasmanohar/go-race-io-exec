package main

import (
	"log"
	"net"
	"os"
)

func main() {
	log.SetPrefix("Child: ")
	log.SetOutput(os.Stdout)

	log.Print("Child process started")

	go logData()
	sendData()
}

func sendData() {
	conn, _ := net.Dial("tcp", ":4000")
	defer conn.Close()

	for i := 1; i <= 10000; i++ {
		conn.Write([]byte("ping"))
	}
}

func logData() {
	for {
		log.Print("Data")
	}
}
