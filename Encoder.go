package grpc

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"net/http"
)

type Encoder struct{ w io.Writer }

func NewEncoder(w io.Writer) *Encoder { return &Encoder{w} }
func (enc *Encoder) Encode(v any) {
	var body, _ = json.Marshal(v)
	enc.w.Write(append(binary.BigEndian.AppendUint32([]byte{0x00}, uint32(len(body))), body...))
	if flusher, ok := enc.w.(http.Flusher); ok {
		flusher.Flush()
	}
}
func (enc *Encoder) Close() { enc.w.(io.Closer).Close() }
