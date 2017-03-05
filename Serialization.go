package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage:	%s	host:port	filename ", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[2])
	checkError(err)
	defer file.Close()
	fmt.Println("file ", file, " berhasil dibuka")

	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	scanner := bufio.NewScanner(file)
	fmt.Println("Membuka file ", file.Name())
	baris := 1
	for scanner.Scan() {
		fmt.Println("Membuka baris ke-", baris)
		fmt.Println(scanner.Text())
		_, err = conn.Write([]byte(scanner.Text()))
		checkError(err)
		fmt.Println("Baris ke-", baris, " terkirim")
		_, err = conn.Write([]byte("\n"))
		checkError(err)
		fmt.Println("")
		baris = baris + 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
