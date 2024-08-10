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

	expectedOutput := "001 Haltenummer: Main Station Provider: SNCB\n002 Haltenummer: Bus Stop Provider: DELIJN\n"

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
