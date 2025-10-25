package main

import (
	"log"

	"github.com/kunalvardey/distributed-file-storage-go/p2p"
)

func main() {

	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddress: ":8080",
		Decoder:       p2p.GOBDecoder{},
		ShakeHands:    p2p.NOPHandshakeFunc,
	}

	tr := p2p.NewTCPTransport(tcpTransportOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
