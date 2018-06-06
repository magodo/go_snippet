package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type MyReader string

var isFirtTime = true

func (r *MyReader) Read(p []byte) (n int, err error) {
	if isFirtTime {
		isFirtTime = false
		p[0] = 'a'
		return 1, io.EOF
	}
	return 0, io.EOF
}

func main() {
	//buf := new(bytes.Buffer)
	//gw := gzip.NewWriter(buf)
	//gw.Write([]byte{'a'})
	//gw.Close()
	//f, _ := os.Create("a.gz")
	//defer f.Close()
	//f.Write(buf.Bytes())
	//fmt.Println(hex.Dump(buf.Bytes()))

	//f, err := os.Open("a.gz")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//gr, err := gzip.NewReader(f)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_, err = io.Copy(buf, gr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(buf)

	buf := new(bytes.Buffer)
	buf.Write([]byte("abcdef"))
	buf.Truncate(2)
	fmt.Println(buf.Bytes())
}

func tgzfiles(out string, files ...string) (err error) {

	pr, pw, err := os.Pipe()
	if err != nil {
		err = errors.Wrap(err, "Failed to create pipe")
		return
	}
	defer pr.Close()

	echannel := make(chan error)
	go func() {
		var err error

		defer func() {
			echannel <- err
		}()

		// close pipe write end after copying files into pipe
		defer pw.Close()

		gw := gzip.NewWriter(pw)
		defer gw.Close()

		err = tarfiles(gw, files...)
		if err != nil {
			err = errors.Wrapf(err, "Failed to tar files: %s", files)
		}
	}()

	outf, _ := os.Create(out)
	defer outf.Close()

	io.Copy(outf, pr)

	return <-echannel
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
