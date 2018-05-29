package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

func main() {
	fcounter := NewFileNewcontentCounter("abc")
	for {
		fmt.Print("start check: ")
		n, err := fcounter.CountNewContentContain("^hello")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
		time.Sleep(time.Second)
	}
}

type FileNewcontentCounter interface {
	// Check how many does newly appended content contain some pattern (regexp)
	CountNewContentContain(pattern string) (int, error)
}

type simpleFileNewcontentCounter struct {
	filename string
	stat     syscall.Stat_t
}

func (fcounter *simpleFileNewcontentCounter) CountNewContentContain(pattern string) (hit int, err error) {

	p, err := regexp.Compile(pattern)
	if err != nil {
		err = errors.Wrapf(err, "regexp compile pattern %s failed", pattern)
		return
	}

	// get file metadata
	oldstat := fcounter.stat
	err = syscall.Stat(fcounter.filename, &fcounter.stat)
	if err != nil {
		// 1. if file is removed, reset internal state
		if err == syscall.ENOENT {
			fcounter.stat = syscall.Stat_t{}
		}
		err = errors.Wrapf(err, "stat %s failed", fcounter.filename)
		return
	}

	if oldstat.Ino != fcounter.stat.Ino {
		// 2. if a same-named file is created, the old stat should be reset for later use
		oldstat = syscall.Stat_t{}
	} else if oldstat.Size > fcounter.stat.Size {
		// 3. if file is truncated, no need to check
		err = errors.New("File is trucated, nothing to check")
		return
	}

	// check whether newly added content contains the pattern
	f, err := os.Open(fcounter.filename)
	if err != nil {
		err = errors.Wrapf(err, "failed to open %s", fcounter.filename)
		return
	}
	defer f.Close()
	// seek to where we read last time
	_, err = f.Seek(oldstat.Size, 0)
	if err != nil {
		err = errors.Wrapf(err, "failed to seek to %d for file %s", oldstat.Size, fcounter.filename)
		return
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if p.Match(scanner.Bytes()) {
			hit++
		}
	}
	if err = scanner.Err(); err != nil {
		err = errors.Wrapf(err, "scan file %s failed", fcounter.filename)
		return
	}

	return
}

func NewFileNewcontentCounter(filename string) *simpleFileNewcontentCounter {
	return &simpleFileNewcontentCounter{filename: filename}
}
