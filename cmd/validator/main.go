package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	generated "github.com/aequo-labs/forgexml-scap/internal/generated/gov/nist/scap/schema/asset-reporting-format/1-1"
)

var (
	inputPath  = flag.String("input", "", "Input XML file or directory (required)")
	outputPath = flag.String("output", "", "Output XML file or directory (optional, default: in-memory only)")
	reportFile = flag.String("report", "", "Path to write aggregate report for directory processing (optional)")
	listXSD    = flag.Bool("list-xsd", false, "List all embedded XSD files")
	showXSD    = flag.String("show-xsd", "", "Show content of specific XSD file")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nValidation Modes:\n")
		fmt.Fprintf(os.Stderr, "  Single file:  --input file.xml [--output out.xml]\n")
		fmt.Fprintf(os.Stderr, "  Directory:    --input ./dir [--output ./outdir] [--report summary.txt]\n")
		fmt.Fprintf(os.Stderr, "\nFlags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  # Validate single file (in-memory, no output file)\n")
		fmt.Fprintf(os.Stderr, "  %s --input file.xml\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  # Validate and write reserialized file\n")
		fmt.Fprintf(os.Stderr, "  %s --input file.xml --output result.xml\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  # Validate all XML files in directory (in-memory)\n")
		fmt.Fprintf(os.Stderr, "  %s --input ./arf-files --report summary.txt\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  # Validate directory and write all reserialized files\n")
		fmt.Fprintf(os.Stderr, "  %s --input ./arf-files --output ./validated --report summary.txt\n\n", os.Args[0])
	}

	flag.Parse()

	// Handle XSD file operations
	if *listXSD {
		listXSDFiles()
		return
	}

	if *showXSD != "" {
		showXSDFile(*showXSD)
		return
	}

	// Validate required input flag
	if *inputPath == "" {
		fmt.Fprintln(os.Stderr, "Error: --input flag is required")
		flag.Usage()
		os.Exit(1)
	}

	// Check if input is file or directory
	info, err := os.Stat(*inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: cannot access input path: %v\n", err)
		os.Exit(1)
	}

	if info.IsDir() {
		// Directory processing
		if err := runBatchValidation(*inputPath, *outputPath, *reportFile); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Single file processing
		writeOutput := *outputPath != ""
		outFile := *outputPath
		if outFile == "" {
			// In-memory processing, but need a path for display
			ext := filepath.Ext(*inputPath)
			base := strings.TrimSuffix(*inputPath, ext)
			outFile = base + "_reserialized.xml"
		}

		if err := runRoundTripTest(*inputPath, outFile, writeOutput); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}
}

// listXSDFiles lists all embedded XSD files
func listXSDFiles() {
	fmt.Println("Embedded XSD Files:")
	fmt.Println("==================")

	err := fs.WalkDir(generated.Schemas, "schemas", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(path, ".xsd") {
			fmt.Println(path)
		}
		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing XSD files: %v\n", err)
		os.Exit(1)
	}
}

// showXSDFile displays the content of a specific XSD file
func showXSDFile(filename string) {
	// Try exact path first
	content, err := fs.ReadFile(generated.Schemas, filename)
	if err != nil {
		// Try with schemas/ prefix
		if !strings.HasPrefix(filename, "schemas/") {
			content, err = fs.ReadFile(generated.Schemas, "schemas/"+filename)
		}
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading XSD file '%s': %v\n", filename, err)
		fmt.Fprintln(os.Stderr, "\nUse --list-xsd to see available files")
		os.Exit(1)
	}

	fmt.Println(string(content))
}

// BatchResult holds validation results for a single file
type BatchResult struct {
	InputFile        string
	Success          bool
	ErrorMessage     string
	FidelityPercent  float64
	Verdict          string // "PERFECT", "SEMANTIC", "FAILURE"
	InputElements    int
	OutputElements   int
	InputAttributes  int
	OutputAttributes int
}

// BatchSummary holds aggregate results for directory processing
type BatchSummary struct {
	TotalFiles       int
	SuccessfulFiles  int
	FailedFiles      int
	PerfectFidelity  int
	SemanticFidelity int
	Results          []BatchResult
}

// runBatchValidation processes all XML files in a directory
func runBatchValidation(inputDir, outputDir, reportPath string) error {
	fmt.Printf("=== Batch Validation ===\n")
	fmt.Printf("Input Directory:  %s\n", inputDir)

	writeOutput := outputDir != ""
	if writeOutput {
		fmt.Printf("Output Directory: %s\n", outputDir)
		// Create output directory if it doesn't exist
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	} else {
		fmt.Printf("Output Directory: (in-memory only)\n")
	}

	if reportPath != "" {
		fmt.Printf("Report File:      %s\n", reportPath)
	}
	fmt.Println()

	// Find all XML files
	var xmlFiles []string
	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".xml") {
			xmlFiles = append(xmlFiles, path)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to scan directory: %w", err)
	}

	if len(xmlFiles) == 0 {
		return fmt.Errorf("no XML files found in %s", inputDir)
	}

	fmt.Printf("Found %d XML files to validate\n\n", len(xmlFiles))

	// Set up worker pool for parallel processing
	numWorkers := runtime.NumCPU()
	fmt.Printf("Using %d parallel workers (CPU count)\n\n", numWorkers)

	// Channels for distributing work and collecting results
	type job struct {
		index   int
		xmlFile string
	}

	jobs := make(chan job, len(xmlFiles))
	results := make(chan BatchResult, len(xmlFiles))

	// Progress tracking
	var processedCount int64
	var progressMutex sync.Mutex

	// Start workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := range jobs {
				relPath, _ := filepath.Rel(inputDir, j.xmlFile)

				// Thread-safe progress update
				current := atomic.AddInt64(&processedCount, 1)
				progressMutex.Lock()
				fmt.Printf("[%d/%d] Worker %d processing: %s\n", current, len(xmlFiles), workerID, relPath)
				progressMutex.Unlock()

				result := BatchResult{
					InputFile: relPath,
				}

				// Determine output path
				var outPath string
				if writeOutput {
					outPath = filepath.Join(outputDir, relPath)
					// Create parent directory for output file
					if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
						result.Success = false
						result.ErrorMessage = fmt.Sprintf("failed to create output directory: %v", err)
						results <- result
						continue
					}
				} else {
					// Generate temp path for display purposes
					ext := filepath.Ext(j.xmlFile)
					base := strings.TrimSuffix(j.xmlFile, ext)
					outPath = base + "_reserialized.xml"
				}

				// Run validation
				validationResult, fidelity, verdict, err := runQuietValidation(j.xmlFile, outPath, writeOutput)
				if err != nil {
					result.Success = false
					result.ErrorMessage = err.Error()
				} else {
					result.Success = true
					result.FidelityPercent = fidelity
					result.Verdict = verdict
					result.InputElements = validationResult.InputElements
					result.OutputElements = validationResult.OutputElements
					result.InputAttributes = validationResult.InputAttributes
					result.OutputAttributes = validationResult.OutputAttributes
				}

				// Print brief result
				progressMutex.Lock()
				if result.Success {
					if result.Verdict == "PERFECT" {
						fmt.Printf("  ✅ PERFECT FIDELITY (%.1f%%)\n", result.FidelityPercent)
					} else if result.Verdict == "SEMANTIC" {
						fmt.Printf("  ✅ SEMANTIC FIDELITY (%.1f%%)\n", result.FidelityPercent)
					} else {
						fmt.Printf("  ❌ FAILED (%.1f%%)\n", result.FidelityPercent)
					}
				} else {
					fmt.Printf("  ❌ ERROR: %s\n", result.ErrorMessage)
				}
				fmt.Println()
				progressMutex.Unlock()

				results <- result
			}
		}(w)
	}

	// Send jobs to workers
	for i, xmlFile := range xmlFiles {
		jobs <- job{index: i, xmlFile: xmlFile}
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()
	close(results)

	// Collect results and build summary
	summary := BatchSummary{
		TotalFiles: len(xmlFiles),
		Results:    make([]BatchResult, 0, len(xmlFiles)),
	}

	for result := range results {
		summary.Results = append(summary.Results, result)

		if result.Success {
			summary.SuccessfulFiles++
			if result.Verdict == "PERFECT" {
				summary.PerfectFidelity++
			} else if result.Verdict == "SEMANTIC" {
				summary.SemanticFidelity++
			}
		} else {
			summary.FailedFiles++
		}
	}

	// Print aggregate summary
	printBatchSummary(summary)

	// Write report file if requested
	if reportPath != "" {
		if err := writeBatchReport(reportPath, summary); err != nil {
			return fmt.Errorf("failed to write report: %w", err)
		}
		fmt.Printf("📄 Report written to: %s\n", reportPath)
	}

	return nil
}

// runQuietValidation runs validation with minimal output
func runQuietValidation(inputPath, outputPath string, writeOutput bool) (*ValidationResult, float64, string, error) {
	// Read input XML
	xmlData, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, 0, "", fmt.Errorf("read failed: %w", err)
	}

	inputMetrics := CountXMLItemsDetailed(xmlData)

	// Unmarshal
	var data generated.AssetReportCollectionElement
	if err := xml.Unmarshal(xmlData, &data); err != nil {
		return nil, 0, "", fmt.Errorf("unmarshal failed: %w", err)
	}

	// Extract and set prefixes
	prefixes := generated.ExtractElementPrefixes(xmlData)
	data.SetElementPrefixes(prefixes)

	// Extract which elements had xmlns for exact xmlns replication (zero delta)
	elementsWithXmlns := generated.ExtractElementsWithXmlns(xmlData)
	data.SetElementsWithXmlns(elementsWithXmlns)

	// Validate struct
	validationResult, err := ValidateStruct(&data)
	if err != nil {
		// Continue even if validation has issues
		validationResult = &ValidationResult{}
	}

	// Marshal
	output, err := data.MarshalIndentClean("", "  ")
	if err != nil {
		return nil, 0, "", fmt.Errorf("marshal failed: %w", err)
	}

	// Smart line ending normalization
	output = normalizeLineEndings(output)
	fullOutput := append([]byte(xml.Header), output...)

	// Write if requested
	if writeOutput {
		if err := os.WriteFile(outputPath, fullOutput, 0644); err != nil {
			return nil, 0, "", fmt.Errorf("write failed: %w", err)
		}
	}

	// Count output
	outputMetrics := CountXMLItemsDetailed(fullOutput)

	// Calculate fidelity
	validationResult.InputElements = inputMetrics.Elements
	validationResult.OutputElements = outputMetrics.Elements
	validationResult.InputAttributes = inputMetrics.Attributes
	validationResult.OutputAttributes = outputMetrics.Attributes
	validationResult.InputXmlns = inputMetrics.XmlnsAttrs
	validationResult.OutputXmlns = outputMetrics.XmlnsAttrs
	validationResult.InputNonXmlns = inputMetrics.NonXmlnsAttrs
	validationResult.OutputNonXmlns = outputMetrics.NonXmlnsAttrs
	fidelity := CalculateFidelity(inputMetrics.Elements, outputMetrics.Elements, inputMetrics.Attributes, outputMetrics.Attributes)

	// Determine verdict via deep comparison
	verdict := "FAILURE"
	deepResult, err := DeepCompareXML(xmlData, fullOutput)
	if err == nil {
		if deepResult.IsPerfectMatch {
			verdict = "PERFECT"
		} else {
			// Check semantic fidelity
			hasSemanticFidelity := deepResult.ElementNames.MismatchCount == 0 &&
				deepResult.AttributeNames.MismatchCount == 0 &&
				deepResult.NamespaceURIs.MismatchCount == 0 &&
				deepResult.ElementValues.MismatchCount == 0 &&
				deepResult.AttributeValues.MismatchCount == 0
			if hasSemanticFidelity {
				verdict = "SEMANTIC"
			}
		}
	}

	return validationResult, fidelity, verdict, nil
}

// printBatchSummary prints aggregate summary
func printBatchSummary(summary BatchSummary) {
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("📊 BATCH VALIDATION SUMMARY")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("Total Files Processed:    %d\n", summary.TotalFiles)
	fmt.Printf("Successful Validations:   %d (%.1f%%)\n",
		summary.SuccessfulFiles,
		float64(summary.SuccessfulFiles)/float64(summary.TotalFiles)*100)
	fmt.Printf("Failed Validations:       %d (%.1f%%)\n",
		summary.FailedFiles,
		float64(summary.FailedFiles)/float64(summary.TotalFiles)*100)
	fmt.Println()
	fmt.Printf("Perfect Fidelity:         %d\n", summary.PerfectFidelity)
	fmt.Printf("Semantic Fidelity:        %d\n", summary.SemanticFidelity)
	fmt.Printf("Data Corruption:          %d\n", summary.SuccessfulFiles-summary.PerfectFidelity-summary.SemanticFidelity)
	fmt.Println(strings.Repeat("=", 70))
}

// writeBatchReport writes detailed report to file
func writeBatchReport(reportPath string, summary BatchSummary) error {
	f, err := os.Create(reportPath)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "Batch Validation Report\n")
	fmt.Fprintf(f, "Generated: %s\n\n", filepath.Base(reportPath))
	fmt.Fprintf(f, "Summary:\n")
	fmt.Fprintf(f, "========\n")
	fmt.Fprintf(f, "Total Files:          %d\n", summary.TotalFiles)
	fmt.Fprintf(f, "Successful:           %d (%.1f%%)\n",
		summary.SuccessfulFiles,
		float64(summary.SuccessfulFiles)/float64(summary.TotalFiles)*100)
	fmt.Fprintf(f, "Failed:               %d (%.1f%%)\n",
		summary.FailedFiles,
		float64(summary.FailedFiles)/float64(summary.TotalFiles)*100)
	fmt.Fprintf(f, "Perfect Fidelity:     %d\n", summary.PerfectFidelity)
	fmt.Fprintf(f, "Semantic Fidelity:    %d\n", summary.SemanticFidelity)
	fmt.Fprintf(f, "Data Corruption:      %d\n\n",
		summary.SuccessfulFiles-summary.PerfectFidelity-summary.SemanticFidelity)

	fmt.Fprintf(f, "Detailed Results:\n")
	fmt.Fprintf(f, "=================\n\n")

	for i, result := range summary.Results {
		fmt.Fprintf(f, "%d. %s\n", i+1, result.InputFile)
		if result.Success {
			fmt.Fprintf(f, "   Status:     %s\n", result.Verdict)
			fmt.Fprintf(f, "   Fidelity:   %.1f%%\n", result.FidelityPercent)
			fmt.Fprintf(f, "   Elements:   %d → %d\n", result.InputElements, result.OutputElements)
			fmt.Fprintf(f, "   Attributes: %d → %d\n", result.InputAttributes, result.OutputAttributes)
		} else {
			fmt.Fprintf(f, "   Status:     FAILED\n")
			fmt.Fprintf(f, "   Error:      %s\n", result.ErrorMessage)
		}
		fmt.Fprintf(f, "\n")
	}

	return nil
}

// normalizeLineEndings intelligently normalizes line endings in XML output.
// It removes &#xd; (carriage return entities) only when they're part of CRLF sequences,
// preserving &#xd; that appears in actual element content (e.g., regex patterns).
//
// Strategy:
// - Remove "&#xd;\n" (CRLF in XML) -> "\n" (LF)
// - Remove "&#xd;&#10;" (CRLF as entities) -> "&#10;" (LF)
// - Keep standalone "&#xd;" that's part of content
func normalizeLineEndings(xmlBytes []byte) []byte {
	str := string(xmlBytes)

	// Pattern 1: &#xd; followed by literal newline (most common CRLF case)
	// This is when Windows line ending gets partially encoded
	str = strings.ReplaceAll(str, "&#xd;\n", "\n")

	// Pattern 2: &#xd;&#10; (both CR and LF encoded as entities)
	// This is when both characters get encoded
	str = strings.ReplaceAll(str, "&#xd;&#10;", "&#10;")

	// Pattern 3: &#13;&#10; (decimal encoding variant)
	str = strings.ReplaceAll(str, "&#13;&#10;", "&#10;")

	// Pattern 4: &#13;\n (decimal CR + literal LF)
	str = strings.ReplaceAll(str, "&#13;\n", "\n")

	// Note: We do NOT remove standalone &#xd; or &#13; because those might be
	// actual data content (like in regex patterns: [^\n\r])

	return []byte(str)
}

// runRoundTripTest performs the marshal/unmarshal round-trip test
func runRoundTripTest(inputPath, outputPath string, writeOutput bool) error {
	fmt.Printf("=== Round-Trip Validation Test ===\n")
	fmt.Printf("Schema: 1-1\n")
	fmt.Printf("Input:  %s\n", inputPath)
	if writeOutput {
		fmt.Printf("Output: %s\n\n", outputPath)
	} else {
		fmt.Printf("Output: (in-memory only, no file written)\n\n")
	}

	// Read input XML
	fmt.Println("📖 Reading input XML...")
	xmlData, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}
	fmt.Printf("✅ Read %d bytes\n\n", len(xmlData))

	// Count input XML items with xmlns breakdown
	fmt.Println("📊 Counting input XML items...")
	inputMetrics := CountXMLItemsDetailed(xmlData)
	fmt.Printf("✅ Input has %d elements and %d attributes\n\n", inputMetrics.Elements, inputMetrics.Attributes)

	// Unmarshal into generated struct
	fmt.Println("🔄 Unmarshaling into generated struct...")
	var data generated.AssetReportCollectionElement
	if err := xml.Unmarshal(xmlData, &data); err != nil {
		return fmt.Errorf("unmarshal failed: %w", err)
	}
	fmt.Println("✅ Unmarshal successful\n")

	// Extract element prefixes from raw XML for proper prefix restoration
	fmt.Println("🔍 Extracting element prefixes from input XML...")
	prefixes := generated.ExtractElementPrefixes(xmlData)
	data.SetElementPrefixes(prefixes)
	fmt.Printf("✅ Extracted %d element prefix mappings\n\n", len(prefixes))

	// Extract which elements had xmlns for exact xmlns replication (zero delta)
	fmt.Println("🔍 Tracking which elements had xmlns in input XML...")
	elementsWithXmlns := generated.ExtractElementsWithXmlns(xmlData)
	data.SetElementsWithXmlns(elementsWithXmlns)
	fmt.Printf("✅ Tracked %d element types with xmlns for replication\n\n", len(elementsWithXmlns))

	// Validate using reflection
	fmt.Println("🔍 Validating struct completeness...")
	validationResult, err := ValidateStruct(&data)
	if err != nil {
		fmt.Printf("⚠️  Validation error: %v\n\n", err)
	} else {
		fmt.Printf("✅ Validated %d struct fields\n\n", validationResult.TotalFields)
	}

	// Marshal back to XML with xmlns filtering
	fmt.Println("🔄 Marshaling back to XML...")
	output, err := data.MarshalIndentClean("", "  ")
	if err != nil {
		return fmt.Errorf("marshal failed: %w", err)
	}
	fmt.Printf("✅ Marshal successful (%d bytes)\n\n", len(output))

	// Smart line ending normalization - only remove &#xd; from CRLF sequences
	// Preserve &#xd; that's part of actual content (e.g., in regex patterns)
	output = normalizeLineEndings(output)

	// Prepare full output with XML header
	fullOutput := append([]byte(xml.Header), output...)

	// Write output if requested
	if writeOutput {
		fmt.Println("💾 Writing output file...")
		if err := os.WriteFile(outputPath, fullOutput, 0644); err != nil {
			return fmt.Errorf("failed to write output: %w", err)
		}
		fmt.Printf("✅ Output written: %s\n\n", outputPath)
	} else {
		fmt.Println("💾 Skipping file write (in-memory processing)\n")
	}

	// Count output XML items with xmlns breakdown
	fmt.Println("📊 Counting output XML items...")
	outputMetrics := CountXMLItemsDetailed(fullOutput)
	fmt.Printf("✅ Output has %d elements and %d attributes\n\n", outputMetrics.Elements, outputMetrics.Attributes)

	// Calculate fidelity with detailed xmlns breakdown
	validationResult.InputElements = inputMetrics.Elements
	validationResult.OutputElements = outputMetrics.Elements
	validationResult.InputAttributes = inputMetrics.Attributes
	validationResult.OutputAttributes = outputMetrics.Attributes
	validationResult.InputXmlns = inputMetrics.XmlnsAttrs
	validationResult.OutputXmlns = outputMetrics.XmlnsAttrs
	validationResult.InputNonXmlns = inputMetrics.NonXmlnsAttrs
	validationResult.OutputNonXmlns = outputMetrics.NonXmlnsAttrs
	validationResult.FidelityPercent = CalculateFidelity(inputMetrics.Elements, outputMetrics.Elements, inputMetrics.Attributes, outputMetrics.Attributes)

	// Perform comprehensive deep comparison
	fmt.Println("🔍 Performing comprehensive deep comparison...")
	deepResult, err := DeepCompareXML(xmlData, fullOutput)
	if err != nil {
		fmt.Printf("⚠️  Deep comparison error: %v\n", err)
		// Continue with basic validation even if deep comparison fails
	} else {
		fmt.Println("✅ Deep comparison completed\n")
		// Display comprehensive deep comparison report
		PrintDeepComparisonReport(deepResult)
	}

	// Display detailed validation report (basic metrics)
	fmt.Println("\n" + strings.Repeat("=", 60))
	PrintValidationReport(validationResult)
	fmt.Println(strings.Repeat("=", 60) + "\n")

	// Summary
	fmt.Println("📊 Summary:")
	fmt.Printf("   Input size:  %d bytes\n", len(xmlData))
	fmt.Printf("   Output size: %d bytes\n", len(fullOutput))
	ratio := float64(len(fullOutput)) / float64(len(xmlData)) * 100
	fmt.Printf("   Size ratio:  %.1f%%\n\n", ratio)

	fmt.Println("✅ Round-trip test COMPLETED")

	return nil
}
