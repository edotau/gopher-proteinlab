package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"gopher-proteinlab/parseio"
	"gopher-proteinlab/uniprot"
	// "github.com/vertgenlab/gonomics/fileio"
)

// Fasta struct representing the Name and Sequence for each FASTA entry
// type Fasta struct {
// 	Name string
// 	Seq  string
// }

// // writeFasta writes a single FASTA record to the provided writer in a wrapped line format
// func writeFasta(file *fileio.EasyWriter, rec Fasta, lineLength int) {
// 	var err error
// 	_, err = fmt.Fprintf(file, ">%s\n", rec.Name)
// 	parseio.ExitOnError(err)
// 	for i := 0; i < len(rec.Seq); i += lineLength {
// 		if i+lineLength > len(rec.Seq) {
// 			_, err = fmt.Fprintf(file, "%s\n", rec.Seq[i:])
// 			parseio.ExitOnError(err)
// 		} else {
// 			_, err = fmt.Fprintf(file, "%s\n", rec.Seq[i:i+lineLength])
// 			parseio.ExitOnError(err)
// 		}
// 	}
// }

// processXml reads a XML file from UniProtand writes FASTA records directly to the output file
func processXml(entry *uniprot.Entry) string {
	txt := parseio.NewTxtBuilder()
	// Accession
	txt.WriteString(entry.Accession[0])
	txt.WriteByte('\t')
	// Name
	txt.WriteString(entry.Name[0])
	txt.WriteByte('\t')
	// Taxon
	txt.WriteString(entry.Organism.Lineage.Taxon[0])
	txt.WriteByte('\t')
	// Sequence
	txt.WriteString(entry.Sequence.Value)
	txt.WriteByte('\t')
	txt.WriteByte('\n')
	return txt.String()
}

// worker function that processes files passed through the jobs channel
func worker(jobs <-chan string, suffix string, wg *sync.WaitGroup) {
	defer wg.Done()

	for inputFilePath := range jobs {
		xmlReader := parseio.NewCodeReader(inputFilePath)
		defer xmlReader.Close()

		// Iterate through the tokens in the XML file
		decoder := xml.NewDecoder(xmlReader)
		
		for entry, err := uniprot.ParseUniProt(decoder); err != io.EOF; entry, err =  uniprot.ParseUniProt(decoder) {
			column := processXml(entry)
			fmt.Println(column)
		}
		//outputFilePath := strings.TrimSuffix(inputFilePath, suffix) + ".fa.gz"
		
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
	// directoryPtr := flag.String("dir", "", "Path to the directory containing files")
	filePath := flag.String("file", "", "Path to the single xml file")
	// suffixPtr := flag.String("suffix", ".xml", "File suffix to process (e.g., .xml, .tsv, .txt)")

	// Override the default usage message
	flag.Usage = usage

	flag.Parse()

	// Check if the directory flag is provided
	// if *directoryPtr == "" {
	// 	flag.Usage()
	// 	log.Fatal("Error: Please provide the directory path using the -dir flag")
	// }

	xmlReader := parseio.NewCodeReader(*filePath)
	defer xmlReader.Close()

	// Iterate through the tokens in the XML file
	decoder := xml.NewDecoder(xmlReader)
	
	for entry, err := uniprot.ParseUniProt(decoder); err != io.EOF; entry, err =  uniprot.ParseUniProt(decoder) {
		column := processXml(entry)
		fmt.Println(column)
	}


	// Process the directory and handle any errors
	// if err := processDirectory(*directoryPtr, *suffixPtr); err != nil {
	// 	log.Fatalf("Error: %v", err)
	// }
}
