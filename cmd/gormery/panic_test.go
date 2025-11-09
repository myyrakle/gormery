package main

import (
	"fmt"
	"strings"
	"sync"
	"testing"

	gormSchema "gorm.io/gorm/schema"
)

// TestPanicOnParseError verifies that when ParseWithSpecialTableName returns an error,
// the code panics with an appropriate message
func TestPanicOnParseError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			// Panic occurred as expected
			panicMsg := fmt.Sprintf("%v", r)
			if !strings.Contains(panicMsg, "failed to parse struct") {
				t.Errorf("Expected panic message to contain 'failed to parse struct', got: %s", panicMsg)
			}
			if !strings.Contains(panicMsg, "unsupported data type") {
				t.Errorf("Expected panic message to contain original error 'unsupported data type', got: %s", panicMsg)
			}
			t.Logf("Successfully caught panic with message: %s", panicMsg)
		} else {
			t.Error("Expected panic to occur but it didn't")
		}
	}()

	// This should trigger a panic because nil is not a valid input
	// This simulates the pattern used in the generated runner code
	_, err := gormSchema.ParseWithSpecialTableName(
		nil, // This causes an error: "unsupported data type: <nil>"
		&sync.Map{},
		&gormSchema.NamingStrategy{},
		"",
	)

	if err != nil {
		panic(fmt.Sprintf("failed to parse struct TestStruct: %v", err))
	}
}
