package navdata

import (
	"encoding/binary"
	"io"
)

type binaryReader struct {
	r io.Reader
	Checksum Checksum
	io.Reader
}

func newBinaryReader(r io.Reader) *binaryReader {
	return &binaryReader{r: r}
}

// @TODO merge with ReadOrPanic function
// readOrPanic is a helper function that triggers a panic when binary.Read()
// returns an error (EOF, ErrUnexpectedEOF, etc.). This allows us to unwind the
// stack in these cases without using `if err != nil` checks everywhere.
//
// see: ReadNavdata(), which stops the panic() from propagating to the user of
// this library.
func readOrPanic(r io.Reader, value interface{}) {
	if err := binary.Read(r, binary.LittleEndian, value); err != nil {
		panic(err)
	}
}

func (this *binaryReader) ReadOrPanic(value interface{}) {
	readOrPanic(this, value)
}

func (this *binaryReader) Read(buf []byte) (n int, err error) {
	n, err = this.r.Read(buf)
	this.Checksum.Add(buf[0:n])
	return
}
