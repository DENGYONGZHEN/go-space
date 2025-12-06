package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

type GobCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	dec  *gob.Decoder
	enc  *gob.Encoder
}

// 编译时接口实现检查
// 这里是检查GobCodec有没有实现Codec
var _ Codec = (*GobCodec)(nil)

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

// Close implements Codec.
func (g *GobCodec) Close() error {
	return g.conn.Close()
}

// ReadBody implements Codec.
func (g *GobCodec) ReadBody(body any) error {
	return g.dec.Decode(body)
}

// ReadHeader implements Codec.
func (g *GobCodec) ReadHeader(h *Header) error {
	return g.dec.Decode(h)
}

// Write implements Codec.
func (g *GobCodec) Write(h *Header, body any) (err error) {
	//return 后，方法结束前 g.Close()
	defer func() {
		_ = g.buf.Flush()
		if err != nil {
			_ = g.Close()
		}
	}()

	if err = g.enc.Encode(h); err != nil {
		log.Println("rpc: gob error encoding header:", err)
		//这里只是return
		return
	}
	if err = g.enc.Encode(body); err != nil {
		log.Println("rpc: gob error encoding body:", err)
		return
	}
	return
}
