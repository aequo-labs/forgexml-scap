// Package types provides custom types for XSD built-in types with 100% fidelity
package types

import (
	"encoding/xml"
	"time"
)

// DateTime represents an XSD dateTime value with support for multiple formats
// and preserves the original string representation for exact round-trip fidelity.
// Supports:
// - RFC3339 with timezone: "2006-01-02T15:04:05Z" or "2006-01-02T15:04:05-07:00"
// - ISO 8601 without timezone: "2006-01-02T15:04:05"
// - Date only: "2006-01-02"
type DateTime struct {
	Value    time.Time // Parsed time value
	Original string    // Original string representation for fidelity
}

// Common datetime formats in order of specificity
var dateTimeFormats = []string{
	time.RFC3339Nano,                // "2006-01-02T15:04:05.999999999Z07:00"
	time.RFC3339,                    // "2006-01-02T15:04:05Z07:00"
	"2006-01-02T15:04:05.999999999", // Nanoseconds without timezone
	"2006-01-02T15:04:05.999999",    // Microseconds without timezone
	"2006-01-02T15:04:05.999",       // Milliseconds without timezone
	"2006-01-02T15:04:05",           // Seconds without timezone (OVAL format)
	"2006-01-02",                    // Date only
}

// UnmarshalXML implements xml.Unmarshaler for flexible datetime parsing
func (dt *DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}

	// Preserve original for exact round-trip
	dt.Original = content

	// Try parsing with each format
	var lastErr error
	for _, format := range dateTimeFormats {
		t, err := time.Parse(format, content)
		if err == nil {
			dt.Value = t
			return nil
		}
		lastErr = err
	}

	// If all formats failed, return the last error
	return lastErr
}

// MarshalXML implements xml.Marshaler to preserve original format
func (dt DateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Use original string if available for exact fidelity
	if dt.Original != "" {
		return e.EncodeElement(dt.Original, start)
	}

	// Fallback to RFC3339 if no original (shouldn't happen in round-trip)
	return e.EncodeElement(dt.Value.Format(time.RFC3339), start)
}

// UnmarshalText implements encoding.TextUnmarshaler for attribute parsing
// This is required for DateTime fields used as XML attributes
func (dt *DateTime) UnmarshalText(text []byte) error {
	content := string(text)
	dt.Original = content

	// Try parsing with each format
	var lastErr error
	for _, format := range dateTimeFormats {
		t, err := time.Parse(format, content)
		if err == nil {
			dt.Value = t
			return nil
		}
		lastErr = err
	}

	// If all formats failed, return the last error
	return lastErr
}

// MarshalText implements encoding.TextMarshaler for attribute marshaling
// This is required for DateTime fields used as XML attributes
func (dt DateTime) MarshalText() ([]byte, error) {
	// Use original string if available for exact fidelity
	if dt.Original != "" {
		return []byte(dt.Original), nil
	}

	// Fallback to RFC3339 if no original (shouldn't happen in round-trip)
	return []byte(dt.Value.Format(time.RFC3339)), nil
}

// String returns the original string representation
func (dt DateTime) String() string {
	if dt.Original != "" {
		return dt.Original
	}
	return dt.Value.Format(time.RFC3339)
}

// Time returns the parsed time.Time value
func (dt DateTime) Time() time.Time {
	return dt.Value
}
