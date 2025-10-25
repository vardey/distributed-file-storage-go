package p2p

import (
	"fmt"
	"net"
)

type TCPTransportOpts struct {
	ListenAddress string
	ShakeHands    HandshakeFunc
	Decoder       Decoder
}

type TCPPeer struct {
	conn net.Conn
	//if we make the connection, then true
	//if we receive and accept, then false
	outbound bool
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rpcch    chan RPC
}

func (peer *TCPPeer) Close() error {
	return peer.conn.Close()
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func Consume(t *TCPTransport) chan<- RPC {
	return t.rpcch
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch:            make(chan RPC),
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {

		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP start and accept error: %s\n", err)
		}
		fmt.Printf("Listen and Accept...")
		go t.handleConnection(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.ShakeHands(peer); err != nil {
		peer.Close()
		fmt.Printf("TCP handshake error: %s\n", err)
		return
	}
	fmt.Printf("New Connection incoming: %+v\n", peer)
	rpc := &RPC{}
	for {
		if err := t.Decoder.Decode(conn, rpc); err != nil {
			fmt.Printf("TCP decode error: %s\n", err)
			continue
		}
		rpc.From = conn.RemoteAddr()
		fmt.Printf("Received message: %+v", rpc)
	}

}
