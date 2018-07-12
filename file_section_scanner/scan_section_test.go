package section_scanner

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestScanSection(t *testing.T) {
	outputs, err := ScanFile("test.txt", sectionFunc)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(len(outputs))
	fmt.Println(strings.Join(outputs, ""))
}

func sectionFunc(line string) (status SectionStatus) {
	p := regexp.MustCompile(`^section: (\d+)`)
	matches := p.FindStringSubmatch(line)

	if len(matches) == 0 {
		return NoHitSection
	}

	if matches[1] == "1" || matches[1] == "3" || matches[1] == "4" || matches[1] == "5" {
		//if matches[1] == "1" {
		return HitMatchedSection
	}

	return HitNonMatchedSection
}
