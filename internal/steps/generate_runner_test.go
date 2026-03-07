package steps

import (
	"strings"
	"testing"
)

func TestGenerateCodeForTarget_WithError(t *testing.T) {
	target := ProecssFileContext{
		filename:    "test/test.go",
		packageName: "test",
		structName:  "TestStruct",
		entityParam: nil,
	}

	code := generateCodeForTarget(0, target)

	// Verify that the generated code contains error handling with panic
	if !strings.Contains(code, "if err != nil {") {
		t.Error("Expected generated code to contain error check")
	}

	if !strings.Contains(code, "panic(fmt.Sprintf(") {
		t.Error("Expected generated code to contain panic call")
	}

	if !strings.Contains(code, "failed to parse struct TestStruct") {
		t.Error("Expected panic message to contain struct name")
	}

	// Verify that createGormFile is called unconditionally (after error check)
	if !strings.Contains(code, "createGormFile(target_0") {
		t.Error("Expected generated code to call createGormFile")
	}

	// Verify that the old pattern (if err == nil) is NOT present
	if strings.Contains(code, "if err == nil") {
		t.Error("Generated code should not contain 'if err == nil' pattern")
	}
}
