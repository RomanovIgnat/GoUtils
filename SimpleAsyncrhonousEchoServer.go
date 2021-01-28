package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalln("Unable to close connection")
		} else {
			log.Println("Connection closed")
		}
	}()

    if _, err := io.Copy(conn, conn); err != nil {
    	log.Fatalln("Unable to read/write data")
	}
}

func main() {
    listener, err := net.Listen("tcp", "0.0.0.0:20080")
    if err != nil {
    	log.Fatalln("Unable to bind the port")
	}

	log.Println("Listening on 0.0.0.0:20080")

    for {
    	conn, err := listener.Accept()
    	log.Println("Received connection")
    	if err != nil {
    		log.Fatalln("Unable to accept connection")
		}

		go echo(conn)
	}

    return
}
