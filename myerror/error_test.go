package myerror

import (
	"errors"
	"fmt"
	"testing"
)

func TestOriginErrNil(t *testing.T) {
	var e error = nil
	oe := AppendError(e, errors.New("foo"))
	fmt.Println(oe)
}
func TestOriginErrNotNil(t *testing.T) {
	var e error = errors.New("foo")
	oe := AppendError(e, errors.New("bar"))
	fmt.Println(oe)
}

func TestInputErrNil(t *testing.T) {
	var e error = nil
	oe := AppendError(e, errors.New("foo"), nil, errors.New("bar"), errors.New("foobar"))
	fmt.Println(oe)
}

func TestMultiCall(t *testing.T) {
	var e error = nil
	e = AppendError(e, errors.New("rick"))
	e = AppendError(e, errors.New("morty"))
	fmt.Println(e)
}
