package builtins

import (
	"testing"
)

func TestSpecificModule(t *testing.T) {
	tests := []struct {
		version string
		module  string
		exist   bool
	}{
		{"5.99.99", "freelist", true},
		{"6.0.0", "freelist", false},
		{"0.99.99", "v8", false},
		{"1.0.0", "v8", true},
		{"1.0.99", "process", false},
		{"1.1.0", "process", true},
		{"8.0.99", "async_hooks", false},
		{"8.1.0", "async_hooks", true},
		{"8.3.99", "http2", false},
		{"8.4.0", "http2", true},
		{"8.4.99", "perf_hooks", false},
		{"8.5.0", "perf_hooks", true},
	}
	for _, tc := range tests {
		found := false
		builtins, err := GetVersion(tc.version)
		if err != nil {
			t.Error(err)
		}
		for _, b := range builtins {
			if b == tc.module {
				found = true
			}
		}
		if found != tc.exist {
			t.Errorf("Version check failed for module %s", tc.module)
		}
	}
}

func TestBuiltinsLength(t *testing.T) {
	tests := []struct {
		version string
		length  int
	}{
		{"0.12.0", 33},
		{"8.5.0", 37},
	}
	for _, tc := range tests {
		builtins, err := GetVersion(tc.version)
		if err != nil {
			t.Error(err)
		}
		if len(builtins) != tc.length {
			t.Errorf("Builtins length check failed for versoin %s", tc.version)
		}
	}
}
