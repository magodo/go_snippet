package myio

import (
	"bytes"
	"compress/gzip"
	"io"
)

type GzReader struct {
	io.Reader
	buf        *bytes.Buffer
	gw         *gzip.Writer
	isGwClosed bool
}

func NewGzReader(r io.Reader) *GzReader {
	buf := new(bytes.Buffer)
	return &GzReader{
		r,
		buf,
		gzip.NewWriter(buf),
		false,
	}
}

// Read() will read and compress the content from internal io.Reader.
// (in contrast to the standard gzip.Reader)
// When error except EOF occurs, n is 0.
func (r *GzReader) Read(p []byte) (n int, err error) {
	_, err = io.CopyN(r.gw, r.Reader, int64(len(p)))
	if err != nil && err != io.EOF {
		r.gw.Close()
		return 0, err
	}
	// since length of GzReader.buf might be longer than output buffer `p`,
	// so every Read() should check if we have consumed all GzReader.buf
	// (instead of checking if last copy above reaches EOF) so as to decide
	// whether we reach EOF.

	if len(r.buf.Bytes()) == 0 {
		// If there is nothing in GzReader.buf, it doesn't mean we have consumed
		// all the compressed content. In this case, we should close the internal
		// gzip writer, which will flush remaining gz-compressed data (something like
		// crc32 and isize, see: http://www.zlib.org/rfc-gzip.html)
		// Then Read() should go on consuming these remaining stuff in GzReader.buf.
		if !r.isGwClosed {
			r.gw.Close()
		}
		n, err = r.buf.Read(p)
		return n, err
	}
	n, _ = r.buf.Read(p)
	return n, nil
}
