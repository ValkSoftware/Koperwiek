package packet

import (
	"log"
	"net"

	server "valksoftware.nl/koperwiek/Server"
)

func HandleConnection(conn net.Conn) {
	addr := conn.RemoteAddr()
	var size int32 = 0
	var id byte

	for {
		buf := make([]byte, 8192)
		log.Println("Reading!")

		count, err := conn.Read(buf)
		if err != nil {
			log.Printf("%s | user has disconnected with error %s", addr, err)
			break
		}

		log.Println(string(buf[:count]))

		if count == 0 {
			log.Printf("%s | sent packet of length 0, dropping...", addr)
			break
		}

		s := Serializer{0, buf}

		if buf[0] == 0xFE {
			server.Disconnect(conn)
			log.Printf("%s | dropping legacy ping...", addr)
			return
		}

		// if int32(count) < size {
		// 	log.Println("read too little...")
		// 	continue
		// }

		size = s.ReadVarint()
		id = s.ReadByte()

		HandlePacket(size, id, s, conn)
		buf = nil
	}

	server.Disconnect(conn)
}
