package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddress}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	conn, err := t.listener.Accept()
	if err != nil {
		fmt.Printf("TCP start and accept error: %s\n", err)
	}

	go t.handleConnection(conn)
}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	fmt.Printf("New Connection incoming: %+v\n", conn)
}
