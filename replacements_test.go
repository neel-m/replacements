package replacements

import (
	"testing"
)

func TestReplacePlaceholders(t *testing.T) {
	var result string
	// test empty string
	if ReplacePlaceholders("", nil) != "" {
		t.Error("Expected empty string")
	}
	// test no replacements
	if ReplacePlaceholders("boo", nil) != "boo" {
		t.Error("Expected boo")
	}
	// test no replacement when no settings
	result = ReplacePlaceholders("##BOO##%%BAR%%", nil)
	if result != "##BOO##%%BAR%%" {
		t.Error("Expected ##BOO##%%BAR%% got:", result)
	}
	// test settings replacement
	choices := make(map[string]string)
	choices["BOO"] = "bar"
	if ReplacePlaceholders("##BOO##", choices) != "bar" {
		t.Error("Expected bar")
	}
	// test empty environment replacement
	result = ReplacePlaceholders("%%NOTVALID%%", nil)
	if result != "%%NOTVALID%%" {
		t.Error("Expected empty string, got:", result)
	}
	// test multiple replacements
	choices["BOO"] = "bar"
	t.Setenv("BAR", "baz")
	result = ReplacePlaceholders("##BOO##%%BAR%%", choices)
	if result != "barbaz" {
		t.Error("Expected barbaz got:", result)
	}
}
