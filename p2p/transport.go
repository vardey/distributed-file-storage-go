package p2p

// Transport represents anything that handles network communication.
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}

// Peer represents a remote node in the network.
type Peer interface {
	Close() error
}
