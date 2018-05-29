package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

// check new logs(compared to last access) written into /var/log/messages
type MessagesContentChecker struct {
	Filename     string
	Stat         syscall.Stat_t
	RotatePeriod int // unit: days
}

var messagesContentchecker = &MessagesContentChecker{Filename: "messages", RotatePeriod: 7}

func (checker *MessagesContentChecker) CountNewContentContain(pattern string) (hit int, err error) {

	p, err := regexp.Compile(pattern)
	if err != nil {
		err = errors.Wrapf(err, "regexp compile pattern %s failed", pattern)
		return
	}

	// get file metadata
	oldStat := checker.Stat
	err = syscall.Stat(checker.Filename, &checker.Stat)
	if err != nil {
		// case 1: If file is removed after a previous successful check, then it means the /var/log/messages is logrotating, but the new file hasn't been created yet.
		// Need to reset stat (prepare for next call and avoid nex call to enter this condition again) and check the last rotated file, from the position we accessed last time.
		if err == syscall.ENOENT && oldStat.Ino != 0 {
			fmt.Printf("%s not exists, it might because the log file is rotating, but new file hasn't been created yet.\n", checker.Filename)
			checker.Stat = syscall.Stat_t{}

			rotatedFilename := checker.Filename + "-" + time.Now().AddDate(0, 0, -checker.RotatePeriod).Format("20060102")
			hit, err = countFileContain(rotatedFilename, oldStat.Size, p)
			return
		}
		// some unexpected error occurs
		err = errors.Wrapf(err, "stat %s failed", checker.Filename)
		return
	}

	if oldStat.Ino != checker.Stat.Ino {
		fmt.Printf("A new %s is created by logrotate\n", checker.Filename)
		// case 2: If a same-named file is created, which means the /var/log/messasges is logrotating, and the new file has been created.
		// Need to check both:
		// - the last rotated file, from position we accessed it last time
		// - the new file, from beginning

		// last rotated file only when we have not yet check it yet
		// (if we have checked it at case 1, then we should not check it again here.)
		if oldStat.Ino != 0 {
			rotatedFilename := checker.Filename + "-" + time.Now().AddDate(0, 0, -checker.RotatePeriod).Format("20060102")
			nOld, err := countFileContain(rotatedFilename, oldStat.Size, p)
			if err != nil {
				err = errors.Wrapf(err, "Failed to count new file content for %s", rotatedFilename)
				return 0, err
			}
			hit += nOld
		}

		// the new file
		nNew, err := countFileContain(checker.Filename, 0, p)
		if err != nil {
			err = errors.Wrapf(err, "Failed to count new file content for %s", checker.Filename)
			return 0, err
		}
		hit += nNew
		return hit, nil
	}

	if oldStat.Size > checker.Stat.Size {
		// case 3: if file is truncated, no need to check
		err = errors.New("File is trucated, nothing to check")
		return
	}

	// normal case
	return countFileContain(checker.Filename, oldStat.Size, p)
}

// Init will just get the stat of target file
func (checker *MessagesContentChecker) Init() (err error) {

	err = syscall.Stat(checker.Filename, &checker.Stat)
	if err != nil {
		err = errors.Wrap(err, "Init MessagesContentChecker failed (maybe called during logratating)")
		return
	}
	return
}

// Count how many times does "pattern" occur in "filename" from "startPoint" to end.
func countFileContain(filename string, startPoint int64, pattern *regexp.Regexp) (hit int, err error) {
	f, err := os.Open(filename)
	if err != nil {
		err = errors.Wrapf(err, "failed to open %s", filename)
		return
	}
	defer f.Close()
	// seek to where we read last time
	_, err = f.Seek(startPoint, 0)
	if err != nil {
		err = errors.Wrapf(err, "failed to seek to %d for file %s", startPoint, filename)
		return
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if pattern.Match(scanner.Bytes()) {
			hit++
		}
	}
	err = scanner.Err()
	return
}

func main() {
	err := messagesContentchecker.Init()
	if err != nil {
		log.Fatal(err)
	}
	for {
		n, err := messagesContentchecker.CountNewContentContain("oom")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
		time.Sleep(2 * time.Second)
	}
}
