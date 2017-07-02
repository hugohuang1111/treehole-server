package internal

// Connect connect
type Connect interface {
	ID() uint64
	CloseSig() chan bool
	SendChan() chan []byte
	RecvChan() chan map[string]interface{}
}
