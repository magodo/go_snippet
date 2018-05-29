package mylib

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func ReadSlowQueryFromTable(db *sql.DB, outTimeLayout string, startTime, endTime time.Time) (out string, err error) {

	dbTimestampLayout := "2006-01-02 15:04:05"
	dbTimeLayout := "15:04:05"

	/* convert time.Time to DB timestamp type representation */
	startTimeDB := startTime.Format(dbTimestampLayout)
	endTimeDB := endTime.Format(dbTimestampLayout)

	/* query slow log */
	slowQueries, err := Vquery(db, "SELECT * FROM mysql.slow_log WHERE start_time BETWEEN ? and ?", startTimeDB, endTimeDB)
	if err != nil {
		err = errors.Wrap(err, "Failed to query mysql.slow_log")
		return
	}

	/* query log_timestamps to identify whether local time or UTC time should be used in output. */

	var tsType string // either SYSTEM or UTC

	tsTypes, err := Vquery(db, "SHOW VARIABLES LIKE 'log_timestamps'")
	if err != nil {
		err = errors.Wrap(err, "Failed to query 'log_timestamps' system variable")
		return
	}
	if tsTypes == nil {
		/* earlier mysql versions has no "log_timestamps" system var, it uses local time-zone by default*/
		tsType = "SYSTEM"
	} else {
		tsType = tsTypes[0]["Value"]
	}

	/* convert each slow query into file format lines */
	outputSlices := []string{}
	for _, slowq := range slowQueries {
		lines := make([]string, 5)

		// line1: time (local/utc)
		// line4: timestamp
		// db timestamp -> time.Time -> local/utc
		t, err := time.Parse(dbTimestampLayout, slowq["start_time"])
		if err != nil {
			err = errors.Wrapf(err, "Failed to parse db timestamp: %s", slowq["start_time"])
			return "", err
		}
		switch tsType {
		case "SYSTEM":
			t = t.Local()
		case "UTC":
			t = t.UTC()
		}
		lines[0] = fmt.Sprintf("# Time: %s", t.Format(outTimeLayout))
		lines[3] = fmt.Sprintf("SET timestamp=%d;", t.Unix())

		// line2: user, host
		lines[1] = fmt.Sprintf("# User@Host: %s", slowq["user_host"])

		// line3: statistics
		qtime, err := time.Parse(dbTimeLayout, slowq["query_time"])
		if err != nil {
			err = errors.Wrapf(err, "Failed to parse query time: %s", slowq["query_time"])
			return "", err
		}
		ltime, err := time.Parse(dbTimeLayout, slowq["lock_time"])
		if err != nil {
			err = errors.Wrapf(err, "Failed to parse lock time: %s", slowq["lock_time"])
			return "", err
		}
		lines[2] = fmt.Sprintf("Query_time: %d  Lock_time: %d  Rows_sent: %s  Rows_examined: %s", qtime.Second(), ltime.Second(), slowq["rows_sent"], slowq["rows_examined"])

		// line5: msg
		lines[4] = fmt.Sprintf("%s;", slowq["sql_text"])

		// gather
		outputSlices = append(outputSlices, strings.Join(lines, "\n"))
	}

	out = strings.Join(outputSlices, "\n")

	return
}
