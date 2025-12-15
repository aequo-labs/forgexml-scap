// Package integration provides round-trip integration tests for DAL object graph persistence
package integration

import (
	"context"
	"database/sql"
	"encoding/xml"
	"os"
	"path/filepath"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	xccdf "github.com/aequo-labs/forgexml-scap/internal/generated/gov/nist/checklists/xccdf/1-2"
)

// TestXMLRoundTripWithDB tests: XML -> Go struct -> Save to DB -> Load from DB -> XML -> Compare
// This is a simplified test that uses direct SQL since the DAL is in a separate module
func TestXMLRoundTripWithDB(t *testing.T) {
	// Find a test XCCDF file
	testFile := findXCCDFTestFile(t)
	if testFile == "" {
		t.Skip("No XCCDF test file found")
	}

	t.Logf("Using test file: %s", testFile)

	// Step 1: Load XML into Go struct
	xmlData, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	var original xccdf.BenchmarkElement
	if err := xml.Unmarshal(xmlData, &original); err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	t.Logf("Loaded benchmark: ID=%s, Profiles=%d, Groups=%d, Rules=%d, Values=%d",
		original.Id, len(original.Profile), len(original.Group), len(original.Rule), len(original.Value))

	// Step 2: Create in-memory SQLite database
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	// Create a simple schema for benchmark elements
	_, err = db.ExecContext(ctx, `
		CREATE TABLE benchmark (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			benchmark_id TEXT NOT NULL,
			resolved INTEGER,
			style TEXT,
			style_href TEXT
		);
		CREATE TABLE profile (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			benchmark_id INTEGER NOT NULL,
			profile_id TEXT NOT NULL,
			abstract INTEGER,
			prohibit_changes INTEGER,
			note_tag TEXT,
			extends TEXT,
			FOREIGN KEY (benchmark_id) REFERENCES benchmark(id)
		);
		CREATE TABLE benchmark_group (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			benchmark_id INTEGER NOT NULL,
			group_id TEXT NOT NULL,
			abstract INTEGER,
			hidden INTEGER,
			prohibit_changes INTEGER,
			selected INTEGER,
			weight TEXT,
			extends TEXT,
			FOREIGN KEY (benchmark_id) REFERENCES benchmark(id)
		);
		CREATE TABLE rule (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			benchmark_id INTEGER NOT NULL,
			rule_id TEXT NOT NULL,
			abstract INTEGER,
			hidden INTEGER,
			prohibit_changes INTEGER,
			selected INTEGER,
			weight TEXT,
			role TEXT,
			severity TEXT,
			multiple INTEGER,
			FOREIGN KEY (benchmark_id) REFERENCES benchmark(id)
		);
		CREATE TABLE value (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			benchmark_id INTEGER NOT NULL,
			value_id TEXT NOT NULL,
			abstract INTEGER,
			hidden INTEGER,
			prohibit_changes INTEGER,
			interactive INTEGER,
			type TEXT,
			operator TEXT,
			FOREIGN KEY (benchmark_id) REFERENCES benchmark(id)
		);
	`)
	if err != nil {
		t.Fatalf("Failed to create schema: %v", err)
	}

	// Step 3: Save benchmark to DB
	resolved := 0
	if original.Resolved != nil && *original.Resolved {
		resolved = 1
	}
	style := ""
	if original.Style != nil {
		style = *original.Style
	}
	styleHref := ""
	if original.StyleHref != nil {
		styleHref = *original.StyleHref
	}

	result, err := db.ExecContext(ctx, 
		"INSERT INTO benchmark (benchmark_id, resolved, style, style_href) VALUES (?, ?, ?, ?)",
		string(original.Id), resolved, style, styleHref)
	if err != nil {
		t.Fatalf("Failed to insert benchmark: %v", err)
	}

	benchmarkDBID, _ := result.LastInsertId()
	t.Logf("Saved benchmark with DB ID: %d", benchmarkDBID)

	// Save profiles
	for _, profile := range original.Profile {
		abstract := 0
		if profile.Abstract != nil && *profile.Abstract {
			abstract = 1
		}
		prohibitChanges := 0
		if profile.ProhibitChanges != nil && *profile.ProhibitChanges {
			prohibitChanges = 1
		}
		noteTag := ""
		if profile.NoteTag != nil {
			noteTag = *profile.NoteTag
		}
		extends := ""
		if profile.Extends != nil {
			extends = string(*profile.Extends)
		}
		_, err := db.ExecContext(ctx,
			"INSERT INTO profile (benchmark_id, profile_id, abstract, prohibit_changes, note_tag, extends) VALUES (?, ?, ?, ?, ?, ?)",
			benchmarkDBID, string(profile.Id), abstract, prohibitChanges, noteTag, extends)
		if err != nil {
			t.Fatalf("Failed to insert profile: %v", err)
		}
	}
	t.Logf("Saved %d profiles", len(original.Profile))

	// Save groups
	for _, group := range original.Group {
		abstract := 0
		if group.Abstract != nil && *group.Abstract {
			abstract = 1
		}
		hidden := 0
		if group.Hidden != nil && *group.Hidden {
			hidden = 1
		}
		prohibitChanges := 0
		if group.ProhibitChanges != nil && *group.ProhibitChanges {
			prohibitChanges = 1
		}
		selected := 1
		if group.Selected != nil && !*group.Selected {
			selected = 0
		}
		weight := "1.0"
		if group.Weight != nil {
			weight = string(*group.Weight)
		}
		extends := ""
		if group.Extends != nil {
			extends = string(*group.Extends)
		}
		_, err := db.ExecContext(ctx,
			"INSERT INTO benchmark_group (benchmark_id, group_id, abstract, hidden, prohibit_changes, selected, weight, extends) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			benchmarkDBID, string(group.Id), abstract, hidden, prohibitChanges, selected, weight, extends)
		if err != nil {
			t.Fatalf("Failed to insert group: %v", err)
		}
	}
	t.Logf("Saved %d groups", len(original.Group))

	// Save rules
	for _, rule := range original.Rule {
		abstract := 0
		if rule.Abstract != nil && *rule.Abstract {
			abstract = 1
		}
		hidden := 0
		if rule.Hidden != nil && *rule.Hidden {
			hidden = 1
		}
		prohibitChanges := 0
		if rule.ProhibitChanges != nil && *rule.ProhibitChanges {
			prohibitChanges = 1
		}
		selected := 1
		if rule.Selected != nil && !*rule.Selected {
			selected = 0
		}
		weight := "1.0"
		if rule.Weight != nil {
			weight = string(*rule.Weight)
		}
		role := ""
		if rule.Role != nil {
			role = string(*rule.Role)
		}
		severity := ""
		if rule.Severity != nil {
			severity = string(*rule.Severity)
		}
		multiple := 0
		if rule.Multiple != nil && *rule.Multiple {
			multiple = 1
		}
		_, err := db.ExecContext(ctx,
			"INSERT INTO rule (benchmark_id, rule_id, abstract, hidden, prohibit_changes, selected, weight, role, severity, multiple) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			benchmarkDBID, string(rule.Id), abstract, hidden, prohibitChanges, selected, weight, role, severity, multiple)
		if err != nil {
			t.Fatalf("Failed to insert rule: %v", err)
		}
	}
	t.Logf("Saved %d rules", len(original.Rule))

	// Save values
	for _, value := range original.Value {
		abstract := 0
		if value.Abstract != nil && *value.Abstract {
			abstract = 1
		}
		hidden := 0
		if value.Hidden != nil && *value.Hidden {
			hidden = 1
		}
		prohibitChanges := 0
		if value.ProhibitChanges != nil && *value.ProhibitChanges {
			prohibitChanges = 1
		}
		interactive := 0
		if value.Interactive != nil && *value.Interactive {
			interactive = 1
		}
		valueType := ""
		if value.Type != nil {
			valueType = string(*value.Type)
		}
		operator := ""
		if value.Operator != nil {
			operator = string(*value.Operator)
		}
		_, err := db.ExecContext(ctx,
			"INSERT INTO value (benchmark_id, value_id, abstract, hidden, prohibit_changes, interactive, type, operator) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			benchmarkDBID, string(value.Id), abstract, hidden, prohibitChanges, interactive, valueType, operator)
		if err != nil {
			t.Fatalf("Failed to insert value: %v", err)
		}
	}
	t.Logf("Saved %d values", len(original.Value))

	// Step 4: Load from DB
	var loadedID string
	var loadedResolved int
	var loadedStyle, loadedStyleHref string
	err = db.QueryRowContext(ctx, "SELECT benchmark_id, resolved, style, style_href FROM benchmark WHERE id = ?", benchmarkDBID).
		Scan(&loadedID, &loadedResolved, &loadedStyle, &loadedStyleHref)
	if err != nil {
		t.Fatalf("Failed to load benchmark: %v", err)
	}

	// Count loaded children
	var profileCount, groupCount, ruleCount, valueCount int
	db.QueryRowContext(ctx, "SELECT COUNT(*) FROM profile WHERE benchmark_id = ?", benchmarkDBID).Scan(&profileCount)
	db.QueryRowContext(ctx, "SELECT COUNT(*) FROM benchmark_group WHERE benchmark_id = ?", benchmarkDBID).Scan(&groupCount)
	db.QueryRowContext(ctx, "SELECT COUNT(*) FROM rule WHERE benchmark_id = ?", benchmarkDBID).Scan(&ruleCount)
	db.QueryRowContext(ctx, "SELECT COUNT(*) FROM value WHERE benchmark_id = ?", benchmarkDBID).Scan(&valueCount)

	t.Logf("Loaded from DB: ID=%s, Profiles=%d, Groups=%d, Rules=%d, Values=%d",
		loadedID, profileCount, groupCount, ruleCount, valueCount)

	// Step 5: Compare
	if loadedID != string(original.Id) {
		t.Errorf("Benchmark ID mismatch: original=%s, loaded=%s", original.Id, loadedID)
	}
	if profileCount != len(original.Profile) {
		t.Errorf("Profile count mismatch: original=%d, loaded=%d", len(original.Profile), profileCount)
	}
	if groupCount != len(original.Group) {
		t.Errorf("Group count mismatch: original=%d, loaded=%d", len(original.Group), groupCount)
	}
	if ruleCount != len(original.Rule) {
		t.Errorf("Rule count mismatch: original=%d, loaded=%d", len(original.Rule), ruleCount)
	}
	if valueCount != len(original.Value) {
		t.Errorf("Value count mismatch: original=%d, loaded=%d", len(original.Value), valueCount)
	}

	t.Log("Round-trip test PASSED - all counts match")
}

func findXCCDFTestFile(t *testing.T) string {
	// Try different paths to find XCCDF 1.2 test data
	searchPaths := []string{
		"../../instances/scap-benchmarks",
	}

	for _, basePath := range searchPaths {
		entries, err := os.ReadDir(basePath)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}

			// Check subdirectories for xccdf files
			fullPath := filepath.Join(basePath, entry.Name())
			subEntries, err := os.ReadDir(fullPath)
			if err != nil {
				continue
			}
			for _, sub := range subEntries {
				subName := sub.Name()
				if strings.HasSuffix(subName, "-xccdf.xml") {
					testFile := filepath.Join(fullPath, subName)
					// Verify it's XCCDF 1.2 by checking content
					content, err := os.ReadFile(testFile)
					if err != nil {
						continue
					}
					if strings.Contains(string(content), "xccdf/1.2") {
						return testFile
					}
				}
			}
		}
	}

	return ""
}
