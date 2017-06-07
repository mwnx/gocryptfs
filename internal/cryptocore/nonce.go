package cryptocore

import (
	"crypto/rand"
	"encoding/binary"
	"log"
)

// RandBytes gets "n" random bytes from /dev/urandom or panics
func RandBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		log.Panic("Failed to read random bytes: " + err.Error())
	}
	return b
}

// RandUint64 returns a secure random uint64
func RandUint64() uint64 {
	b := RandBytes(8)
	return binary.BigEndian.Uint64(b)
}

type nonceGenerator struct {
	nonceLen int // bytes
	fifo     chan []byte
}

const (
	nonceGenThreads = 1
	nonceFifoDepth  = 10
)

func newNonceGen(l int) *nonceGenerator {
	n := nonceGenerator{
		nonceLen: l,
		fifo:     make(chan []byte, nonceFifoDepth),
	}
	for i := 0; i < nonceGenThreads; i++ {
		go n.collect()
	}
	return &n
}

func (n *nonceGenerator) collect() {
	for {
		n.fifo <- RandBytes(n.nonceLen)
	}
}

// Get a random "nonceLen"-byte nonce
func (n *nonceGenerator) Get() []byte {
	return <-n.fifo
}
