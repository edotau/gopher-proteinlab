package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
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
func writeFasta(file *fileio.EasyWriter, rec Fasta) {
	var err error
	_, err = fmt.Fprintf(file, ">%s\n", rec.Name)
	parseio.ExitOnError(err)
	for i := 0; i < len(rec.Seq); i += 50 {
		if i+50 > len(rec.Seq) {
			_, err = fmt.Fprintf(file, "%s\n", rec.Seq[i:])
			parseio.ExitOnError(err)
		} else {
			_, err = fmt.Fprintf(file, "%s\n", rec.Seq[i:i+50])
			parseio.ExitOnError(err)
		}
	}
}

func writeTsv(file *fileio.EasyWriter, row string) {
	_, err := fmt.Fprint(file, row)
	parseio.ExitOnError(err)
}

// writeResults reads a XML file from UniProt and writes TSV and FASTA records to respective output files
func writeResults(inputFilename, tableDir, seqDir, outputFileBase string) error {
	xmlReader := parseio.NewCodeReader(inputFilename)
	defer xmlReader.Close()

	// Create TSV file in the tables directory
	tsvFile := fileio.EasyCreate(filepath.Join(tableDir, outputFileBase+".tsv.gz"))
	defer tsvFile.Close()

	// Create FASTA file in the sequences directory
	faFile := fileio.EasyCreate(filepath.Join(seqDir, outputFileBase+".fa.gz"))
	defer faFile.Close()

	writeTsv(tsvFile, "Accession\tDataset\tName\tTaxon\tSequence\n")

	decoder := xml.NewDecoder(xmlReader)

	for {
		entry, err := uniprot.ParseUniProt(decoder)
		if err == io.EOF {
			break // End of file reached
		}
		if err != nil {
			return err // Return any other errors encountered during parsing
		}

		// Write TSV and FASTA entries
		writeTsv(tsvFile, processXml(entry))
		writeFasta(faFile, Fasta{Name: entry.Name, Seq: entry.Sequence.Value})
	}

	return nil
}

// processXml processes each UniProt entry and returns a TSV row string
func processXml(entry *uniprot.Entry) string {
	txt := parseio.NewTxtBuilder()

	// Accession
	txt.WriteString(entry.Accession)
	txt.WriteByte('\t')

	// Dataset
	txt.WriteString(entry.Dataset)
	txt.WriteByte('\t')

	// Name
	txt.WriteString(entry.Name)
	txt.WriteByte('\t')

	// First Taxon in Lineage (for demonstration)
	if len(entry.Organism.Lineage.Taxon) > 0 {
		txt.WriteString(entry.Organism.Lineage.Taxon[0])
	} else {
		txt.WriteString("N/A")
	}
	txt.WriteByte('\t')

	// Sequence
	txt.WriteString(entry.Sequence.Value)
	txt.WriteByte('\n')

	return txt.String()
}

// worker is the function that processes files passed through the jobs channel
func worker(jobs <-chan string, tableDir, seqDir, suffix string, wg *sync.WaitGroup) {
	defer wg.Done()

	for inputFilePath := range jobs {
		outputFileBase := strings.TrimSuffix(filepath.Base(inputFilePath), suffix)

		log.Printf("Processing file: %s", inputFilePath)
		err := writeResults(inputFilePath, tableDir, seqDir, outputFileBase)

		if err != nil {
			log.Printf("Error processing file %s: %v", inputFilePath, err)
		}
	}
}

// processDirectory reads all files from a directory with a given suffix and processes each one concurrently using limited workers
func processDirectory(directory, suffix string) error {
	// Create 'tables' and 'sequences' directories if they don't exist
	tableDir := filepath.Join(directory, "tables")
	seqDir := filepath.Join(directory, "sequences")
	if err := os.MkdirAll(tableDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", tableDir, err)
	}
	if err := os.MkdirAll(seqDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", seqDir, err)
	}

	files, err := os.ReadDir(directory)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", directory, err)
	}

	// Get the number of available CPUs and create a job queue
	numCPUs := runtime.NumCPU()
	jobs := make(chan string, numCPUs)
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < numCPUs; i++ {
		wg.Add(1)
		go worker(jobs, tableDir, seqDir, suffix, &wg)
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
	fmt.Println("  -suffix\tThe file suffix to process (e.g., .xml, .tsv). Default is .xml.")
	fmt.Println("\nExample:")
	fmt.Println("  go run uniprotXml.go -dir=./data -suffix=.xml")
}

func main() {
	// Define command-line flags for the directory and file suffix
	directoryPtr := flag.String("dir", "", "Path to the directory containing files")
	suffixPtr := flag.String("suffix", ".xml", "File suffix to process (e.g., .xml, .tsv)")

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
