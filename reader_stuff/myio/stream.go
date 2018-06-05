package myio

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"hash"
	"io"
)

////////////////

//StreamReader records how many bytes is read from stream and calculate the md5sum
type StreamReader struct {
	N    int
	hash hash.Hash
	io.Reader
}

//NewStreamReader create a new StreamReader object
func NewStreamReader(r io.Reader) *StreamReader {
	return &StreamReader{Reader: r, hash: md5.New()}
}

func (r *StreamReader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	if err != nil {
		return
	}
	r.N += n
	// no error checked since hash.Hash's write never return error
	r.hash.Write(p[:n])
	return
}

//GetMD5 get the md5sum string representation of the read stream
func (r *StreamReader) GetMD5() string {
	return hex.EncodeToString(r.hash.Sum(nil)[:])
}

////////////////

type FirstBreakWriter struct {
	isNew bool
	io.Writer
}

func (w *FirstBreakWriter) Write(p []byte) (n int, err error) {
	if w.isNew {
		w.isNew = false
		return 0, errors.New("Just break, for no reason.")
	}
	return w.Writer.Write(p)
}

func NewFirstBreakWriter(w io.Writer) *FirstBreakWriter {
	return &FirstBreakWriter{true, w}
}

////////////////
