package myio

import (
	"archive/tar"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"hash"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type TgzReader struct {
	io.Reader
	errchan chan error
	hash    hash.Hash
}

func (r *TgzReader) Hash() string {
	return hex.EncodeToString(r.hash.Sum(nil)[:])
}

func (r *TgzReader) Read(b []byte) (n int, err error) {
	select {
	case err = <-r.errchan:
		return
	default:
		n, err = r.Reader.Read(b)
		r.hash.Write(b[:n])
		return
	}
}

//
//                         +---------------+
//          +-------+      | +--------+    |
//          |    +--+------+-+--+     |    |
//   <-read |    |r | PIPE | |w | gzip|tar | <- write
//          |    +--+------|-+--+     |    |
//          +-------+      | +--------+    |
//                         +---------------+

func NewTgzReader(files ...string) (r *TgzReader, err error) {

	pr, pw, err := os.Pipe()
	tr := &TgzReader{pr, nil, md5.New()}

	if err != nil {
		err = errors.Wrap(err, "Failed to create pipe")
		return
	}

	go func() {

		// close pipe write end after copying files into pipe
		defer pw.Close()

		gw := gzip.NewWriter(pw)
		defer gw.Close()

		err = tarfiles(gw, files...)
		if err != nil {
			err = errors.Wrapf(err, "Failed to tar files: %s", files)
			tr.errchan <- err
		}
	}()

	return tr, nil
}

func Tgzfiles(out string, files ...string) (err error) {

	w, err := os.Create(out)
	if err != nil {
		err = errors.Wrapf(err, "Failed to open %s", out)
		return
	}
	defer w.Close()

	gw := gzip.NewWriter(w)
	defer gw.Close()

	err = tarfiles(gw, files...)
	if err != nil {
		err = errors.Wrapf(err, "Failed to tar files: %s", files)
	}
	return
}

func tarfiles(w io.Writer, files ...string) (err error) {
	//fmt.Printf("tarIt: %s : BEGIN\n", files)
	//defer fmt.Printf("tarIt: %s : END\n", files)

	tw := tar.NewWriter(w)
	defer tw.Close()

	for _, file := range files {
		err = filepath.Walk(file, func(path string, info os.FileInfo, inerr error) (err error) {
			if inerr != nil {
				err = errors.Wrapf(err, "Something wrong when walking through %s", path)
				return
			}
			if info.IsDir() {
				return
			}
			return tarfile(tw, path)
		})
	}

	return
}

func tarfile(w *tar.Writer, file string) (err error) {
	//fmt.Printf("tarAddFile: %s : BEGIN\n", file)
	//defer fmt.Printf("tarAddFile: %s : END\n", file)

	f, err := os.Open(file)
	if err != nil {
		err = errors.Wrapf(err, "Failed to open %s", file)
		return
	}
	defer f.Close()

	info, err := os.Stat(file)
	if err != nil {
		err = errors.Wrapf(err, "Failed to stat %s", file)
		return
	}

	w.WriteHeader(&tar.Header{
		Name:    file,
		Size:    info.Size(),
		Mode:    int64(info.Mode()),
		ModTime: info.ModTime(),
	})

	_, err = io.Copy(w, f)
	return
}
