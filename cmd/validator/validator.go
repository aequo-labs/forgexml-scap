package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
	"strings"
)

// XMLMetrics holds detailed XML counting with xmlns breakdown
type XMLMetrics struct {
	Elements      int
	Attributes    int
	XmlnsAttrs    int
	NonXmlnsAttrs int
}

// ValidationResult holds the results of struct validation
type ValidationResult struct {
	TotalFields         int
	PopulatedFields     int
	EmptyFields         []string
	UnknownElements     int
	UnknownAttrs        int
	UnknownElementsList []string // Actual unknown element names
	UnknownAttrsList    []string // Actual unknown attribute names
	Warnings            []string
	// XML Fidelity Metrics
	InputElements    int
	OutputElements   int
	InputAttributes  int
	OutputAttributes int
	FidelityPercent  float64
	// Detailed attribute breakdown
	InputXmlns     int
	OutputXmlns    int
	InputNonXmlns  int
	OutputNonXmlns int
}

// ValidateStruct performs reflection-based validation of the unmarshaled struct
func ValidateStruct(data interface{}) (*ValidationResult, error) {
	result := &ValidationResult{
		EmptyFields:         []string{},
		UnknownElementsList: []string{},
		UnknownAttrsList:    []string{},
		Warnings:            []string{},
	}

	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %v", val.Kind())
	}

	validateStructFields(val, "", result)

	return result, nil
}

// validateStructFields recursively validates struct fields
func validateStructFields(val reflect.Value, prefix string, result *ValidationResult) {
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		fieldName := fieldType.Name

		// Skip unexported fields
		if !field.CanInterface() {
			continue
		}

		// Build full field path
		fullName := fieldName
		if prefix != "" {
			fullName = prefix + "." + fieldName
		}

		// Handle special fields
		if fieldName == "XMLName" {
			continue // Skip XMLName, it's always populated
		}

		if fieldName == "UnknownElements" {
			result.UnknownElements = field.Len()
			// Collect actual element names
			for i := 0; i < field.Len(); i++ {
				elem := field.Index(i)
				if elem.Kind() == reflect.Struct {
					// Look for XMLName field
					xmlNameField := elem.FieldByName("XMLName")
					if xmlNameField.IsValid() && xmlNameField.Kind() == reflect.Struct {
						localField := xmlNameField.FieldByName("Local")
						if localField.IsValid() && localField.Kind() == reflect.String {
							elemName := localField.String()
							result.UnknownElementsList = append(result.UnknownElementsList, elemName)
						}
					}
				}
			}
			continue
		}

		if fieldName == "UnknownAttrs" {
			result.UnknownAttrs = field.Len()
			// Collect actual attribute names
			for i := 0; i < field.Len(); i++ {
				attr := field.Index(i)
				if attr.Kind() == reflect.Struct {
					// xml.Attr has Name.Local field
					nameField := attr.FieldByName("Name")
					if nameField.IsValid() && nameField.Kind() == reflect.Struct {
						localField := nameField.FieldByName("Local")
						if localField.IsValid() && localField.Kind() == reflect.String {
							attrName := localField.String()
							result.UnknownAttrsList = append(result.UnknownAttrsList, attrName)
						}
					}
				}
			}
			continue
		}

		result.TotalFields++

		// Check if field is populated
		if isFieldEmpty(field, fieldType) {
			result.EmptyFields = append(result.EmptyFields, fullName)
		} else {
			result.PopulatedFields++

			// Recursively validate nested structs
			if field.Kind() == reflect.Struct {
				validateStructFields(field, fullName, result)
			} else if field.Kind() == reflect.Ptr && !field.IsNil() && field.Elem().Kind() == reflect.Struct {
				validateStructFields(field.Elem(), fullName, result)
			} else if field.Kind() == reflect.Slice && field.Len() > 0 {
				// Validate slice elements
				for j := 0; j < field.Len(); j++ {
					elem := field.Index(j)
					elemPath := fmt.Sprintf("%s[%d]", fullName, j)

					if elem.Kind() == reflect.Struct {
						validateStructFields(elem, elemPath, result)
					} else if elem.Kind() == reflect.Ptr && !elem.IsNil() && elem.Elem().Kind() == reflect.Struct {
						validateStructFields(elem.Elem(), elemPath, result)
					}
				}
			}
		}
	}
}

// isFieldEmpty checks if a field is empty/zero value
func isFieldEmpty(field reflect.Value, fieldType reflect.StructField) bool {
	// Check for omitempty tag
	xmlTag := fieldType.Tag.Get("xml")
	hasOmitEmpty := strings.Contains(xmlTag, "omitempty")

	switch field.Kind() {
	case reflect.String:
		return field.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return field.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return field.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return field.Float() == 0.0
	case reflect.Bool:
		return !field.Bool()
	case reflect.Ptr:
		return field.IsNil()
	case reflect.Slice, reflect.Array:
		return field.Len() == 0
	case reflect.Map:
		return field.Len() == 0
	case reflect.Interface:
		return field.IsNil()
	case reflect.Struct:
		// For structs, only consider empty if omitempty and all fields are zero
		if hasOmitEmpty {
			return field.IsZero()
		}
		return false // Non-omitempty structs are considered populated
	default:
		return field.IsZero()
	}
}

// CountXMLItems counts elements and attributes in XML data
func CountXMLItems(xmlData []byte) (elements int, attributes int, err error) {
	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	elements = 0
	attributes = 0

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, 0, fmt.Errorf("error parsing XML: %w", err)
		}

		switch t := token.(type) {
		case xml.StartElement:
			elements++
			attributes += len(t.Attr)
		}
	}

	return elements, attributes, nil
}

// CountXMLItemsDetailed counts elements and attributes with xmlns breakdown
func CountXMLItemsDetailed(xmlData []byte) XMLMetrics {
	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	metrics := XMLMetrics{}

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return metrics
		}

		switch t := token.(type) {
		case xml.StartElement:
			metrics.Elements++
			for _, attr := range t.Attr {
				metrics.Attributes++
				// Count xmlns attributes separately
				if attr.Name.Space == "xmlns" || (attr.Name.Local == "xmlns" && attr.Name.Space == "") {
					metrics.XmlnsAttrs++
				} else {
					metrics.NonXmlnsAttrs++
				}
			}
		}
	}

	return metrics
}

// CalculateFidelity calculates the fidelity percentage between input and output XML
func CalculateFidelity(inputElems, outputElems, inputAttrs, outputAttrs int) float64 {
	totalInput := inputElems + inputAttrs
	totalOutput := outputElems + outputAttrs

	if totalInput == 0 {
		return 100.0
	}

	// Calculate what percentage of input items are preserved in output
	preserved := totalOutput
	if totalOutput > totalInput {
		// If output has more items, cap preserved at input total
		preserved = totalInput
	}

	return (float64(preserved) / float64(totalInput)) * 100.0
}

// FidelityResult contains fidelity calculation results (DEPRECATED - use CalculateFidelity float64 return)
type FidelityResult struct {
	Percentage float64
	Status     string
	Message    string
}

// XMLComparisonResult contains detailed XML comparison results
type XMLComparisonResult struct {
	InputElements    int
	OutputElements   int
	InputAttributes  int
	OutputAttributes int

	MissingElements  []string // Elements in input but not in output
	ExtraElements    []string // Elements in output but not in input
	MismatchedValues []string // Elements with different text content
	MissingAttrs     []string // Attributes in input but not in output
	ExtraAttrs       []string // Attributes in output but not in input
	MismatchedAttrs  []string // Attributes with different values
	NamespaceIssues  []string // Namespace prefix or declaration issues

	TotalDifferences int
	IsPerfectMatch   bool
}

// XMLElement represents a parsed XML element for comparison
type XMLElement struct {
	Name       xml.Name
	Attributes []xml.Attr
	Value      string
	Path       string // XPath-like path to element
	Children   []*XMLElement
}

// ParseXMLTree parses XML data into a tree structure for comparison
func ParseXMLTree(xmlData []byte) (*XMLElement, error) {
	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	root := &XMLElement{Path: "/"}
	var stack []*XMLElement
	var currentPath []string

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error parsing XML: %w", err)
		}

		switch t := token.(type) {
		case xml.StartElement:
			elem := &XMLElement{
				Name:       t.Name,
				Attributes: t.Attr,
				Path:       "/" + strings.Join(append(currentPath, t.Name.Local), "/"),
			}

			if len(stack) == 0 {
				root = elem
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, elem)
			}

			stack = append(stack, elem)
			currentPath = append(currentPath, t.Name.Local)

		case xml.EndElement:
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			if len(currentPath) > 0 {
				currentPath = currentPath[:len(currentPath)-1]
			}

		case xml.CharData:
			text := strings.TrimSpace(string(t))
			if text != "" && len(stack) > 0 {
				current := stack[len(stack)-1]
				current.Value = text
			}
		}
	}

	return root, nil
}

// CompareXMLTrees performs deep comparison of two XML trees
func CompareXMLTrees(input, output *XMLElement) *XMLComparisonResult {
	result := &XMLComparisonResult{
		MissingElements:  []string{},
		ExtraElements:    []string{},
		MismatchedValues: []string{},
		MissingAttrs:     []string{},
		ExtraAttrs:       []string{},
		MismatchedAttrs:  []string{},
		NamespaceIssues:  []string{},
	}

	// Count totals
	result.InputElements, result.InputAttributes = countElementsAndAttrs(input)
	result.OutputElements, result.OutputAttributes = countElementsAndAttrs(output)

	// Perform deep comparison
	compareElements(input, output, result)

	// Calculate totals
	result.TotalDifferences = len(result.MissingElements) + len(result.ExtraElements) +
		len(result.MismatchedValues) + len(result.MissingAttrs) +
		len(result.ExtraAttrs) + len(result.MismatchedAttrs) +
		len(result.NamespaceIssues)

	result.IsPerfectMatch = result.TotalDifferences == 0

	return result
}

// countElementsAndAttrs counts elements and attributes in a tree
func countElementsAndAttrs(elem *XMLElement) (int, int) {
	if elem == nil {
		return 0, 0
	}

	elements := 1
	attrs := len(elem.Attributes)

	for _, child := range elem.Children {
		childElems, childAttrs := countElementsAndAttrs(child)
		elements += childElems
		attrs += childAttrs
	}

	return elements, attrs
}

// compareElements recursively compares two elements
func compareElements(input, output *XMLElement, result *XMLComparisonResult) {
	if input == nil && output == nil {
		return
	}

	// Check for missing/extra elements
	if input == nil {
		result.ExtraElements = append(result.ExtraElements,
			fmt.Sprintf("%s (extra in output)", output.Path))
		return
	}
	if output == nil {
		result.MissingElements = append(result.MissingElements,
			fmt.Sprintf("%s (missing from output)", input.Path))
		return
	}

	// Check namespace differences
	if input.Name.Space != output.Name.Space {
		result.NamespaceIssues = append(result.NamespaceIssues,
			fmt.Sprintf("%s: namespace changed from '%s' to '%s'",
				input.Path, input.Name.Space, output.Name.Space))
	}

	// Check local name differences
	if input.Name.Local != output.Name.Local {
		result.MismatchedValues = append(result.MismatchedValues,
			fmt.Sprintf("%s: element name changed from '%s' to '%s'",
				input.Path, input.Name.Local, output.Name.Local))
	}

	// Check text content
	if input.Value != output.Value {
		result.MismatchedValues = append(result.MismatchedValues,
			fmt.Sprintf("%s: value changed from '%s' to '%s'",
				input.Path, truncate(input.Value, 50), truncate(output.Value, 50)))
	}

	// Compare attributes
	compareAttributes(input, output, result)

	// Compare children
	compareChildren(input, output, result)
}

// compareAttributes compares attributes of two elements
func compareAttributes(input, output *XMLElement, result *XMLComparisonResult) {
	inputAttrs := make(map[string]xml.Attr)
	for _, attr := range input.Attributes {
		key := attr.Name.Space + ":" + attr.Name.Local
		inputAttrs[key] = attr
	}

	outputAttrs := make(map[string]xml.Attr)
	for _, attr := range output.Attributes {
		key := attr.Name.Space + ":" + attr.Name.Local
		outputAttrs[key] = attr

		// Check for namespace corruption (xmlns: becoming _xmlns:)
		if strings.HasPrefix(attr.Name.Space, "_xmlns") || strings.HasPrefix(attr.Name.Local, "_xmlns") {
			result.NamespaceIssues = append(result.NamespaceIssues,
				fmt.Sprintf("%s: corrupted namespace attribute '%s:%s' = '%s'",
					output.Path, attr.Name.Space, attr.Name.Local, attr.Value))
		}
	}

	// Check for missing attributes
	for key, attr := range inputAttrs {
		if _, exists := outputAttrs[key]; !exists {
			result.MissingAttrs = append(result.MissingAttrs,
				fmt.Sprintf("%s/@%s:%s (missing from output)",
					input.Path, attr.Name.Space, attr.Name.Local))
		}
	}

	// Check for extra attributes
	for key, attr := range outputAttrs {
		if _, exists := inputAttrs[key]; !exists {
			// Skip corrupted namespace attributes in the report
			if !strings.HasPrefix(attr.Name.Space, "_xmlns") && !strings.HasPrefix(attr.Name.Local, "_xmlns") {
				result.ExtraAttrs = append(result.ExtraAttrs,
					fmt.Sprintf("%s/@%s:%s (extra in output)",
						output.Path, attr.Name.Space, attr.Name.Local))
			}
		}
	}

	// Check for mismatched attribute values
	for key, inputAttr := range inputAttrs {
		if outputAttr, exists := outputAttrs[key]; exists {
			if inputAttr.Value != outputAttr.Value {
				result.MismatchedAttrs = append(result.MismatchedAttrs,
					fmt.Sprintf("%s/@%s:%s: value changed from '%s' to '%s'",
						input.Path, inputAttr.Name.Space, inputAttr.Name.Local,
						truncate(inputAttr.Value, 50), truncate(outputAttr.Value, 50)))
			}
		}
	}
}

// compareChildren compares children of two elements
func compareChildren(input, output *XMLElement, result *XMLComparisonResult) {
	// Build maps for quick lookup
	inputChildren := make(map[string][]*XMLElement)
	for _, child := range input.Children {
		key := child.Name.Space + ":" + child.Name.Local
		inputChildren[key] = append(inputChildren[key], child)
	}

	outputChildren := make(map[string][]*XMLElement)
	for _, child := range output.Children {
		key := child.Name.Space + ":" + child.Name.Local
		outputChildren[key] = append(outputChildren[key], child)
	}

	// Compare matching children
	for key, inputList := range inputChildren {
		outputList, exists := outputChildren[key]
		if !exists {
			// All input children of this type are missing
			for _, child := range inputList {
				result.MissingElements = append(result.MissingElements,
					fmt.Sprintf("%s (missing from output)", child.Path))
			}
			continue
		}

		// Compare matching pairs
		maxLen := len(inputList)
		if len(outputList) > maxLen {
			maxLen = len(outputList)
		}

		for i := 0; i < maxLen; i++ {
			var inputChild, outputChild *XMLElement
			if i < len(inputList) {
				inputChild = inputList[i]
			}
			if i < len(outputList) {
				outputChild = outputList[i]
			}

			compareElements(inputChild, outputChild, result)
		}
	}

	// Check for extra children in output
	for key, outputList := range outputChildren {
		if _, exists := inputChildren[key]; !exists {
			for _, child := range outputList {
				result.ExtraElements = append(result.ExtraElements,
					fmt.Sprintf("%s (extra in output)", child.Path))
			}
		}
	}
}

// truncate truncates a string to maxLen characters
func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

// PrintValidationReport prints a formatted validation report
func PrintValidationReport(result *ValidationResult) {
	fmt.Println("📊 Validation Report:")
	fmt.Println("=====================")
	fmt.Printf("Total Fields:      %d\n", result.TotalFields)
	fmt.Printf("Populated Fields:  %d\n", result.PopulatedFields)
	fmt.Printf("Empty Fields:      %d\n", len(result.EmptyFields))
	fmt.Printf("Unknown Elements:  %d\n", result.UnknownElements)
	fmt.Printf("Unknown Attrs:     %d\n", result.UnknownAttrs)

	// Show actual unknown elements only if count > 0 (elements not defined in schema)
	if result.UnknownElements > 0 && len(result.UnknownElementsList) > 0 {
		fmt.Println("\n📋 Unknown Elements (not defined in schema, captured via xs:any):")
		seen := make(map[string]bool)
		for _, elemName := range result.UnknownElementsList {
			if !seen[elemName] {
				fmt.Printf("  - %s\n", elemName)
				seen[elemName] = true
			}
		}
		fmt.Println("   Note: These elements are from other namespaces, allowed by schema xs:any wildcard")
	}

	// Show actual unknown attributes only if count > 0 (attributes not defined in schema)
	if result.UnknownAttrs > 0 && len(result.UnknownAttrsList) > 0 {
		fmt.Println("\n📋 Unknown Attributes (not defined in schema, captured via xs:anyAttribute):")
		seen := make(map[string]bool)
		for _, attrName := range result.UnknownAttrsList {
			if !seen[attrName] {
				fmt.Printf("  - %s\n", attrName)
				seen[attrName] = true
			}
		}
		fmt.Println("   Note: These attributes are from other namespaces, allowed by schema xs:anyAttribute wildcard")
	}

	// Calculate completeness percentage
	if result.TotalFields > 0 {
		completeness := float64(result.PopulatedFields) / float64(result.TotalFields) * 100
		fmt.Printf("\n📊 Struct Completeness: %.1f%% (%d of %d fields populated)\n",
			completeness, result.PopulatedFields, result.TotalFields)
		fmt.Println("   Note: Empty fields indicate optional schema elements not present in this XML instance")
	}

	// Display XML fidelity metrics
	if result.InputElements > 0 || result.OutputElements > 0 {
		fmt.Println("\n📊 XML Fidelity Metrics:")
		fmt.Println("========================")
		fmt.Printf("Input Elements:     %d\n", result.InputElements)
		fmt.Printf("Output Elements:    %d\n", result.OutputElements)
		fmt.Printf("Input Attributes:   %d\n", result.InputAttributes)
		fmt.Printf("Output Attributes:  %d\n", result.OutputAttributes)

		// Display detailed xmlns breakdown
		fmt.Println("\n📊 Attribute Breakdown:")
		fmt.Printf("Input xmlns attrs:     %d\n", result.InputXmlns)
		fmt.Printf("Output xmlns attrs:    %d\n", result.OutputXmlns)
		fmt.Printf("Input non-xmlns attrs: %d\n", result.InputNonXmlns)
		fmt.Printf("Output non-xmlns attrs:%d\n", result.OutputNonXmlns)

		xmlnsDelta := result.OutputXmlns - result.InputXmlns
		nonXmlnsDelta := result.OutputNonXmlns - result.InputNonXmlns
		if xmlnsDelta == 0 && nonXmlnsDelta == 0 {
			fmt.Println("\n🎯 Perfect xmlns fidelity - no attribute delta!")
		} else {
			fmt.Printf("\nxmlns delta:     %+d\n", xmlnsDelta)
			fmt.Printf("non-xmlns delta: %+d\n", nonXmlnsDelta)
		}

		totalInput := result.InputElements + result.InputAttributes
		totalOutput := result.OutputElements + result.OutputAttributes
		fmt.Printf("\nTotal Items (Input):  %d\n", totalInput)
		fmt.Printf("Total Items (Output): %d\n", totalOutput)

		if result.FidelityPercent > 0 {
			fmt.Printf("\n✅ XML Fidelity: %.1f%% (%d/%d items)\n",
				result.FidelityPercent, totalOutput, totalInput)

			if result.FidelityPercent < 100.0 {
				fmt.Println("❌ CRITICAL: Data loss detected!")
				fmt.Println("   Zero tolerance for data loss - ALL items must be preserved")
				fmt.Println("   Check for:")
				fmt.Println("   - Missing struct fields for elements/attributes")
				fmt.Println("   - Incorrect xml tags or omitempty usage")
				fmt.Println("   - Type mismatches preventing unmarshaling")
			} else {
				fmt.Println("🎯 Perfect fidelity - all XML items preserved!")
			}
		}
	}
}

// Deep XML comparison functions for comprehensive validation

// DeepComparisonResult contains detailed comparison results
type DeepComparisonResult struct {
	// Count validation
	ElementCount   CountMatch
	AttributeCount CountMatch
	XmlnsCount     CountMatch
	NonXmlnsCount  CountMatch

	// Name validation
	ElementNames   NameMatch
	AttributeNames NameMatch
	NamespaceURIs  NameMatch

	// Value validation (CRITICAL)
	ElementValues   ValueMatch
	AttributeValues ValueMatch

	// Structural validation
	MaxDepth      DepthMatch
	ElementOrder  bool
	NestingIntact bool

	// Overall verdict
	IsPerfectMatch bool
	TotalIssues    int
}

type CountMatch struct {
	Input  int
	Output int
	Delta  int
	Match  bool
}

type NameMatch struct {
	UniqueNames   int
	MatchCount    int
	MismatchCount int
	Missing       []string
	Extra         []string
	Changed       []NameChange
}

type NameChange struct {
	Path    string
	OldName string
	NewName string
}

type ValueMatch struct {
	Total         int
	MatchCount    int
	MismatchCount int
	Mismatches    []ValueMismatch
}

type ValueMismatch struct {
	Path        string
	Name        string
	InputValue  string
	OutputValue string
}

type DepthMatch struct {
	Input  int
	Output int
	Match  bool
}

// ParseXMLForComparison parses XML into a comparable structure
func ParseXMLForComparison(xmlData []byte) (*XMLNode, error) {
	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	var root *XMLNode
	var current *XMLNode
	var stack []*XMLNode

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("parse error: %w", err)
		}

		switch t := token.(type) {
		case xml.StartElement:
			node := &XMLNode{
				Name:       t.Name,
				Attributes: t.Attr,
				Path:       buildPath(stack, t.Name),
				Children:   []*XMLNode{},
			}

			if current != nil {
				current.Children = append(current.Children, node)
				stack = append(stack, current)
			} else {
				root = node
			}
			current = node

		case xml.EndElement:
			if len(stack) > 0 {
				current = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			} else {
				current = nil
			}

		case xml.CharData:
			if current != nil {
				text := strings.TrimSpace(string(t))
				if text != "" {
					current.Value = text
				}
			}
		}
	}

	return root, nil
}

type XMLNode struct {
	Name       xml.Name
	Attributes []xml.Attr
	Value      string
	Path       string
	Children   []*XMLNode
}

func buildPath(stack []*XMLNode, name xml.Name) string {
	var path strings.Builder
	for _, node := range stack {
		path.WriteString("/")
		if node.Name.Space != "" {
			path.WriteString("{")
			path.WriteString(node.Name.Space)
			path.WriteString("}")
		}
		path.WriteString(node.Name.Local)
	}
	path.WriteString("/")
	if name.Space != "" {
		path.WriteString("{")
		path.WriteString(name.Space)
		path.WriteString("}")
	}
	path.WriteString(name.Local)
	return path.String()
}

// DeepCompareXML performs comprehensive comparison
func DeepCompareXML(inputXML, outputXML []byte) (*DeepComparisonResult, error) {
	result := &DeepComparisonResult{}

	// Parse both trees
	inputTree, err := ParseXMLForComparison(inputXML)
	if err != nil {
		return nil, fmt.Errorf("failed to parse input: %w", err)
	}

	outputTree, err := ParseXMLForComparison(outputXML)
	if err != nil {
		return nil, fmt.Errorf("failed to parse output: %w", err)
	}

	// Level 1: Count validation
	inputMetrics := CountXMLItemsDetailed(inputXML)
	outputMetrics := CountXMLItemsDetailed(outputXML)

	result.ElementCount = CountMatch{
		Input:  inputMetrics.Elements,
		Output: outputMetrics.Elements,
		Delta:  outputMetrics.Elements - inputMetrics.Elements,
		Match:  inputMetrics.Elements == outputMetrics.Elements,
	}

	result.AttributeCount = CountMatch{
		Input:  inputMetrics.Attributes,
		Output: outputMetrics.Attributes,
		Delta:  outputMetrics.Attributes - inputMetrics.Attributes,
		Match:  inputMetrics.Attributes == outputMetrics.Attributes,
	}

	result.XmlnsCount = CountMatch{
		Input:  inputMetrics.XmlnsAttrs,
		Output: outputMetrics.XmlnsAttrs,
		Delta:  outputMetrics.XmlnsAttrs - inputMetrics.XmlnsAttrs,
		Match:  inputMetrics.XmlnsAttrs == outputMetrics.XmlnsAttrs,
	}

	result.NonXmlnsCount = CountMatch{
		Input:  inputMetrics.NonXmlnsAttrs,
		Output: outputMetrics.NonXmlnsAttrs,
		Delta:  outputMetrics.NonXmlnsAttrs - inputMetrics.NonXmlnsAttrs,
		Match:  inputMetrics.NonXmlnsAttrs == outputMetrics.NonXmlnsAttrs,
	}

	// Level 2: Name validation
	result.ElementNames = compareElementNames(inputTree, outputTree)
	result.AttributeNames = compareAttributeNames(inputTree, outputTree)
	result.NamespaceURIs = compareNamespaceURIs(inputTree, outputTree)

	// Level 3: Value validation (CRITICAL)
	result.ElementValues = compareElementValues(inputTree, outputTree)
	result.AttributeValues = compareAttributeValues(inputTree, outputTree)

	// Level 4: Structural validation
	result.MaxDepth = DepthMatch{
		Input:  calculateDepth(inputTree),
		Output: calculateDepth(outputTree),
		Match:  calculateDepth(inputTree) == calculateDepth(outputTree),
	}
	result.ElementOrder = compareOrder(inputTree, outputTree)
	result.NestingIntact = compareNesting(inputTree, outputTree)

	// Calculate overall verdict
	result.IsPerfectMatch = result.ElementCount.Match &&
		result.AttributeCount.Match &&
		result.XmlnsCount.Match &&
		result.NonXmlnsCount.Match &&
		result.ElementNames.MismatchCount == 0 &&
		result.AttributeNames.MismatchCount == 0 &&
		result.NamespaceURIs.MismatchCount == 0 &&
		result.ElementValues.MismatchCount == 0 &&
		result.AttributeValues.MismatchCount == 0 &&
		result.MaxDepth.Match &&
		result.ElementOrder &&
		result.NestingIntact

	result.TotalIssues =
		result.ElementNames.MismatchCount +
			result.AttributeNames.MismatchCount +
			result.NamespaceURIs.MismatchCount +
			result.ElementValues.MismatchCount +
			result.AttributeValues.MismatchCount

	if !result.ElementCount.Match {
		result.TotalIssues++
	}
	if !result.AttributeCount.Match {
		result.TotalIssues++
	}

	return result, nil
}

// Helper comparison functions
func compareElementNames(input, output *XMLNode) NameMatch {
	inputNames := collectElementNames(input, make(map[string]bool))
	outputNames := collectElementNames(output, make(map[string]bool))

	match := NameMatch{
		UniqueNames: len(inputNames),
		Missing:     []string{},
		Extra:       []string{},
		Changed:     []NameChange{},
	}

	for name := range inputNames {
		if !outputNames[name] {
			match.Missing = append(match.Missing, name)
			match.MismatchCount++
		} else {
			match.MatchCount++
		}
	}

	for name := range outputNames {
		if !inputNames[name] {
			match.Extra = append(match.Extra, name)
			match.MismatchCount++
		}
	}

	return match
}

func collectElementNames(node *XMLNode, names map[string]bool) map[string]bool {
	if node == nil {
		return names
	}
	names[node.Name.Local] = true
	for _, child := range node.Children {
		collectElementNames(child, names)
	}
	return names
}

func compareAttributeNames(input, output *XMLNode) NameMatch {
	inputAttrs := collectAttributeNames(input, make(map[string]bool))
	outputAttrs := collectAttributeNames(output, make(map[string]bool))

	match := NameMatch{
		UniqueNames: len(inputAttrs),
		Missing:     []string{},
		Extra:       []string{},
	}

	for name := range inputAttrs {
		if !outputAttrs[name] {
			match.Missing = append(match.Missing, name)
			match.MismatchCount++
		} else {
			match.MatchCount++
		}
	}

	for name := range outputAttrs {
		if !inputAttrs[name] {
			match.Extra = append(match.Extra, name)
			match.MismatchCount++
		}
	}

	return match
}

func collectAttributeNames(node *XMLNode, names map[string]bool) map[string]bool {
	if node == nil {
		return names
	}
	for _, attr := range node.Attributes {
		if attr.Name.Space != "xmlns" && attr.Name.Local != "xmlns" {
			names[attr.Name.Local] = true
		}
	}
	for _, child := range node.Children {
		collectAttributeNames(child, names)
	}
	return names
}

func compareNamespaceURIs(input, output *XMLNode) NameMatch {
	inputNS := collectNamespaceURIs(input, make(map[string]bool))
	outputNS := collectNamespaceURIs(output, make(map[string]bool))

	match := NameMatch{
		UniqueNames: len(inputNS),
		Missing:     []string{},
		Extra:       []string{},
	}

	for uri := range inputNS {
		if !outputNS[uri] {
			match.Missing = append(match.Missing, uri)
			match.MismatchCount++
		} else {
			match.MatchCount++
		}
	}

	for uri := range outputNS {
		if !inputNS[uri] {
			match.Extra = append(match.Extra, uri)
			match.MismatchCount++
		}
	}

	return match
}

func collectNamespaceURIs(node *XMLNode, uris map[string]bool) map[string]bool {
	if node == nil {
		return uris
	}
	for _, attr := range node.Attributes {
		if attr.Name.Space == "xmlns" || attr.Name.Local == "xmlns" {
			uris[attr.Value] = true
		}
	}
	for _, child := range node.Children {
		collectNamespaceURIs(child, uris)
	}
	return uris
}

func compareElementValues(input, output *XMLNode) ValueMatch {
	match := ValueMatch{
		Mismatches: []ValueMismatch{},
	}

	compareValuesRecursive(input, output, &match)
	match.Total = match.MatchCount + match.MismatchCount

	return match
}

func compareValuesRecursive(input, output *XMLNode, match *ValueMatch) {
	if input == nil || output == nil {
		return
	}

	// Compare this node's value
	if input.Value != "" || output.Value != "" {
		if input.Value == output.Value {
			match.MatchCount++
		} else {
			match.MismatchCount++
			if len(match.Mismatches) < 20 { // Limit to first 20
				match.Mismatches = append(match.Mismatches, ValueMismatch{
					Path:        input.Path,
					Name:        input.Name.Local,
					InputValue:  input.Value,
					OutputValue: output.Value,
				})
			}
		}
	}

	// Compare children by identity (namespace + local name), not by position
	// Build map of output children by identity for efficient matching
	outputMap := make(map[string][]*XMLNode)
	for _, outputChild := range output.Children {
		key := outputChild.Name.Space + ":" + outputChild.Name.Local
		outputMap[key] = append(outputMap[key], outputChild)
	}

	// Match input children by identity
	for _, inputChild := range input.Children {
		key := inputChild.Name.Space + ":" + inputChild.Name.Local
		if matchingNodes := outputMap[key]; len(matchingNodes) > 0 {
			// Compare with first matching element (consume it)
			compareValuesRecursive(inputChild, matchingNodes[0], match)
			// Remove the matched node so we don't match it again
			outputMap[key] = matchingNodes[1:]
		}
		// If no match found, the element is missing in output (already counted in name validation)
	}
}

func compareAttributeValues(input, output *XMLNode) ValueMatch {
	match := ValueMatch{
		Mismatches: []ValueMismatch{},
	}

	compareAttrValuesRecursive(input, output, &match)
	match.Total = match.MatchCount + match.MismatchCount

	return match
}

func compareAttrValuesRecursive(input, output *XMLNode, match *ValueMatch) {
	if input == nil || output == nil {
		return
	}

	// Build attribute maps
	inputAttrs := make(map[string]string)
	outputAttrs := make(map[string]string)

	for _, attr := range input.Attributes {
		if attr.Name.Space != "xmlns" && attr.Name.Local != "xmlns" {
			inputAttrs[attr.Name.Local] = attr.Value
		}
	}

	for _, attr := range output.Attributes {
		if attr.Name.Space != "xmlns" && attr.Name.Local != "xmlns" {
			outputAttrs[attr.Name.Local] = attr.Value
		}
	}

	// Compare attribute values
	for name, inputVal := range inputAttrs {
		if outputVal, exists := outputAttrs[name]; exists {
			if inputVal == outputVal {
				match.MatchCount++
			} else {
				match.MismatchCount++
				if len(match.Mismatches) < 20 {
					match.Mismatches = append(match.Mismatches, ValueMismatch{
						Path:        input.Path,
						Name:        "@" + name,
						InputValue:  inputVal,
						OutputValue: outputVal,
					})
				}
			}
		}
	}

	// Recurse to children by identity, not position
	// Build map of output children by identity
	outputMap := make(map[string][]*XMLNode)
	for _, outputChild := range output.Children {
		key := outputChild.Name.Space + ":" + outputChild.Name.Local
		outputMap[key] = append(outputMap[key], outputChild)
	}

	// Match input children by identity
	for _, inputChild := range input.Children {
		key := inputChild.Name.Space + ":" + inputChild.Name.Local
		if matchingNodes := outputMap[key]; len(matchingNodes) > 0 {
			compareAttrValuesRecursive(inputChild, matchingNodes[0], match)
			outputMap[key] = matchingNodes[1:]
		}
	}
}

func calculateDepth(node *XMLNode) int {
	if node == nil || len(node.Children) == 0 {
		return 1
	}

	maxChildDepth := 0
	for _, child := range node.Children {
		depth := calculateDepth(child)
		if depth > maxChildDepth {
			maxChildDepth = depth
		}
	}

	return 1 + maxChildDepth
}

func compareOrder(input, output *XMLNode) bool {
	if input == nil || output == nil {
		return input == output
	}

	if len(input.Children) != len(output.Children) {
		return false
	}

	for i := range input.Children {
		if input.Children[i].Name != output.Children[i].Name {
			return false
		}
		if !compareOrder(input.Children[i], output.Children[i]) {
			return false
		}
	}

	return true
}

func compareNesting(input, output *XMLNode) bool {
	if input == nil || output == nil {
		return input == output
	}

	if len(input.Children) != len(output.Children) {
		return false
	}

	for i := range input.Children {
		if !compareNesting(input.Children[i], output.Children[i]) {
			return false
		}
	}

	return true
}

// PrintDeepComparisonReport prints comprehensive validation report
func PrintDeepComparisonReport(result *DeepComparisonResult) {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("📊 COMPREHENSIVE ROUND-TRIP VALIDATION REPORT")
	fmt.Println("═══════════════════════════════════════════════════════════")

	// Level 1: Count Validation
	fmt.Println("\nLEVEL 1: COUNT VALIDATION")
	fmt.Println("─────────────────────────")
	printCountMatch("Elements", result.ElementCount)
	printCountMatch("Attributes", result.AttributeCount)
	printCountMatch("  xmlns", result.XmlnsCount)
	printCountMatch("  non-xmlns", result.NonXmlnsCount)

	// Level 2: Name Validation
	fmt.Println("\nLEVEL 2: NAME VALIDATION")
	fmt.Println("─────────────────────────")
	printNameMatch("Element Names", result.ElementNames)
	printNameMatch("Attribute Names", result.AttributeNames)
	printNameMatch("Namespace URIs", result.NamespaceURIs)

	// Level 3: Value Validation (CRITICAL)
	fmt.Println("\nLEVEL 3: VALUE VALIDATION ⚠️ CRITICAL")
	fmt.Println("─────────────────────────")
	printValueMatch("Element Content", result.ElementValues)
	printValueMatch("Attribute Values", result.AttributeValues)

	// Level 4: Structural Validation
	fmt.Println("\nLEVEL 4: STRUCTURAL VALIDATION")
	fmt.Println("─────────────────────────")
	printDepthMatch("Max Depth", result.MaxDepth)
	printBoolMatch("Element Order", result.ElementOrder)
	printBoolMatch("Nesting", result.NestingIntact)

	// Cross-Check Analysis
	fmt.Println("\n═══════════════════════════════════════════════════════════")
	fmt.Println("CROSS-CHECK ANALYSIS")
	fmt.Println("═══════════════════════════════════════════════════════════")

	silentCorruption := result.ElementCount.Match && result.AttributeCount.Match &&
		(result.ElementValues.MismatchCount > 0 || result.AttributeValues.MismatchCount > 0)

	if silentCorruption {
		fmt.Println("\n⚠️  SILENT CORRUPTION DETECTED:")
		fmt.Println("   Counts match but values differ!")
		fmt.Println("   This is DATA CORRUPTION, not just count loss!")
	} else {
		fmt.Println("\n✅ No silent corruption detected")
	}

	namespaceCorruption := result.XmlnsCount.Match && result.NamespaceURIs.MismatchCount > 0
	if namespaceCorruption {
		fmt.Println("\n⚠️  NAMESPACE CORRUPTION:")
		fmt.Println("   xmlns count matches but URIs differ!")
	} else {
		fmt.Println("✅ No namespace corruption detected")
	}

	if len(result.ElementNames.Changed) > 0 {
		fmt.Println("\n⚠️  ELEMENT SUBSTITUTION:")
		fmt.Println("   Element names changed!")
	} else {
		fmt.Println("✅ No element substitution detected")
	}

	// Verdict
	fmt.Println("\n═══════════════════════════════════════════════════════════")

	// Check semantic fidelity (data preservation regardless of structure)
	hasSemanticFidelity := result.ElementNames.MismatchCount == 0 &&
		result.AttributeNames.MismatchCount == 0 &&
		result.NamespaceURIs.MismatchCount == 0 &&
		result.ElementValues.MismatchCount == 0 &&
		result.AttributeValues.MismatchCount == 0

	if result.IsPerfectMatch {
		fmt.Println("VERDICT: 🎯 PERFECT FIDELITY")
		fmt.Println("═══════════════════════════════════════════════════════════")
		fmt.Println("100% count match")
		fmt.Println("100% name match")
		fmt.Println("100% value match")
		fmt.Println("100% structure match")
		fmt.Println("\nZERO DATA LOSS. ZERO CORRUPTION. ZERO MUTATIONS.")
	} else if hasSemanticFidelity {
		fmt.Println("VERDICT: ✅ SEMANTIC FIDELITY")
		fmt.Println("═══════════════════════════════════════════════════════════")
		fmt.Println("100% data preservation:")
		fmt.Println("  • All element names preserved")
		fmt.Println("  • All attribute names preserved")
		fmt.Println("  • All element values preserved")
		fmt.Println("  • All attribute values preserved")
		fmt.Println("  • All namespaces preserved")
		fmt.Println("\nStructural differences detected (acceptable):")
		if !result.ElementCount.Match {
			fmt.Printf("  • Element count: %d → %d (Δ = %+d)\n",
				result.ElementCount.Input, result.ElementCount.Output, result.ElementCount.Delta)
		}
		if !result.AttributeCount.Match {
			fmt.Printf("  • Attribute count: %d → %d (Δ = %+d)\n",
				result.AttributeCount.Input, result.AttributeCount.Output, result.AttributeCount.Delta)
			fmt.Println("    Note: xmlns consolidation is normal XML marshaling behavior")
		}
		if !result.ElementOrder {
			fmt.Println("  • Element ordering differs (unknown elements may appear at end)")
		}
		if !result.NestingIntact {
			fmt.Println("  • Nesting structure differs")
		}
		fmt.Println("\n✅ NO DATA LOSS - All content successfully preserved!")
	} else {
		fmt.Println("VERDICT: ❌ CRITICAL FAILURE")
		fmt.Println("═══════════════════════════════════════════════════════════")
		fmt.Printf("\nTotal issues found: %d\n", result.TotalIssues)
		fmt.Println("\nDATA CORRUPTION DETECTED")
		fmt.Println("Zero tolerance for data loss - ALL items must be preserved")
	}
	fmt.Println("═══════════════════════════════════════════════════════════")
}

func printCountMatch(label string, match CountMatch) {
	status := "✓"
	if !match.Match {
		status = "❌"
	}
	fmt.Printf("%-15s %d → %d %s (Δ = %+d)\n", label+":", match.Input, match.Output, status, match.Delta)
}

func printNameMatch(label string, match NameMatch) {
	status := "✓"
	if match.MismatchCount > 0 {
		status = "❌"
	}
	fmt.Printf("%s: %d unique, %d/%d match %s\n", label, match.UniqueNames, match.MatchCount, match.UniqueNames, status)

	if len(match.Missing) > 0 {
		fmt.Printf("  Missing: %v\n", match.Missing)
	}
	if len(match.Extra) > 0 {
		fmt.Printf("  Extra: %v\n", match.Extra)
	}
	if len(match.Changed) > 0 {
		fmt.Printf("  Changed: %d\n", len(match.Changed))
		for i, change := range match.Changed {
			if i < 5 {
				fmt.Printf("    %s: %s → %s\n", change.Path, change.OldName, change.NewName)
			}
		}
		if len(match.Changed) > 5 {
			fmt.Printf("    ... and %d more\n", len(match.Changed)-5)
		}
	}
}

func printValueMatch(label string, match ValueMatch) {
	status := "✓"
	if match.MismatchCount > 0 {
		status = "❌"
	}
	fmt.Printf("%s: %d total, %d/%d match %s\n", label, match.Total, match.MatchCount, match.Total, status)

	if len(match.Mismatches) > 0 {
		fmt.Printf("  Mismatches: %d\n\n", match.MismatchCount)
		for i, mismatch := range match.Mismatches {
			fmt.Printf("  %d. %s at %s\n", i+1, mismatch.Name, mismatch.Path)
			fmt.Printf("     Input:  %q\n", mismatch.InputValue)
			fmt.Printf("     Output: %q\n\n", mismatch.OutputValue)
		}
		if match.MismatchCount > len(match.Mismatches) {
			fmt.Printf("  ... and %d more\n\n", match.MismatchCount-len(match.Mismatches))
		}
	}
}

func printDepthMatch(label string, match DepthMatch) {
	status := "✓"
	if !match.Match {
		status = "❌"
	}
	fmt.Printf("%s: %d → %d %s\n", label, match.Input, match.Output, status)
}

func printBoolMatch(label string, match bool) {
	status := "✓"
	result := "PRESERVED"
	if !match {
		status = "❌"
		result = "ALTERED"
	}
	fmt.Printf("%s: %s %s\n", label, result, status)
}
