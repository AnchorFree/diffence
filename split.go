package diffence

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// split out logic from scan.go
func beginsWithCommitID(s string) bool {
	return true
}

// SplitDiffs splits a single diff txt into an individual DiffItem for each file changed
func SplitDiffs(r io.Reader, l List) error {

	// increase buffer size
	scanner := bufio.NewScanner(r)
	buf := make([]byte, 0, 64*1024)    // 64kb starting token buffer size
	scanner.Buffer(buf, 1024*1024*100) // 100MB max token buffer size: https://help.github.com/articles/working-with-large-files/
	scanner.Split(ScanDiffs)

	// copy to temporary buffer because gets overwritten otherwise
	// https://golang.org/pkg/bufio/#Scanner.Bytes
	// "Bytes returns the most recent token generated by a call to Scan. The underlying array may point to data that will be OVERWRITTEN by a subsequent call to Scan. It does no allocation."
	buffer := bytes.NewBuffer(make([]byte, 0))
	BOF := true

	for scanner.Scan() {
		buffer.Write(scanner.Bytes())
		raw := buffer.String()
		// strip cruft from BOF (Beginning Of File) if necessary
		// if BOF && !strings.HasPrefix(raw, diffSep) {
		if BOF && !(strings.HasPrefix(raw, diffSep) || beginsWithCommitID(raw)) {
			BOF = false
			buffer.Reset()
			continue
		}
		l.Push(raw)
		buffer.Reset()
	}

	return scanner.Err()
}
