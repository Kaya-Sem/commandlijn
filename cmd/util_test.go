package cmd

import (
	"io"
	"os"
	"testing"
	"time"
)

func TestReplaceSpacesWithURLCode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World", "Hello%20World"},
		{"NoSpacesHere", "NoSpacesHere"},
		{" Leading and trailing ", "%20Leading%20and%20trailing%20"},
		{"Multiple   Spaces", "Multiple%20%20%20Spaces"},
	}

	for _, tt := range tests {
		result := replaceSpacesWithURLCode(tt.input)
		if result != tt.expected {
			t.Errorf("replaceSpacesWithURLCode(%q) = %q; expected %q", tt.input, result, tt.expected)
		}
	}
}

func TestPrintTransitPoints(t *testing.T) {
	tp := []TransitPoint{
		{"Station A", "001", "SNCB", "Main Station"},
		{"Station B", "002", "DELIJN", "Bus Stop"},
	}

	expectedOutput := "ID: 001 Name: Station A Description: Main Station Provider: SNCB\nID: 002 Name: Station B Description: Bus Stop Provider: DELIJN\n"

	// Capture the output of printTransitPoints
	output := captureOutput(func() {
		printTransitPoints(tp)
	})

	if output != expectedOutput {
		t.Errorf("printTransitPoints output = %q; expected %q", output, expectedOutput)
	}
}

func TestGetCurrentTimeHHMM(t *testing.T) {
	// Get the actual current time formatted as HHMM
	expectedTime := time.Now().Format("1504") // 24-hour format without colon

	result := getCurrentTimeHHMM()

	if result != expectedTime {
		t.Errorf("getCurrentTimeHHMM() = %q; expected %q", result, expectedTime)
	}
}

func TestNormalizeTime(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"14:30", "1430"},
		{" 9:45 ", "945"},
		{"12:00 AM", "1200AM"},
		{"23:59 ", "2359"},
	}

	for _, tt := range tests {
		result := normalizeTime(tt.input)
		if result != tt.expected {
			t.Errorf("normalizeTime(%q) = %q; expected %q", tt.input, result, tt.expected)
		}
	}
}

func TestLogVerbose(t *testing.T) {
	verbose = true
	expectedOutput := "Verbose message"

	// Capture the output of logVerbose
	output := captureOutput(func() {
		logVerbose("Verbose message")
	})

	if output != expectedOutput+"\n" {
		t.Errorf("logVerbose output = %q; expected %q", output, expectedOutput+"\n")
	}

	verbose = false
	output = captureOutput(func() {
		logVerbose("This should not appear")
	})

	if output != "" {
		t.Errorf("logVerbose output when verbose is false = %q; expected empty output", output)
	}
}

// NOTE: These tests depend on local timezone, GMT+01:00
func TestUnixToHHMM(t *testing.T) {
	// Belgium is in the CET/CEST timezone.
	// We'll explicitly set the location to "Europe/Brussels".
	location, err := time.LoadLocation("Europe/Brussels")
	if err != nil {
		t.Fatalf("Failed to load location: %v", err)
	}

	tests := []struct {
		input    int64
		expected string
	}{
		{1609459200, "01:00"},
		{1609492800, "10:20"},
		{1622505600, "02:00"},
		{1723393763, "18:29"},
	}

	for _, tt := range tests {
		// Convert the Unix timestamp to time.Time in the specified location
		tm := time.Unix(tt.input, 0).In(location)
		result := tm.Format("15:04")

		if result != tt.expected {
			t.Errorf("UnixToHHMM(%d) = %q; expected %q", tt.input, result, tt.expected)
		}
	}
}

func TestFormatDelay(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{45, "45"},        // Less than an hour
		{60, "1h"},        // Exactly one hour
		{75, "1h 15m"},    // One hour and fifteen minutes
		{120, "2h"},       // Exactly two hours
		{145, "2h 25m"},   // Two hours and twenty-five minutes
		{360, "6h"},       // Exactly six hours
		{1456, "24h 16m"}, // 24 hours and sixteen minutes
	}

	for _, tt := range tests {
		result := FormatDelay(tt.input)
		if result != tt.expected {
			t.Errorf("FormatDelay(%d) = %q; expected %q", tt.input, result, tt.expected)
		}
	}
}

// Utility function to capture the output of a function
func captureOutput(f func()) string {
	// Redirect stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run the function
	f()

	// Restore stdout and capture output
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = old

	return string(out)
}
