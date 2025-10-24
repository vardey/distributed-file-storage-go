package p2p

type Decoder interface {
	Decode(any) error
}
