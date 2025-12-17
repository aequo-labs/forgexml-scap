package state

import (
	"fmt"
	"testing"
)

func TestPrintRootTypes(t *testing.T) {
	s := NewXMLDocumentState()
	types := s.GetRootElementTypes()
	fmt.Printf("Root types count: %d\n", len(types))
	for i, typ := range types {
		fmt.Printf("  %d: %s\n", i+1, typ)
	}
	if len(types) < 2 {
		t.Errorf("Expected more than 1 root type, got %d", len(types))
	}
}
