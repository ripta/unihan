package generate

import (
	"archive/zip"
	"bufio"
	"fmt"
	"strings"

	"github.com/ripta/unihan/pkg/definitions"
)

func ParseUnihanReadings(zr *zip.Reader) (definitions.UnihanReadings, error) {
	r, err := getFile(zr, "Unihan_Readings.txt")
	if err != nil {
		return nil, err
	}

	defer r.Close()

	u := map[rune]definitions.UnihanReading{}

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := s.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}

		fields := strings.Split(line, "\t")
		if len(fields) != 3 {
			continue
		}

		var codepoint rune
		if _, err := fmt.Sscanf(fields[0], "%U", &codepoint); err != nil {
			return nil, fmt.Errorf("parsing codepoint %q: %w", fields[0], err)
		}

		fieldName := fields[1][1:]
		fieldPos, ok := definitions.UnihanReadingKeyMap[fieldName]
		if !ok {
			return nil, fmt.Errorf("invalid field name %q", fieldName)
		}

		value := fields[2]

		if v, ok := u[codepoint]; ok {
			v[fieldPos] = value
			u[codepoint] = v
		} else {
			v := definitions.UnihanReading{}
			v[fieldPos] = value
			u[codepoint] = v
		}
	}

	return u, s.Err()
}
