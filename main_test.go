package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	tests := []struct {
		name           string
		versionVar     string
		expectedOutput string
	}{
		{
			name:           "version set",
			versionVar:     "1.0.0",
			expectedOutput: "1.0.0",
		},
		{
			name:           "version empty",
			versionVar:     "",
			expectedOutput: "dev",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set version variable
			oldVersion := Version
			Version = tt.versionVar
			defer func() { Version = oldVersion }()

			// Capture output
			output := captureOutput(func() {
				versionCmd.Run(versionCmd, []string{})
			})

			if output != tt.expectedOutput {
				t.Errorf("expected %q, got %q", tt.expectedOutput, output)
			}
		})
	}
}

func TestVersionFunction(t *testing.T) {
	tests := []struct {
		name       string
		versionVar string
		expected   string
	}{
		{
			name:       "version set",
			versionVar: "1.2.3",
			expected:   "1.2.3",
		},
		{
			name:       "version empty",
			versionVar: "",
			expected:   "dev",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldVersion := Version
			Version = tt.versionVar
			defer func() { Version = oldVersion }()

			result := version()
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestRunCommand(t *testing.T) {
	tests := []struct {
		name               string
		inputJSON          string
		expectedOutputJSON string
	}{
		{
			name: "valid simple script",
			inputJSON: `{
				"script": "send [USD/2 100] (source = @foo destination = @bar)",
				"variables": {},
				"metadata": {},
				"balances": {
					"foo": {
						"USD/2": 100,
						"EUR/2": 100
					}
				},
				"featureFlags": {}
			}`,
			expectedOutputJSON: `{
				"postings": [
					{
						"source": "foo",
						"destination": "bar",
						"amount": 100,
						"asset": "USD/2"
					}
				],
				"txMeta": {},
				"accountsMeta": {}
			}`,
		},
		{
			name: "valid script with feature flag",
			inputJSON: `{
				"script": "vars { monetary $mon = [USD/2 100] \n number $n = get_amount($mon) \n } send [USD/2 $n] ( source = { oneof { @foo @bar } } destination = @dest )",
				"variables": {},
				"metadata": {},
				"balances": {
					"foo": {
						"USD/2": 100,
						"EUR/2": 100
					}
				},
				"featureFlags": {
					"experimental-oneof": {},
					"experimental-get-amount-function": {},
					"experimental-mid-script-function-call": {}
				}
			}`,
			expectedOutputJSON: `{
				"postings": [
					{
						"source": "foo",
						"destination": "dest",
						"amount": 100,
						"asset": "USD/2"
					}
				],
				"txMeta": {},
				"accountsMeta": {}
			}`,
		},
		{
			name: "empty script",
			inputJSON: `{
				"script": "",
				"variables": {},
				"metadata": {},
				"balances": {
					"foo": {
						"USD/2": 100
					}
				},
				"featureFlags": {}
			}`,
			expectedOutputJSON: `{
				"postings": [],
				"txMeta": {},
				"accountsMeta": {}
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputJSON := []byte(tt.inputJSON)

			// Create a pipe to simulate stdin
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("failed to create pipe: %v", err)
			}

			// Save original stdin
			originalStdin := os.Stdin
			defer func() { os.Stdin = originalStdin }()

			// Set our pipe as stdin
			os.Stdin = r

			// Write input to pipe
			go func() {
				defer w.Close()
				w.Write(inputJSON)
			}()

			// Capture stdout and stderr
			var stdout, stderr bytes.Buffer
			originalStdout := os.Stdout
			originalStderr := os.Stderr
			defer func() {
				os.Stdout = originalStdout
				os.Stderr = originalStderr
			}()

			// Create pipes for stdout and stderr
			rOut, wOut, _ := os.Pipe()
			rErr, wErr, _ := os.Pipe()
			os.Stdout = wOut
			os.Stderr = wErr

			// Capture output in goroutines with proper synchronization
			stdoutDone := make(chan bool)
			stderrDone := make(chan bool)

			go func() {
				io.Copy(&stdout, rOut)
				rOut.Close()
				stdoutDone <- true
			}()
			go func() {
				io.Copy(&stderr, rErr)
				rErr.Close()
				stderrDone <- true
			}()

			// Run the function in a goroutine to handle panics
			done := make(chan bool)
			go func() {
				defer func() {
					if r := recover(); r != nil {
						// Handle panic from run()
						// It is empty because the error is already captured by stderr
					}
					done <- true
				}()
				run()
			}()

			// Wait for completion
			<-done

			// Close write ends to flush
			wOut.Close()
			wErr.Close()

			// Wait for all copying to complete
			<-stdoutDone
			<-stderrDone

			// Check results
			if stderr.Len() > 0 {
				t.Errorf("unexpected error output: %s", stderr.String())
			}

			if stdout.Len() > 0 {
				// Parse actual output
				var actualResult map[string]any
				err := json.Unmarshal(stdout.Bytes(), &actualResult)
				if err != nil {
					t.Errorf("actual output is not valid JSON: %v\nOutput: %s", err, stdout.String())
					return
				}

				// Parse expected output
				var expectedResult map[string]any
				err = json.Unmarshal([]byte(tt.expectedOutputJSON), &expectedResult)
				if err != nil {
					t.Errorf("expected output is not valid JSON: %v", err)
					return
				}

				// Compare the structures
				if !compareJSONObjects(actualResult, expectedResult) {
					actualJSON, _ := json.MarshalIndent(actualResult, "", "  ")
					expectedJSON, _ := json.MarshalIndent(expectedResult, "", "  ")
					t.Errorf("Output mismatch:\nActual:\n%s\n\nExpected:\n%s", actualJSON, expectedJSON)
				}
			} else {
				t.Error("expected output, got none")
			}
		})
	}
}

func TestRunCommandWithInvalidInput(t *testing.T) {
	tests := []struct {
		name          string
		inputJSON     string
		expectedError string
	}{
		{
			name:          "invalid JSON",
			inputJSON:     "invalid json",
			expectedError: "invalid character 'i' looking for beginning of value",
		},
		{
			name:          "malformed JSON object",
			inputJSON:     `{"script": "send [USD 100]", invalid}`,
			expectedError: "invalid character 'i' looking for beginning of object key string",
		},
		{
			name: "invalid script",
			inputJSON: `{
				"script": "send [USD/2 100] (source  @foo destination = @bar)",
				"variables": {},
				"metadata": {},
				"balances": {
					"foo": {
						"USD/2": 100,
						"EUR/2": 100
					}
				},
				"featureFlags": {}
			}`,
			expectedError: "Got errors while parsing:\nmissing '=' at '@'\n  0 | send [USD/2 100] (source  @foo destination = @bar)\n    |                           \n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputJSON := []byte(tt.inputJSON)

			// Create a pipe to simulate stdin
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("failed to create pipe: %v", err)
			}

			// Save original stdin
			originalStdin := os.Stdin
			defer func() { os.Stdin = originalStdin }()

			// Set our pipe as stdin
			os.Stdin = r

			// Write input to pipe
			go func() {
				defer w.Close()
				w.Write(inputJSON)
			}()

			// Capture stdout and stderr
			var stdout, stderr bytes.Buffer
			originalStdout := os.Stdout
			originalStderr := os.Stderr
			defer func() {
				os.Stdout = originalStdout
				os.Stderr = originalStderr
			}()

			// Create pipes for stdout and stderr
			rOut, wOut, _ := os.Pipe()
			rErr, wErr, _ := os.Pipe()
			os.Stdout = wOut
			os.Stderr = wErr

			// Capture output in goroutines with proper synchronization
			stdoutDone := make(chan bool)
			stderrDone := make(chan bool)

			go func() {
				io.Copy(&stdout, rOut)
				rOut.Close()
				stdoutDone <- true
			}()
			go func() {
				io.Copy(&stderr, rErr)
				rErr.Close()
				stderrDone <- true
			}()

			// Run the function in a goroutine to handle panics
			done := make(chan bool)
			go func() {
				defer func() {
					if r := recover(); r != nil {
						// Handle panic from run()
						// It is empty because the error is already captured by stderr
					}
					done <- true
				}()
				run()
			}()

			// Wait for completion
			<-done

			// Close write ends to flush
			wOut.Close()
			wErr.Close()

			// Wait for all copying to complete
			<-stdoutDone
			<-stderrDone

			// Check results
			if stderr.Len() > 0 {
				errorOutput := stderr.String()
				if errorOutput != tt.expectedError {
					t.Errorf("Error output mismatch.\nExpected: %s\nActual: %s", tt.expectedError, errorOutput)
				}
			} else {
				t.Error("expected error output, got none")
			}
		})
	}
}

// Helper function to compare JSON objects deeply
func compareJSONObjects(actual, expected map[string]any) bool {
	if len(actual) != len(expected) {
		return false
	}

	for key, expectedValue := range expected {
		actualValue, exists := actual[key]
		if !exists {
			return false
		}

		if !compareJSONValues(actualValue, expectedValue) {
			return false
		}
	}

	return true
}

// Helper function to compare JSON values of any type
func compareJSONValues(actual, expected any) bool {
	switch expectedVal := expected.(type) {
	case map[string]any:
		actualMap, ok := actual.(map[string]any)
		if !ok {
			return false
		}
		return compareJSONObjects(actualMap, expectedVal)
	case []any:
		actualArray, ok := actual.([]any)
		if !ok {
			return false
		}
		if len(actualArray) != len(expectedVal) {
			return false
		}
		for i, expectedItem := range expectedVal {
			if !compareJSONValues(actualArray[i], expectedItem) {
				return false
			}
		}
		return true
	default:
		return actual == expected
	}
}

// Helper function to capture output from a function
func captureOutput(f func()) string {
	var stdout, stderr bytes.Buffer
	originalStdout := os.Stdout
	originalStderr := os.Stderr
	defer func() {
		os.Stdout = originalStdout
		os.Stderr = originalStderr
	}()

	rOut, wOut, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	rErr, wErr, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	os.Stdout = wOut
	os.Stderr = wErr

	stdoutDone := make(chan bool)
	stderrDone := make(chan bool)

	go func() {
		io.Copy(&stdout, rOut)
		rOut.Close()
		stdoutDone <- true
	}()
	go func() {
		io.Copy(&stderr, rErr)
		rErr.Close()
		stderrDone <- true
	}()

	f()
	wOut.Close()
	wErr.Close()

	<-stdoutDone
	<-stderrDone

	// Return combined output (stdout first, then stderr)
	combined := stdout.String() + stderr.String()
	return combined
}
