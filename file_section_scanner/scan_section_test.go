package section_scanner

import (
	"regexp"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestScanSection(t *testing.T) {
	outputs, err := ScanFile("test.txt", sectionFunc)
	if err != nil {
		t.Fatal(err)
	}

	spew.Dump(outputs)
}

func sectionFunc(line string) (status SectionStatus) {
	p := regexp.MustCompile(`^section: (\d+)`)
	matches := p.FindStringSubmatch(line)

	if len(matches) == 0 {
		return NoHitSection
	}

	if matches[1] == "1" || matches[1] == "3" {
		return HitMatchedSection
	}

	return HitNonMatchedSection
}
