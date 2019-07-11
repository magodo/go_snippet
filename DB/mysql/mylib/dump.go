package mylib

import (
	"fmt"
	"regexp"
)

func mysqlDumpQuery(stmt string, args ...interface{}) (fullStmt string, err error) {
	sms := regexp.MustCompile(`\?`).FindAllStringIndex(stmt, -1)
	lastIdx := 0
	out := ""
	i := 0
	for _, sm := range sms {
		startIdx, endIdx := sm[0], sm[1]
		arg := args[i]
		argStr := ""
		switch arg.(type) {
		case string:
			argStr = fmt.Sprintf("'%v'", arg)
		default:
			argStr = fmt.Sprintf("%v", arg)
		}
		out = out + stmt[lastIdx:startIdx] + argStr
		lastIdx = endIdx
		i++
	}
	out = out + stmt[lastIdx:len(stmt)]
	return out, nil
}
