package _7val

import "net"

type RaknetSession struct {
	conn                net.Conn
	sendSeq, recvSeq    uint32
	ackQueue, nackQueue map[uint32]bool
	splitTable          map[uint16]map[uint32][]byte
}

// Init initializes session and sets valid initial value.
func (rs *RaknetSession) Init() {

}
