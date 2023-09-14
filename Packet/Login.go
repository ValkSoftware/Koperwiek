package packet

import (
	"log"
	"net"

	"github.com/google/uuid"
	server "valksoftware.nl/koperwiek/Server"
)

func HandleLogin(size int32, id byte, s Serializer, conn net.Conn) {
	switch id {
	case 0x00:
		log.Println("Login Start received")

		username := s.ReadString()
		has_uuid := s.ReadBool()

		a := server.GetClient(conn.RemoteAddr())
		a.SetUsername(username)

		if has_uuid {
			a.SetUUID(s.ReadUUID())
		}

		// since we don't support online mode for now
		// we will immediately send the success packet
		s.Clear()
		s.WriteByte(0x02)
		u, _ := uuid.New().MarshalBinary()
		s.WriteUUID(string(u))
		s.WriteString(username)
		s.WriteVarint(0)

		return
	default:
		return
	}
}
