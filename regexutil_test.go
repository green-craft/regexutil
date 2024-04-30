package regexutil

import (
	"testing"
)

func TestDropPrefixAndUnderscores(t *testing.T) {

	result, _ := DropPrefixAndUnderscores("SomePrefix_value1", 1)
	if result != "value1" {
		t.Errorf("Expected SomePrefix value1, got %s", result)
	}

	result, _ = DropPrefixAndUnderscores("SomePrefix_value1_value2", 2)
	if result != "value1 value2" {
		t.Errorf("Expected value1 value2, got %s", result)
	}

	result, _ = DropPrefixAndUnderscores("SomePrefix_value1_value2_value3", 3)
	if result != "value1 value2 value3" {
		t.Errorf("Expected value1 value2 value3, got %s", result)
	}

	// This should not match and will error out
	var err error
	result, err = DropPrefixAndUnderscores("Some-Prefix_value1_value2", 2)
	if err == nil {
		t.Errorf("Expected no matches, got %s", result)
	}

}
