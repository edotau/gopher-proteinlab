package stdio

import (
	"bufio"
	"io"
	"os"
)

// Scanalyzer wraps around bufio.Scanner and adds a close method.
type Scanalyzer struct {
	*bufio.Scanner
	close func() error
}

// Close is the method to close the underlying resource, such as a file.
func (s *Scanalyzer) Close() error {
	if s.close != nil {
		return s.close()
	}
	return nil
}

// NewScannerio creates a new Scanalyzer from any io.Reader with an optional close function.
func NewScannerio(r io.Reader, file *os.File) *Scanalyzer {
	return &Scanalyzer{
		Scanner: bufio.NewScanner(r),
		close:   file.Close,
	}
}

// func main() {
//     // Example: Using Scanalyzer to scan lines from a file and ensure the file is closed after use
//     file, err := os.Open("example.txt")
//     if err != nil {
//         fmt.Println("Error opening file:", err)
//         return
//     }

//     // Create a new Scanalyzer with a close function to close the file after scanning
//     scanner := NewScannerio(file, file.Close)

//     // Use the scanner to scan the file line by line
//     for scanner.Scan() {
//         fmt.Println(scanner.Text())
//     }

//     // Check for any scan errors
//     if err := scanner.Err(); err != nil {
//         fmt.Println("Error scanning:", err)
//     }

//     // Close the file after scanning
//     if err := scanner.Close(); err != nil {
//         fmt.Println("Error closing file:", err)
//     }
// }
