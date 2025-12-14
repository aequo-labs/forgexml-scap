package template

import (
	"encoding/json"
	"fmt"
	"html/template"
	"strings"
	"time"
)

// GetStandardFuncMap returns a template.FuncMap with commonly used formatting functions
// These functions are useful across HTML templates, email templates, API responses, and logging
func GetStandardFuncMap() template.FuncMap {
	return template.FuncMap{
		"json": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"formatNumber": func(v interface{}) string {
			switch val := v.(type) {
			case int:
				return addCommas(fmt.Sprintf("%d", val))
			case int64:
				return addCommas(fmt.Sprintf("%d", val))
			case float64:
				return fmt.Sprintf("%.2f", val)
			case float32:
				return fmt.Sprintf("%.2f", val)
			default:
				return fmt.Sprintf("%v", val)
			}
		},
		"formatCurrency": func(v interface{}) string {
			switch val := v.(type) {
			case float64:
				return fmt.Sprintf("$%.2f", val)
			case float32:
				return fmt.Sprintf("$%.2f", val)
			case int:
				return fmt.Sprintf("$%d.00", val)
			case int64:
				return fmt.Sprintf("$%d.00", val)
			default:
				return fmt.Sprintf("$%v", val)
			}
		},
		"formatDate": func(t time.Time) string {
			return t.Format("2006-01-02")
		},
		"formatDateTime": func(t time.Time) string {
			return t.Format("2006-01-02 15:04:05")
		},
		"formatBytes": func(bytes int64) string {
			const unit = 1024
			if bytes < unit {
				return fmt.Sprintf("%d B", bytes)
			}
			div, exp := int64(unit), 0
			for n := bytes / unit; n >= unit; n /= unit {
				div *= unit
				exp++
			}
			return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
		},
		"formatPercent": func(v float64) string {
			return fmt.Sprintf("%.1f%%", v*100)
		},
		"toUpper":   strings.ToUpper,
		"toLower":   strings.ToLower,
		"truncate": func(s string, length int) string {
			if len(s) <= length {
				return s
			}
			return s[:length] + "..."
		},
		"add": func(a, b interface{}) interface{} {
			switch va := a.(type) {
			case int:
				if vb, ok := b.(int); ok {
					return va + vb
				}
			case int64:
				if vb, ok := b.(int64); ok {
					return va + vb
				}
			case float64:
				if vb, ok := b.(float64); ok {
					return va + vb
				}
			case float32:
				if vb, ok := b.(float32); ok {
					return va + vb
				}
			}
			return a
		},
		"multiply": func(a, b interface{}) interface{} {
			switch va := a.(type) {
			case int:
				if vb, ok := b.(int); ok {
					return va * vb
				}
			case int64:
				if vb, ok := b.(int64); ok {
					return va * vb
				}
			case float64:
				if vb, ok := b.(float64); ok {
					return va * vb
				}
			case float32:
				if vb, ok := b.(float32); ok {
					return va * vb
				}
			}
			return a
		},
	}
}

// Individual formatting functions for direct use in Go code (non-template contexts)

// FormatNumber formats a number with comma separators
func FormatNumber(v interface{}) string {
	switch val := v.(type) {
	case int:
		return addCommas(fmt.Sprintf("%d", val))
	case int64:
		return addCommas(fmt.Sprintf("%d", val))
	case float64:
		return fmt.Sprintf("%.2f", val)
	case float32:
		return fmt.Sprintf("%.2f", val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

// addCommas adds comma separators to a number string
func addCommas(s string) string {
	// Handle negative numbers
	negative := false
	if len(s) > 0 && s[0] == '-' {
		negative = true
		s = s[1:]
	}
	
	// Add commas from right to left
	n := len(s)
	if n <= 3 {
		if negative {
			return "-" + s
		}
		return s
	}
	
	// Build result with commas
	result := ""
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			result += ","
		}
		result += string(s[i])
	}
	
	if negative {
		return "-" + result
	}
	return result
}

// FormatCurrency formats a value as currency with $ symbol
func FormatCurrency(v interface{}) string {
	switch val := v.(type) {
	case float64:
		return fmt.Sprintf("$%.2f", val)
	case float32:
		return fmt.Sprintf("$%.2f", val)
	case int:
		return fmt.Sprintf("$%d.00", val)
	case int64:
		return fmt.Sprintf("$%d.00", val)
	default:
		return fmt.Sprintf("$%v", val)
	}
}

// FormatBytes formats byte counts in human-readable format (B, KB, MB, GB, etc.)
func FormatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// FormatPercent formats a decimal as a percentage
func FormatPercent(v float64) string {
	return fmt.Sprintf("%.1f%%", v*100)
}

// FormatDate formats a time.Time as YYYY-MM-DD
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatDateTime formats a time.Time as YYYY-MM-DD HH:MM:SS
func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}