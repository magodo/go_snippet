package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/pkg/errors"
)

var device string

func init() {
	flag.StringVar(&device, "d", "", "device path")
	flag.Parse()
}

func main() {
	isssd, err := isSSD(device)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s is ssd: %v\n", device, isssd)
}

func isSSD(device string) (result bool, err error) {
	const BLKROTATIONAL uint = 0x12<<8 + 126
	fd, err := os.Open(device)
	if err != nil {
		err = errors.Wrapf(err, "os.Open failed for %s", device)
		return
	}

	isRotational := 0
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd.Fd(), uintptr(BLKROTATIONAL), uintptr(unsafe.Pointer(&isRotational)))
	if errno != syscall.Errno(0) {
		err = errors.New("syscall.Syscall BLKROTATIONAL failed")
		return
	}

	if isRotational == 0 {
		result = true
	} else {
		result = false
	}
	return
}
