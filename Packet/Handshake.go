package packet

import (
	"log"
	"net"

	server "valksoftware.nl/koperwiek/Server"
)

func HandleHandshake(size int32, id byte, s Serializer, conn net.Conn) {
	switch id {
	case 0x00:
		log.Println("Handshake received")
		version := s.ReadVarint()
		if version != 763 {
			conn.Close()
		}
		s.ReadString()
		s.ReadShort()

		state := s.ReadVarint()
		if state == 1 {
			server.UpdateClient(conn.RemoteAddr(), server.StateStatus)
			log.Println("Upgraded to status")
			return
		}

		if state == 2 {
			server.UpdateClient(conn.RemoteAddr(), server.StateLogin)
			log.Println("Upgraded to login")
			return
		}

		return
	default:
		return
	}

}
