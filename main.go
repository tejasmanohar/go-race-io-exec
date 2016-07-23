package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	log.SetPrefix("Parent: ")
	log.SetOutput(os.Stdout)

	go startServer()

	cmd := exec.Command("go", "run", "child/main.go")

	cmdStdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Error reading STDOUT: ", err.Error())
	}
	go printPipe(cmdStdout)

	cmd.Run()
	os.Exit(0)
}

func startServer() {
	ln, _ := net.Listen("tcp", ":4000")
	defer ln.Close()
	log.Print("Listening")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	for {
		msg := make([]byte, 1024)
		if _, err := conn.Read(msg); err != nil {
			conn.Close()
		}

		log.Print("Received ", string(msg))
		conn.Write([]byte("pong"))
	}
}

func printPipe(in io.Reader) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		log.Print("Piped STDOUT: ", line)
	}
}
