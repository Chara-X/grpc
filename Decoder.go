package grpc

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

type Decoder struct{ r io.Reader }

func NewDecoder(r io.Reader) *Decoder { return &Decoder{r} }
func (dec *Decoder) Decode(v any) error {
	var header = make([]byte, 5)
	var err error
	if _, err = dec.r.Read(header); err == nil {
		return json.NewDecoder(io.LimitReader(dec.r, int64(binary.BigEndian.Uint32(header[1:])))).Decode(v)
	}
	return err
}
func (dec *Decoder) Close() { dec.r.(io.Closer).Close() }
