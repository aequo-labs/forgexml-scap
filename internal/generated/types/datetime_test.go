// Package types provides custom types for XSD built-in types with 100% fidelity
package types

import (
	"encoding/xml"
	"strings"
	"testing"
	"time"
)

// wrapper is a named struct type for XML marshaling tests
type wrapper struct {
	XMLName xml.Name `xml:"wrapper"`
	DT      DateTime `xml:"dt"`
}

// TestDateTime_UnmarshalXML tests XML element unmarshaling
func TestDateTime_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		xml      string
		wantYear int
		wantErr  bool
	}{
		{
			name:     "RFC3339 with Z timezone",
			xml:      `<wrapper><dt>2024-03-15T10:30:00Z</dt></wrapper>`,
			wantYear: 2024,
		},
		{
			name:     "RFC3339 with offset",
			xml:      `<wrapper><dt>2024-03-15T10:30:00-05:00</dt></wrapper>`,
			wantYear: 2024,
		},
		{
			name:     "ISO8601 without timezone",
			xml:      `<wrapper><dt>2024-03-15T10:30:00</dt></wrapper>`,
			wantYear: 2024,
		},
		{
			name:     "Date only",
			xml:      `<wrapper><dt>2024-03-15</dt></wrapper>`,
			wantYear: 2024,
		},
		{
			name:     "With milliseconds",
			xml:      `<wrapper><dt>2024-03-15T10:30:00.123</dt></wrapper>`,
			wantYear: 2024,
		},
		{
			name:     "With microseconds",
			xml:      `<wrapper><dt>2024-03-15T10:30:00.123456</dt></wrapper>`,
			wantYear: 2024,
		},
		{
			name:     "With nanoseconds",
			xml:      `<wrapper><dt>2024-03-15T10:30:00.123456789</dt></wrapper>`,
			wantYear: 2024,
		},
		{
			name:    "Invalid format",
			xml:     `<wrapper><dt>not-a-date</dt></wrapper>`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var w wrapper
			err := xml.Unmarshal([]byte(tt.xml), &w)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if w.DT.Value.Year() != tt.wantYear {
				t.Errorf("year = %d, want %d", w.DT.Value.Year(), tt.wantYear)
			}
		})
	}
}

// TestDateTime_MarshalXML tests XML element marshaling preserves original
func TestDateTime_MarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		dt       DateTime
		contains string
	}{
		{
			name: "preserves original format",
			dt: DateTime{
				Original: "2024-03-15T10:30:00",
				Value:    time.Date(2024, 3, 15, 10, 30, 0, 0, time.UTC),
			},
			contains: "2024-03-15T10:30:00",
		},
		{
			name: "falls back to RFC3339 when no original",
			dt: DateTime{
				Value: time.Date(2024, 3, 15, 10, 30, 0, 0, time.UTC),
			},
			contains: "2024-03-15T10:30:00Z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := wrapper{DT: tt.dt}
			data, err := xml.Marshal(&w)
			if err != nil {
				t.Fatalf("marshal error: %v", err)
			}
			if got := string(data); !strings.Contains(got, tt.contains) {
				t.Errorf("marshaled = %s, want to contain %s", got, tt.contains)
			}
		})
	}
}

// TestDateTime_UnmarshalText tests attribute unmarshaling
func TestDateTime_UnmarshalText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		wantYear int
		wantErr  bool
	}{
		{
			name:     "RFC3339",
			text:     "2024-03-15T10:30:00Z",
			wantYear: 2024,
		},
		{
			name:     "ISO8601",
			text:     "2024-03-15T10:30:00",
			wantYear: 2024,
		},
		{
			name:    "Invalid",
			text:    "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dt DateTime
			err := dt.UnmarshalText([]byte(tt.text))
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if dt.Value.Year() != tt.wantYear {
				t.Errorf("year = %d, want %d", dt.Value.Year(), tt.wantYear)
			}
			if dt.Original != tt.text {
				t.Errorf("Original = %s, want %s", dt.Original, tt.text)
			}
		})
	}
}

// TestDateTime_MarshalText tests attribute marshaling
func TestDateTime_MarshalText(t *testing.T) {
	tests := []struct {
		name string
		dt   DateTime
		want string
	}{
		{
			name: "preserves original",
			dt: DateTime{
				Original: "2024-03-15T10:30:00",
				Value:    time.Date(2024, 3, 15, 10, 30, 0, 0, time.UTC),
			},
			want: "2024-03-15T10:30:00",
		},
		{
			name: "falls back to RFC3339",
			dt: DateTime{
				Value: time.Date(2024, 3, 15, 10, 30, 0, 0, time.UTC),
			},
			want: "2024-03-15T10:30:00Z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.dt.MarshalText()
			if err != nil {
				t.Fatalf("MarshalText error: %v", err)
			}
			if string(got) != tt.want {
				t.Errorf("MarshalText = %s, want %s", got, tt.want)
			}
		})
	}
}

// TestDateTime_String tests the String method
func TestDateTime_String(t *testing.T) {
	tests := []struct {
		name string
		dt   DateTime
		want string
	}{
		{
			name: "returns original when set",
			dt: DateTime{
				Original: "2024-03-15T10:30:00",
				Value:    time.Date(2024, 3, 15, 10, 30, 0, 0, time.UTC),
			},
			want: "2024-03-15T10:30:00",
		},
		{
			name: "returns RFC3339 when no original",
			dt: DateTime{
				Value: time.Date(2024, 3, 15, 10, 30, 0, 0, time.UTC),
			},
			want: "2024-03-15T10:30:00Z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.String(); got != tt.want {
				t.Errorf("String() = %s, want %s", got, tt.want)
			}
		})
	}
}

// TestDateTime_Time tests the Time method
func TestDateTime_Time(t *testing.T) {
	expected := time.Date(2024, 3, 15, 10, 30, 0, 0, time.UTC)
	dt := DateTime{
		Value:    expected,
		Original: "2024-03-15T10:30:00Z",
	}
	if got := dt.Time(); !got.Equal(expected) {
		t.Errorf("Time() = %v, want %v", got, expected)
	}
}

// TestDateTime_RoundTrip tests full XML round-trip fidelity
func TestDateTime_RoundTrip(t *testing.T) {
	originals := []string{
		"2024-03-15T10:30:00Z",
		"2024-03-15T10:30:00-05:00",
		"2024-03-15T10:30:00",
		"2024-03-15",
		"2024-03-15T10:30:00.123",
	}

	for _, orig := range originals {
		t.Run(orig, func(t *testing.T) {
			// Unmarshal
			xmlIn := `<wrapper><dt>` + orig + `</dt></wrapper>`
			var w wrapper
			if err := xml.Unmarshal([]byte(xmlIn), &w); err != nil {
				t.Fatalf("unmarshal error: %v", err)
			}

			// Marshal back
			data, err := xml.Marshal(&w)
			if err != nil {
				t.Fatalf("marshal error: %v", err)
			}

			// Verify original preserved
			if got := string(data); !strings.Contains(got, orig) {
				t.Errorf("round-trip lost original: got %s, want to contain %s", got, orig)
			}
		})
	}
}
