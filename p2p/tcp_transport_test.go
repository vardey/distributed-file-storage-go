package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	address := ":8080"
	transport := NewTCPTransport(address)
	assert.Equal(t, address, transport.listenAddress)

	assert.Nil(t, transport.ListenAndAccept())

}
