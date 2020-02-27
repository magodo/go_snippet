package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	stdout, stderr, _ := CaptureRun(func() error {
		log.Println("log.Println")
		fmt.Fprintln(os.Stdout, "fmt to stdout")
		fmt.Fprintln(os.Stderr, "fmt to stderr")
		return nil
	})

	fmt.Printf(`
stdout
===
%s

stderr
===
%s
`, stdout, stderr)
}

// CaptureRun run the specified function and return its stdout and stderr as stirng
// , all together with the returned error of function to caller.
// This is useful for function which call library out of your control that will log
// to stdout/stderr in variable ways, which you might want to suppress in some cases.
// (e.g. if everything goes fine, and non-verbose mode, then nothing should be output)
func CaptureRun(f func() error) (stdout, stderr string, err error) {
	// setup stdout
	outpipeR, outpipeW, err := os.Pipe()
	if err != nil {
		return "", "", fmt.Errorf("failed to create pipe for stdout: %w", err)
	}
	outCh := make(chan string)
	go func() {
		buf := bytes.NewBufferString("")
		io.Copy(buf, outpipeR)
		// close read end, since have read everything out
		outpipeR.Close()
		outCh <- buf.String()
	}()

	// replace stdout
	defer func(o *os.File) { os.Stdout = o }(os.Stdout)
	os.Stdout = outpipeW

	// setup stderr
	errpipeR, errpipeW, err := os.Pipe()
	if err != nil {
		return "", "", fmt.Errorf("failed to create pipe for stderr: %w", err)
	}
	errCh := make(chan string)
	go func() {
		buf := bytes.NewBufferString("")
		io.Copy(buf, errpipeR)
		// close read end, since have read everything out
		errpipeR.Close()
		errCh <- buf.String()
	}()

	// replace stdout
	defer func(o *os.File) { os.Stderr = o }(os.Stderr)
	os.Stderr = errpipeW

	// replace log writer
	defer log.SetOutput(log.Writer())
	log.SetOutput(errpipeW)

	// invoke function
	err = f()

	// close pipe write end so that the copy routines could finish
	outpipeW.Close()
	errpipeW.Close()

	return <-outCh, <-errCh, err
}
