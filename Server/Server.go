package server

import (
	"net"
)

var s Server = Server{
	map[net.Addr]Client{},
}

type Server struct {
	clients map[net.Addr]Client
}

func AddClient(address net.Addr, client Client) {
	s.clients[address] = client
}

func removeClient(address net.Addr) {
	delete(s.clients, address)
}

func GetClient(a net.Addr) Client {
	return s.clients[a]
}

func UpdateClient(a net.Addr, c Client) {
	s.clients[a] = c
}

func Disconnect(a net.Conn) {
	removeClient(a.RemoteAddr())
	a.Close()
}

func GetIndentifier(a net.Addr) string {
	if s.clients[a].username == "" {
		return a.String()
	} else {
		return s.clients[a].username
	}

}
