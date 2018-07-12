package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/magodo/go_snippet/reader_stuff/myio"
)

/**
 * some reader ==================================> some read/writer(e.g. file)  (main link)
 *                |										|
 *                | (tee)								| retry
 *                |     +------------------+  cp        v
 *                +---->| writer ->reader  | ----> break writer (weak link)
 *					    +------------------+
 *                           (pipe)
 *
 */
func test_break_writer(r io.Reader) {

	pReader, pWriter := io.Pipe()
	tr := io.TeeReader(r, pWriter)

	ch := make(chan error)
	firstBreakWriter := myio.NewFirstBreakWriter(ioutil.Discard)
	// weak link start
	go func() {
		_, err := io.Copy(firstBreakWriter, pReader)
		ch <- err
	}()

	localFile := "foobar.txt"
	f, _ := os.Create(localFile)
	io.Copy(f, tr) // main link start
	f.Close()
	pWriter.Close() // close the pipe after main link finishes

	err := <-ch

	// retry weak link
	for i := 0; i < 3 && err != nil; i++ {
		fmt.Printf("retry: %d/3\n", i)
		f, _ := os.Open(localFile)
		_, err = io.Copy(firstBreakWriter, f)
		f.Close()
	}
}

/**
 * some reader ======================================> some read/writer(e.g. file)  (this is main link)
 *                |											| retry
 *                | (tee)									v
 *                |     +------------------+  exec   +-----------+
 *                +---->| writer -> 	   reader/stdin          |
 *					    +------------------+         +-----------+
 *                           (pipe)						(child proc)
 *
 */
func test_child_proc(r io.Reader) {
	pReader, pWriter := io.Pipe()
	tr := io.TeeReader(r, pWriter)

	ch := make(chan error)

	// weak link start
	go func() {
		ch <- runInChild(pReader)
	}()

	localFile := "foobar.txt"
	f, _ := os.Create(localFile)
	io.Copy(f, tr) // main link start
	f.Close()
	pWriter.Close() // close the pipe after main link finishes
	err := <-ch

	// retry weak link
	for i := 0; i < 3 && err != nil; i++ {
		fmt.Printf("retry: %d/3\n", i)
		f, _ := os.Open(localFile)
		err = runInChild(f)
		f.Close()
	}
}

func main1() {
	src := strings.NewReader("abc")
	r := myio.NewStreamReader(src)
	//test_break_writer(r)
	test_child_proc(r)
	fmt.Println(r.N)
	fmt.Println(r.GetMD5())
}

func main2() {
	f1, err := os.Open("random.bin")
	if err != nil {
		log.Fatal(err)
	}
	r := myio.NewGzReader(f1)

	f, err := os.Create("bar.gz")
	if err != nil {
		log.Fatal(err)
	}

	n, err := io.Copy(f, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Copied: ", n)
}

func main3() {
	r, err := myio.NewTgzReader("a.sh")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("foo.tgz")
	if err != nil {
		log.Fatal(err)
	}
	n, err := io.Copy(f, r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Copied: ", n)
}

func main() {
	main3()
}

////////////

func runInChild(r io.Reader) error {
	cmd := exec.Command("/bin/bash", "-c", `"echo -n [; cat; echo -n ]"`)
	var stdOut, stdErr bytes.Buffer
	cmd.Stdin = r
	cmd.Stderr = &stdErr
	cmd.Stderr = &stdOut
	fmt.Println("Start")
	err := cmd.Run()
	fmt.Println("End")
	if err != nil {
		fmt.Printf("Run command failed: %s\nStderr:%s\n", err, stdErr.String())
		return err
	}
	fmt.Printf("Stdout: %s\n", stdOut.String())
	return nil
}

func printall(r io.Reader) error {
	buf, err := ioutil.ReadAll(r)
	if err == nil {
		fmt.Println(buf)
	}
	return err
}
