# Go Struct for EMBL Data

We will define structs that capture the key elements from the EMBL format (like `ID`, `AC`, `DE`, `FT`, etc.).

```go
package main

import (
 "bufio"
 "fmt"
 "os"
 "strings"
)

// EMBL Struct Definitions

type EMBLEntry struct {
 ID        string
 Accession []string
 Keywords  []string
 Source    string
 Features  []Feature
 Sequence  string
}

type Feature struct {
 Key       string
 Location  string
 Qualifiers map[string]string
}

```

### Parser for EMBL Format

This function will read an EMBL-formatted file and map the data to the structs.

```go
func parseEMBL(filename string) (*EMBLEntry, error) {
 file, err := os.Open(filename)
 if err != nil {
  return nil, fmt.Errorf("error opening file: %v", err)
 }
 defer file.Close()

 scanner := bufio.NewScanner(file)
 entry := &EMBLEntry{}
 currentFeature := Feature{}

 for scanner.Scan() {
  line := scanner.Text()

  // ID line (entry ID)
  if strings.HasPrefix(line, "ID") {
   entry.ID = strings.TrimSpace(line[5:])
  }

  // AC line (accession numbers)
  if strings.HasPrefix(line, "AC") {
   accessions := strings.Split(line[5:], ";")
   for _, acc := range accessions {
    entry.Accession = append(entry.Accession, strings.TrimSpace(acc))
   }
  }

  // KW line (keywords)
  if strings.HasPrefix(line, "KW") {
   keywords := strings.Split(line[5:], ";")
   for _, keyword := range keywords {
    entry.Keywords = append(entry.Keywords, strings.TrimSpace(keyword))
   }
  }

  // FT line (features)
  if strings.HasPrefix(line, "FT") {
   ftLine := strings.TrimSpace(line[5:])
   if strings.HasPrefix(ftLine, "source") || strings.HasPrefix(ftLine, "CDS") {
    if currentFeature.Key != "" {
     entry.Features = append(entry.Features, currentFeature)
    }
    // New feature
    currentFeature = Feature{
     Key:       strings.Fields(ftLine)[0],
     Location:  strings.Fields(ftLine)[1],
     Qualifiers: make(map[string]string),
    }
   } else {
    // Qualifier line
    if strings.HasPrefix(ftLine, "/") {
     parts := strings.SplitN(ftLine, "=", 2)
     if len(parts) == 2 {
      qualifierKey := strings.TrimSpace(parts[0])
      qualifierValue := strings.Trim(strings.TrimSpace(parts[1]), "\"")
      currentFeature.Qualifiers[qualifierKey] = qualifierValue
     }
    }
   }
  }

  // SQ line (sequence)
  if strings.HasPrefix(line, "SQ") {
   sequence := ""
   for scanner.Scan() {
    seqLine := scanner.Text()
    if seqLine == "//" {
     break
    }
    sequence += strings.TrimSpace(seqLine)
   }
   entry.Sequence = sequence
  }
 }

 // Append the last feature
 if currentFeature.Key != "" {
  entry.Features = append(entry.Features, currentFeature)
 }

 if err := scanner.Err(); err != nil {
  return nil, fmt.Errorf("error reading file: %v", err)
 }

 return entry, nil
}
```

### Main Function to Handle Multiple Formats

Here's how we can add support for multiple formats by calling the appropriate parser based on the `fmt` parameter.

```go
func runParser(fmt, filename string) error {
 switch fmt {
 case "EMBL":
  entry, err := parseEMBL(filename)
  if err != nil {
   return fmt.Errorf("error parsing EMBL file: %v", err)
  }
  fmt.Printf("Parsed EMBL Entry: %+v\n", entry)
 case "GenBank":
  // GenBank parser function to be implemented
 case "DDBJ":
  // DDBJ parser function to be implemented
 default:
  return fmt.Errorf("unsupported format: %s", fmt)
 }
 return nil
}

func main() {
 filename := "embl_sample.txt"
 format := "EMBL" // Change this to "GenBank" or "DDBJ" as needed

 err := runParser(format, filename)
 if err != nil {
  fmt.Printf("Error: %v\n", err)
 }
}
```

This implementation provides the basic framework for parsing **EMBL** data. Let me know when you're ready for the next format, and I can implement it for **GenBank**!
