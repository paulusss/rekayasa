package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	fmt.Println("TCP Address:", tcpAddr)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)

	}

}

func handleClient(conn net.Conn) {
	defer conn.Close() //akan dijalankan terakhir sebelum mengeluarkan return value
	fmt.Println("Mendapat koneksi dari ", conn.RemoteAddr().String())
	namaFile := (conn.RemoteAddr().String()) + " - " + (time.Now().String())
	f, err := os.Create(namaFile)
	checkError(err)
	f.Close()
	fmt.Println("File " + namaFile + " berhasil dibuat")
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		f, err := os.OpenFile(namaFile, os.O_APPEND|os.O_WRONLY, 0600)
		checkError(err)
		defer f.Close()
		n2, err := f.Write(buf[0:n])
		checkError(err)
		fmt.Printf("Sedang menulis %d bytes\n", n2)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
