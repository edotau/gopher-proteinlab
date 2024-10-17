package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/vertgenlab/gonomics/fileio"
	"gopher-proteinlab/stdio"
)

// Fasta struct representing the Name and Sequence for each FASTA entry
type Fasta struct {
	Name string
	Seq  string
}

// writeFasta writes a single FASTA record to the provided writer in a wrapped line format
func writeFasta(file *fileio.EasyWriter, rec Fasta, lineLength int) {
	var err error
	_, err = fmt.Fprintf(file, ">%s\n", rec.Name)
	stdio.CatchError(err)
	for i := 0; i < len(rec.Seq); i += lineLength {
		if i+lineLength > len(rec.Seq) {
			_, err = fmt.Fprintf(file, "%s\n", rec.Seq[i:])
			stdio.CatchError(err)
		} else {
			_, err = fmt.Fprintf(file, "%s\n", rec.Seq[i:i+lineLength])
			stdio.CatchError(err)
		}
	}
}

// processTsv reads a TSV file and writes FASTA records directly to the output file
func processTsv(inputFilename, outputFilename string) error {
	reader := fileio.EasyOpen(inputFilename)
	defer reader.Close()

	outputFile := fileio.EasyCreate(outputFilename)
	defer outputFile.Close()

	// Read the header line and validate
	line, done := fileio.EasyNextLine(reader)
	if done {
		return fmt.Errorf("Error: File %s is empty", inputFilename)
	}

	header := strings.Split(line, "\t")
	fmt.Printf("Parsed header: %v\n", header) // Debugging line to print the header

	if len(header) < 3 {
		return fmt.Errorf("Error: Invalid header format in %s, expected 'Id\tName\tSequence'", inputFilename)
	}

	// Process each line and write FASTA records directly
	for line, done = fileio.EasyNextLine(reader); !done; line, done = fileio.EasyNextLine(reader) {
		row := strings.Split(line, "\t")
		if len(row) != 3 {
			return fmt.Errorf("Error: Invalid TSV format in %s - each row should have 3 columns", inputFilename)
		}
		writeFasta(outputFile, Fasta{Name: row[1], Seq: row[2]}, 50) // Write with a line length of 80 characters
	}

	return nil
}

// worker function that processes files passed through the jobs channel
func worker(jobs <-chan string, suffix string, wg *sync.WaitGroup) {
	defer wg.Done()

	for inputFilePath := range jobs {
		outputFilePath := strings.TrimSuffix(inputFilePath, suffix) + ".fa.gz"
		if err := processTsv(inputFilePath, outputFilePath); err != nil {
			log.Printf("Error processing file %s: %v", inputFilePath, err)
		}
	}
}

// processDirectory reads all files from a directory with a given suffix, processes each one concurrently with limited CPU workers
func processDirectory(directory, suffix string) error {
	files, err := os.ReadDir(directory)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", directory, err)
	}

	// Get the number of available CPUs and limit the number of goroutines
	numCPUs := runtime.NumCPU()
	jobs := make(chan string, numCPUs)
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numCPUs; i++ {
		wg.Add(1)
		go worker(jobs, suffix, &wg)
	}

	// Enqueue jobs (files) for the workers
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), suffix) {
			continue
		}
		inputFilePath := filepath.Join(directory, file.Name())
		jobs <- inputFilePath
	}

	// Close the jobs channel and wait for all workers to finish
	close(jobs)
	wg.Wait()

	return nil
}

func usage() {
	fmt.Println("Usage: go run main.go -dir=<directory> [-suffix=<suffix>]")
	fmt.Println("\nOptions:")
	fmt.Println("  -dir\t\tThe path to the directory containing files to process.")
	fmt.Println("  -suffix\tThe file suffix to process (e.g., .tsv, .txt). Default is .tsv.")
	fmt.Println("\nExample:")
	fmt.Println("  go run main.go -dir=./data -suffix=.txt")
}

func main() {
	// Define command-line flags for the directory and file suffix
	directoryPtr := flag.String("dir", "", "Path to the directory containing files")
	suffixPtr := flag.String("suffix", ".tsv", "File suffix to process (e.g., .tsv, .txt)")

	// Override the default usage message
	flag.Usage = usage

	flag.Parse()

	// Check if the directory flag is provided
	if *directoryPtr == "" {
		flag.Usage()
		log.Fatal("Error: Please provide the directory path using the -dir flag")
	}

	// Process the directory and handle any errors
	if err := processDirectory(*directoryPtr, *suffixPtr); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
