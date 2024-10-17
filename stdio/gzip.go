package stdio

import (
	"bufio"
	"bytes"
)

// IsGzip checks if the file is gzip-compressed by peeking at its magic number
func IsGzip(reader *bufio.Reader) bool {
	if buffer, err := reader.Peek(2); CatchError(err) {
		return bytes.Equal(buffer, []byte{0x1f, 0x8b})
	}
	return false
}
