package hcio_test

import (
	"testing"

	"github.com/dansage/hcio"
)

// TestDefaultOptions verifies that checks created without options are given the default options
func TestDefaultOptions(t *testing.T) {
	// create a check without any specified options
	check := hcio.NewCheck("00000000-0000-0000-0000-000000000000")

	// verify the options are set to the defaults
	if check.Options.BaseUrl != hcio.DefaultOptions.BaseUrl {
		t.Errorf("option BaseUrl not set to default value, got: %q, want: %q", check.Options.BaseUrl, hcio.DefaultOptions.BaseUrl)
	}
	if check.Options.MaxRetries != hcio.DefaultOptions.MaxRetries {
		t.Errorf("option MaxRetries not set to default value, got: %d, want: %d", check.Options.MaxRetries, hcio.DefaultOptions.MaxRetries)
	}
	if check.Options.UserAgent != hcio.DefaultOptions.UserAgent {
		t.Errorf("option UserAgent not set to default value, got: %q, want: %q", check.Options.UserAgent, hcio.DefaultOptions.UserAgent)
	}
}

// TestMergedOptions verifies that checks created with some but not all options are given a merged set of options
func TestMergedOptions(t *testing.T) {
	opts := hcio.Options{
		UserAgent: "HCio tests",
	}

	// create a check with a user agent specified
	check := hcio.NewCheck("00000000-0000-0000-0000-000000000000", opts)

	// verify the user agent matches the value specified
	if check.Options.UserAgent != opts.UserAgent {
		t.Errorf("option UserAgent not set to user specified value, got: %q, want: %q", check.Options.UserAgent, opts.UserAgent)
	}

	// verify the options are set to the defaults
	if check.Options.BaseUrl != hcio.DefaultOptions.BaseUrl {
		t.Errorf("option BaseUrl not set to default value, got: %q, want: %q", check.Options.BaseUrl, hcio.DefaultOptions.BaseUrl)
	}
	if check.Options.MaxRetries != hcio.DefaultOptions.MaxRetries {
		t.Errorf("option MaxRetries not set to default value, got: %q, want: %q", check.Options.MaxRetries, hcio.DefaultOptions.MaxRetries)
	}
}
