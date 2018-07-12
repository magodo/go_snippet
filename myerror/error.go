package myerror

import "fmt"

func AppendError(err error, ierrs ...error) (oerr error) {
	oerr = err
	if len(ierrs) == 0 {
		return
	}

	for _, e := range ierrs {
		if e == nil {
			continue
		}
		if oerr == nil {
			oerr = e
			continue
		}
		oerr = fmt.Errorf("%s\nWith another error: %s", oerr.Error(), e.Error())
	}
	return
}
