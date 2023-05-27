package replacements

import (
	"os"
	"regexp"
	"strings"
)

// settingsPattern supports replacing ##NAME## with settings values
var settingsPattern = regexp.MustCompile(`(##)([A-Z0-9_\-\.]*)(##)`)

// capture group that contains the key, must match settingsPattern
const settingsCaptureGroupKey = 2

// envPattern supports replacing %%NAME%% with environment variables
var envPattern = regexp.MustCompile(`(%%)([A-Z0-9_\-\.]*)(%%)`)

const envCaptureGroupKey = 2

// replaces placeholders in a string with config values (settings or env vars)
func ReplacePlaceholders(s string, settings map[string]string) string {
	if s == "" {
		return ""
	}
	// replace with settings values
	replacements := settingsPattern.FindAllStringSubmatch(s, -1)
	for _, item := range replacements {
		key := item[settingsCaptureGroupKey]
		if value, ok := settings[key]; ok {
			s = strings.Replace(s, "##"+key+"##", value, -1)
		}
	}
	// replace with environment variables
	replacements = envPattern.FindAllStringSubmatch(s, -1)
	for _, item := range replacements {
		key := item[envCaptureGroupKey]
		if value, ok := os.LookupEnv(key); ok {
			s = strings.Replace(s, "%%"+key+"%%", value, -1)
		}
	}
	return s
}
