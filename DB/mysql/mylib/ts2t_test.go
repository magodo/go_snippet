package mylib

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestTs2t(t *testing.T) {
	now := time.Now()

	ts := strconv.FormatFloat(float64(now.UnixNano())/1e9, 'f', -1, 64)
	time_, err := TimestampReprToTime(ts)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(time_)
	fmt.Println(now)

}
