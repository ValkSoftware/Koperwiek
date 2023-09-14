package main

import (
	_ "embed"
	"log"
	"net"
	"os"

	packet "valksoftware.nl/koperwiek/Packet"
	server "valksoftware.nl/koperwiek/Server"
)

//go:embed VERSION.txt
var Version string

func main() {

	log.SetFlags(log.Ltime | log.Lshortfile | log.Lmsgprefix)
	log.SetOutput(os.Stdout)

	log.Printf("starting koperwiek version %s...\n", Version)
	log.Printf("running MC version 1.20.1")

	tcp_server, err := net.Listen("tcp", "0.0.0.0:25565")
	if err != nil {
		log.Fatal("server can't start", err)
	}

	for {
		conn, err := tcp_server.Accept()
		if err != nil {
			log.Fatal("connection failed:", conn.RemoteAddr())
		}
		server.AddClient(conn.RemoteAddr(), server.NewClient(server.StateHandshake))

		go packet.HandleConnection(conn)
	}
}
