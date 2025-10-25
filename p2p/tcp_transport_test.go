package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	address := ":8080"
	tcpTransportOpts := TCPTransportOpts{
		ListenAddress: address,
		ShakeHands:    NOPHandshakeFunc,
	}

	transport := NewTCPTransport(tcpTransportOpts)
	assert.Equal(t, address, transport.ListenAddress)

	assert.Nil(t, transport.ListenAndAccept())

}
