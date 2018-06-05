package mylib

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func TimestampReprToTime(tsRepr string) (t time.Time, err error) {
	ts, err := strconv.ParseFloat(tsRepr, 64)
	if err != nil {
		err = errors.Wrapf(err, "Failed to convert timestamp representation %s to float64", tsRepr)
		return
	}
	t = time.Unix(0, int64(ts*1e9))
	return
}
