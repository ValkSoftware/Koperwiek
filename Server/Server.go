package server

import (
	"net"
)

type ClientState int

const (
	StateHandshake ClientState = 0
	StateStatus    ClientState = 1
	StateLogin     ClientState = 2
	StatePlay      ClientState = 3
)

var s Server = Server{
	map[net.Addr]ClientState{},
}

type Server struct {
	clients map[net.Addr]ClientState
}

func AddClient(address net.Addr, state ClientState) {
	s.clients[address] = state
}

func UpdateClient(address net.Addr, state ClientState) {
	s.clients[address] = state
}

func RemoveClient(address net.Addr) {
	delete(s.clients, address)
}

func GetClientState(a net.Addr) ClientState {
	return s.clients[a]
}

func Disconnect(a net.Conn) {
	RemoveClient(a.RemoteAddr())
	a.Close()
}
