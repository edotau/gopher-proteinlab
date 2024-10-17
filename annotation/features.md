# Struct for Feature Keys and Qualifiers

We'll extend the existing `Feature` struct to support more complex qualifiers and integrate the mandatory/optional handling of feature keys.

```go
// Entry struct represents a generic structure for parsed data from EMBL, GenBank, and DDBJ
type Entry struct {
 ID        string
 Accession []string
 Keywords  []string
 Source    string
 Features  []Feature
 Sequence  string
}

// Feature struct to store feature qualifiers and their parsed values
type Feature struct {
 Allele        string
 Altitude      string
 Anticodon     Anticodon
 BioMaterial   string
 Replace       string
 TranslExcept  []TranslExcept
}

// Anticodon struct to hold the structured data for the /anticodon qualifier
type Anticodon struct {
 Location  string
 AminoAcid string
 Sequence  string
}

// TranslExcept struct to hold the structured data for the /transl_except qualifier
type TranslExcept struct {
 Position  string
 AminoAcid string
}

```

### Generic Feature Parsing

A function that can handle parsing feature tables across all formats, with specific logic depending on the `fmt` (EMBL, GenBank, or DDBJ).

```go
func parseFeatures(scanner *bufio.Scanner, fmt string) ([]Feature, error) {
 var features []Feature
 var currentFeature Feature

 for scanner.Scan() {
  line := scanner.Text()
  if strings.HasPrefix(line, "FT") || strings.HasPrefix(line, "FEATURES") {
   ftLine := strings.TrimSpace(line[5:])
   if ftLine != "" {
    // Handle a new feature key
    parts := strings.Fields(ftLine)
    if len(parts) > 1 {
     // Append the last feature before processing a new one
     if currentFeature.Key != "" {
      features = append(features, currentFeature)
     }
     // Start a new feature
     currentFeature = Feature{
      Key:        parts[0],
      Location:   parts[1],
      Qualifiers: make(map[string]string),
     }
    } else if strings.HasPrefix(ftLine, "/") {
     // Handle qualifiers
     qualifierParts := strings.SplitN(ftLine, "=", 2)
     if len(qualifierParts) == 2 {
      qualKey := strings.TrimSpace(qualifierParts[0])
      qualValue := strings.Trim(strings.TrimSpace(qualifierParts[1]), "\"")
      currentFeature.Qualifiers[qualKey] = qualValue
     }
    }
   }
  }
 }

 // Append the final feature
 if currentFeature.Key != "" {
  features = append(features, currentFeature)
 }

 if err := scanner.Err(); err != nil {
  return nil, fmt.Errorf("error parsing features: %v", err)
 }

 return features, nil
}

```

### Parsing Function for GenBank

This parser handles the specific layout of GenBank entries, including fields like `LOCUS`, `DEFINITION`, and `FEATURES`.

```go
func parseGenBank(filename string) (*GenBankEntry, error) {
 file, err := os.Open(filename)
 if err != nil {
  return nil, fmt.Errorf("error opening file: %v", err)
 }
 defer file.Close()

 scanner := bufio.NewScanner(file)
 entry := &GenBankEntry{}
 features := []Feature{}
 currentFeature := Feature{}

 for scanner.Scan() {
  line := scanner.Text()

  // LOCUS line
  if strings.HasPrefix(line, "LOCUS") {
   entry.Source = strings.TrimSpace(line[12:])
  }

  // DEFINITION line
  if strings.HasPrefix(line, "DEFINITION") {
   entry.Definition = strings.TrimSpace(line[12:])
  }

  // FEATURES line (Start of feature parsing)
  if strings.HasPrefix(line, "FEATURES") {
   ftFeatures, err := parseFeatures(scanner, "GenBank")
   if err != nil {
    return nil, err
   }
   features = append(features, ftFeatures...)
  }
 }

 if err := scanner.Err(); err != nil {
  return nil, fmt.Errorf("error reading file: %v", err)
 }

 fmt.Printf("Parsed GenBank Entry with %d features\n", len(features))
 return entry, nil
}
```

### Parsing Function for DDBJ

DDBJ format is very similar to GenBank, but we provide a dedicated parser for consistency.

```go
func parseDDBJ(filename string) (*DDBJEntry, error) {
 file, err := os.Open(filename)
 if err != nil {
  return nil, fmt.Errorf("error opening file: %v", err)
 }
 defer file.Close()

 scanner := bufio.NewScanner(file)
 entry := &DDBJEntry{}
 features := []Feature{}

 for scanner.Scan() {
  line := scanner.Text()

  // LOCUS line
  if strings.HasPrefix(line, "LOCUS") {
   entry.Source = strings.TrimSpace(line[12:])
  }

  // DEFINITION line
  if strings.HasPrefix(line, "DEFINITION") {
   entry.Definition = strings.TrimSpace(line[12:])
  }

  // FEATURES line (Start of feature parsing)
  if strings.HasPrefix(line, "FEATURES") {
   ftFeatures, err := parseFeatures(scanner, "DDBJ")
   if err != nil {
    return nil, err
   }
   features = append(features, ftFeatures...)
  }
 }

 if err := scanner.Err(); err != nil {
  return nil, fmt.Errorf("error reading file: %v", err)
 }

 fmt.Printf("Parsed DDBJ Entry with %d features\n", len(features))
 return entry, nil
}
```

### Main Parser Function

The main function now handles all three formats—`EMBL`, `GenBank`, and `DDBJ`. It dispatches the appropriate parser based on the `fmt` parameter.

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
  entry, err := parseGenBank(filename)
  if err != nil {
   return fmt.Errorf("error parsing GenBank file: %v", err)
  }
  fmt.Printf("Parsed GenBank Entry: %+v\n", entry)

 case "DDBJ":
  entry, err := parseDDBJ(filename)
  if err != nil {
   return fmt.Errorf("error parsing DDBJ file: %v", err)
  }
  fmt.Printf("Parsed DDBJ Entry: %+v\n", entry)

 default:
  return fmt.Errorf("unsupported format: %s", fmt)
 }
 return nil
}

func main() {
 filename := "genbank_sample.txt"
 format := "GenBank" // Change to "EMBL" or "DDBJ" as needed

 err := runParser(format, filename)
 if err != nil {
  fmt.Printf("Error: %v\n", err)
 }
}
```

### Feature Key Validation (Optional Enhancement)

You can add a validation layer inside the `parseFeatures` function to check that each feature key has the correct mandatory and optional qualifiers based on the reference provided.

For instance:

```go
func validateFeatureKey(feature Feature) error {
 // Add validation logic for mandatory qualifiers
 // e.g. CDS must have /gene, /product, etc.
 switch feature.Key {
 case "CDS":
  if _, ok := feature.Qualifiers["gene"]; !ok {
   return fmt.Errorf("CDS feature missing mandatory qualifier: gene")
  }
  // Add more validation as needed for other feature keys
 }
 return nil
}
```
