package main

import (
	"crypto/tls"
	"io"
	"net"
	"os"
)

func main() {
	cert, err := tls.LoadX509KeyPair("mycert1.cer", "mycert1.key")
	if err != nil {
		panic(err)
	}

	l, err := tls.Listen("tcp", "127.0.0.1:9000", &tls.Config{
		Certificates: []tls.Certificate{cert},
	})
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	w := io.MultiWriter(os.Stdout, conn)
	go io.Copy(w, conn)
	go io.Copy(conn, os.Stdin)
}
