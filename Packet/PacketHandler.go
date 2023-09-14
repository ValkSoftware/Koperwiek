package packet

import (
	"log"
	"net"

	server "valksoftware.nl/koperwiek/Server"
)

func HandlePacket(size int32, id byte, s Serializer, conn net.Conn) {

	log.Println(size)
	log.Println(id)

	switch server.GetClient(conn.RemoteAddr()).GetState() {
	case server.StateHandshake:
		HandleHandshake(size, id, s, conn)
		return
	case server.StateStatus:
		log.Println("Status ping received")
		if id == 0 {
			s.Clear()
			s.WriteByte(0x00)
			s.WriteString(server.CreateServerListPingResponse())
			conn.Write(s.Finish())
			return
		} else if id == 1 {
			a := s.ReadLong()
			s.Clear()
			s.WriteByte(0x01)
			s.WriteLong(a)
			conn.Write(s.Finish())
			return
		}
	case server.StateLogin:
		HandleLogin(size, id, s, conn)
	}
}
