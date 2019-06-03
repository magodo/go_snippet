package dump

import (
	"fmt"
	"regexp"
	"strconv"
)

func showPostgreSQLStmt(stmt string, args ...interface{}) (string, error) {
	sms := regexp.MustCompile(`\$\d+`).FindAllStringIndex(stmt, -1)
	firstIdx := 0
	out := ""
	for _, sm := range sms {
		startIdx, endIdx := sm[0], sm[1]
		argIdx, err := strconv.Atoi(stmt[startIdx+1 : endIdx])
		if err != nil {
			return "", err
		}
		arg := args[argIdx-1]
		argStr := ""
		switch arg.(type) {
		case string:
			argStr = fmt.Sprintf("'%v'", arg)
		default:
			argStr = fmt.Sprintf("%v", arg)
		}
		out = out + stmt[firstIdx:startIdx] + argStr
		firstIdx = endIdx
	}
	return out, nil
}
