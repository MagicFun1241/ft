package ft

import (
	"encoding/ascii85"
	"github.com/vmihailenco/msgpack/v5"
)

type token struct {
	Payload []byte `msgpack:"p"`
	Hash    uint64 `msgpack:"h"`
}

func Sign(payload interface{}, secret []byte) []byte {
	b, _ := msgpack.Marshal(&payload)
	h := hash(b, secret)

	b, _ = msgpack.Marshal(&token{Payload: b, Hash: h})
	r := make([]byte, ascii85.MaxEncodedLen(len(b)))
	ascii85.Encode(r, b)
	return r
}
