package openrouter

import (
	"bufio"
	"io"
	"strings"
)

type EventReader struct {
	scanner *bufio.Scanner
}

func NewEventReader(r io.Reader) *EventReader {
	return &EventReader{
		scanner: bufio.NewScanner(r),
	}
}

// ReadEvent reads one SSE block ("data: ...\n" ... "\n")
func (r *EventReader) ReadEvent() (string, error) {
	var builder strings.Builder

	for r.scanner.Scan() {
		line := r.scanner.Text()
		if line == "" {
			break
		}
		if strings.HasPrefix(line, ":") {
			continue
		}
		if after, ok := strings.CutPrefix(line, "data: "); ok {
			builder.WriteString(after)
		}
	}

	if err := r.scanner.Err(); err != nil {
		return "", err
	}

	return builder.String(), nil
}
