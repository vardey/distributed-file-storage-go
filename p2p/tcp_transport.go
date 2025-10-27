package p2p

import (
	"fmt"
	"net"
)

type TCPTransportOpts struct {
	ListenAddress string
	ShakeHands    HandshakeFunc
	Decoder       Decoder
	OnPeer        func(*TCPPeer) error
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
	fmt.Printf("droppin peer: %+v", peer)
	if peer.conn != nil {
		return peer.conn.Close()
	}
	return nil
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch:            make(chan RPC),
	}
}

func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
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
	defer peer.Close()

	if err := t.ShakeHands(peer); err != nil {
		peer.Close()
		fmt.Printf("TCP handshake error: %s\n", err)
		return
	}
	fmt.Printf("New Connection incoming: %+v\n", peer)
	if t.OnPeer != nil {
		if err := t.OnPeer(peer); err != nil {
			return

		}
	}

	rpc := RPC{}
	for {
		if err := t.Decoder.Decode(conn, &rpc); err != nil {
			fmt.Printf("TCP decode error: %s\n", err)
			return
		}
		rpc.From = conn.RemoteAddr()
		t.rpcch <- rpc
	}

}
