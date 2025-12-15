// Package main provides a validator CLI for testing SCAP XML parsing
package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	oval "github.com/aequo-labs/forgexml-scap/internal/generated/org/mitre/oval/xmlschema/oval-definitions-5"
	xccdf "github.com/aequo-labs/forgexml-scap/internal/generated/gov/nist/checklists/xccdf/1-2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: validator <xml-file>")
		fmt.Println("Validates OVAL and XCCDF XML files against generated Go structs")
		os.Exit(1)
	}

	xmlFile := os.Args[1]
	
	// Read the XML file
	data, err := os.ReadFile(xmlFile)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", xmlFile, err)
		os.Exit(1)
	}

	filename := filepath.Base(xmlFile)
	
	// Detect type and parse accordingly
	if strings.Contains(strings.ToLower(filename), "oval") {
		validateOVAL(data, filename)
	} else if strings.Contains(strings.ToLower(filename), "xccdf") {
		validateXCCDF(data, filename)
	} else {
		// Try both
		fmt.Println("Auto-detecting file type...")
		if tryOVAL(data) {
			validateOVAL(data, filename)
		} else if tryXCCDF(data) {
			validateXCCDF(data, filename)
		} else {
			fmt.Printf("Could not determine file type for: %s\n", filename)
			os.Exit(1)
		}
	}
}

func tryOVAL(data []byte) bool {
	var ovalDef oval.Oval_definitionsElement
	return xml.Unmarshal(data, &ovalDef) == nil && ovalDef.XMLName.Local != ""
}

func tryXCCDF(data []byte) bool {
	var benchmark xccdf.BenchmarkElement
	return xml.Unmarshal(data, &benchmark) == nil && benchmark.XMLName.Local != ""
}

func validateOVAL(data []byte, filename string) {
	var ovalDef oval.Oval_definitionsElement
	err := xml.Unmarshal(data, &ovalDef)
	if err != nil {
		fmt.Printf("Error parsing OVAL XML: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully parsed OVAL: %s\n", filename)
	fmt.Println("---")
	
	// Generator info
	fmt.Println("Generator:")
	if ovalDef.Generator.Product_name != nil {
		fmt.Printf("  Product: %s\n", *ovalDef.Generator.Product_name)
	}
	if ovalDef.Generator.Schema_version != "" {
		fmt.Printf("  Schema Version: %s\n", ovalDef.Generator.Schema_version)
	}
	fmt.Printf("  Timestamp: %v\n", ovalDef.Generator.Timestamp)

	// Definitions count
	if ovalDef.Definitions != nil && len(ovalDef.Definitions.Definition) > 0 {
		fmt.Printf("\nDefinitions: %d\n", len(ovalDef.Definitions.Definition))
		for i, def := range ovalDef.Definitions.Definition {
			if i >= 3 {
				fmt.Printf("  ... and %d more\n", len(ovalDef.Definitions.Definition)-3)
				break
			}
			fmt.Printf("  - %s\n", def.Metadata.Title)
		}
	}

	if ovalDef.Tests != nil {
		fmt.Printf("\nTests: %d\n", len(ovalDef.Tests.Test))
	}
	if ovalDef.Objects != nil {
		fmt.Printf("Objects: %d\n", len(ovalDef.Objects.Object))
	}
	if ovalDef.States != nil {
		fmt.Printf("States: %d\n", len(ovalDef.States.State))
	}
	if ovalDef.Variables != nil {
		fmt.Printf("Variables: present\n")
	}

	fmt.Println("\n✅ OVAL validation successful!")
}

func validateXCCDF(data []byte, filename string) {
	var benchmark xccdf.BenchmarkElement
	err := xml.Unmarshal(data, &benchmark)
	if err != nil {
		fmt.Printf("Error parsing XCCDF XML: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully parsed XCCDF: %s\n", filename)
	fmt.Println("---")
	
	// Title
	if len(benchmark.Title) > 0 {
		fmt.Printf("Title: %s\n", benchmark.Title[0].Value)
	}

	// Status
	if len(benchmark.Status) > 0 {
		fmt.Printf("Status: %s (date: %s)\n", benchmark.Status[0].Value, benchmark.Status[0].Date)
	}

	// Version
	if benchmark.Version.Value != "" {
		fmt.Printf("Version: %s\n", benchmark.Version.Value)
	}

	// Profiles
	if len(benchmark.Profile) > 0 {
		fmt.Printf("\nProfiles: %d\n", len(benchmark.Profile))
		for i, profile := range benchmark.Profile {
			if i >= 3 {
				fmt.Printf("  ... and %d more\n", len(benchmark.Profile)-3)
				break
			}
			if len(profile.Title) > 0 {
				fmt.Printf("  - %s\n", profile.Title[0].Value)
			}
		}
	}

	// Groups
	if len(benchmark.Group) > 0 {
		fmt.Printf("\nGroups: %d\n", len(benchmark.Group))
	}

	// Rules
	if len(benchmark.Rule) > 0 {
		fmt.Printf("Rules: %d\n", len(benchmark.Rule))
	}

	// Values
	if len(benchmark.Value) > 0 {
		fmt.Printf("Values: %d\n", len(benchmark.Value))
	}

	fmt.Println("\n✅ XCCDF validation successful!")
}
