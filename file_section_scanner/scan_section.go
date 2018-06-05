package section_scanner

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type SectionStatus int

const (
	// hit a new section satisfying criteria
	HitMatchedSection SectionStatus = iota
	// hit a new section not satisfying criteria
	HitNonMatchedSection
	// no hit
	NoHitSection
)

//SectionStatusFunc determine if a new section is hit based on the current scanning line
type SectionStatusFunc func(string) SectionStatus

func ScanFile(path string, secFunc SectionStatusFunc) (contents []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		err = errors.Wrapf(err, "Failed to open: %s", path)
		return
	}
	r := bufio.NewReader(f)

	isInSection := false
	lineOfSection := []string{}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if isInSection {
					// finish this section and return
					contents = append(contents, strings.Join(lineOfSection, "\n"))
					return contents, nil
				}
			} else {
				err = errors.Wrapf(err, "Failed to read string up to new line: %s", err)
				return nil, err
			}
		}

		switch secFunc(line) {
		case HitMatchedSection:
			if isInSection {
				// finish last section
				contents = append(contents, strings.Join(lineOfSection, "\n"))
			}
			// add this line into new section
			lineOfSection = []string{line}
			isInSection = true

		case HitNonMatchedSection:
			if isInSection {
				// finish last section
				contents = append(contents, strings.Join(lineOfSection, "\n"))
			}
			isInSection = false

		case NoHitSection:
			if isInSection {
				lineOfSection = append(lineOfSection, line)
			}
		}

	}
}
