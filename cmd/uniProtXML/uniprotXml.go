package main

import (
	"encoding/xml"
	"io"

	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"gopher-proteinlab/parseio"
	"gopher-proteinlab/uniprot"

	"github.com/vertgenlab/gonomics/fileio"
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
	parseio.ExitOnError(err)
	for i := 0; i < len(rec.Seq); i += lineLength {
		if i+lineLength > len(rec.Seq) {
			_, err = fmt.Fprintf(file, "%s\n", rec.Seq[i:])
			parseio.ExitOnError(err)
		} else {
			_, err = fmt.Fprintf(file, "%s\n", rec.Seq[i:i+lineLength])
			parseio.ExitOnError(err)
		}
	}
}

func writeTsv(file *fileio.EasyWriter, row string) {
	_, err := fmt.Fprint(file, row)
	parseio.ExitOnError(err)
}

// writeResults reads a XML file from UniProtand writes FASTA records directly to the output file
func writeResults(inputFilename, outputFilePath string) error {
	xmlReader := parseio.NewCodeReader(inputFilename)
	defer xmlReader.Close()
	
	tsvFile := fileio.EasyCreate(outputFilePath+".tsv.gz")
	writeTsv(tsvFile, "Accession\tDataset\tName\tSequence\n")
	
	decoder := xml.NewDecoder(xmlReader)
	for {
		// Parse each entry and process it
		entry, err := uniprot.ParseUniProt(decoder)
		if err == io.EOF {
			break // End of file reached
		}
		writeTsv(tsvFile, processXml(entry))
	}
	
	writeTsv(tsvFile, "\n")
	tsvFile.Close()
	return nil
}

// processXml
func processXml(entry *uniprot.Entry) string {
	txt := parseio.NewTxtBuilder()
	// Accession
	txt.WriteString(entry.Accession)
	txt.WriteByte('\t')

	txt.WriteString(entry.Dataset)
	txt.WriteByte('\t')
	// Name
	txt.WriteString(entry.Name)
	txt.WriteByte('\t')
	// Taxon
	txt.WriteString(entry.Organism.Lineage.Taxon[0])
	txt.WriteByte('\t')
	// Sequence
	txt.WriteString(entry.Sequence.Value)

	txt.WriteByte('\n')
	return txt.String()
}

// func processTsv(inputFilename, outputFilename string) error {
// 	reader := fileio.EasyOpen(inputFilename)
// 	defer reader.Close()

// 	outputFile := fileio.EasyCreate(outputFilename)
// 	defer outputFile.Close()

// 	// Read the header line and validate
// 	line, done := fileio.EasyNextLine(reader)
// 	if done {
// 		return fmt.Errorf("Error: File %s is empty", inputFilename)
// 	}

// 	header := strings.Split(line, "\t")
// 	fmt.Printf("Parsed header: %v\n", header) // Debugging line to print the header

// 	if len(header) < 3 {
// 		return fmt.Errorf("Error: Invalid header format in %s, expected 'Id\tName\tSequence'", inputFilename)
// 	}

// 	// Process each line and write FASTA records directly
// 	for line, done = fileio.EasyNextLine(reader); !done; line, done = fileio.EasyNextLine(reader) {
// 		row := strings.Split(line, "\t")
// 		if len(row) != 3 {
// 			return fmt.Errorf("Error: Invalid TSV format in %s - each row should have 3 columns", inputFilename)
// 		}
// 		writeFasta(outputFile, Fasta{Name: row[1], Seq: row[2]}, 50) // Write with a line length of 80 characters
// 	}

// 	return nil
// }

// worker function that processes files passed through the jobs channel
func worker(jobs <-chan string, suffix string, wg *sync.WaitGroup) {
	defer wg.Done()

	for inputFilePath := range jobs {
		outputFilePath := strings.TrimSuffix(filepath.Base(inputFilePath), suffix)
		writeResults(inputFilePath, outputFilePath)
		
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
	fmt.Println("Usage: go run uniprotXml.go -dir=<directory> [-suffix=<suffix>]")
	fmt.Println("\nOptions:")
	fmt.Println("  -dir\t\tThe path to the directory containing files to process.")
	fmt.Println("  -suffix\tThe file suffix to process (e.g., .tsv, .txt). Default is .tsv.")
	fmt.Println("\nExample:")
	fmt.Println("  go run uniprotXml.go -dir=./data -suffix=.txt")
}

func main() {
	// Define command-line flags for the directory and file suffix
	directoryPtr := flag.String("dir", "", "Path to the directory containing files")
	suffixPtr := flag.String("suffix", ".xml", "File suffix to process (e.g., .xml, .tsv, .txt)")

	// Override the default usage message
	flag.Usage = usage

	flag.Parse()

	// Check if the directory flag is provided
	if *directoryPtr == "" {
		flag.Usage()
		log.Fatal("Error: Please provide the directory path using the -dir flag")
	}

	if *suffixPtr == "" {
		flag.Usage()
		log.Fatal("Error: Please provide the suffix string using the -suffix flag")
	}

	// Process the directory and handle any errors
	if err := processDirectory(*directoryPtr, *suffixPtr); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
