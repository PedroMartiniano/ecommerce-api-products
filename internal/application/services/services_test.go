package services

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Setup code here if needed
	// For example: database connections, mock setup, etc.
	
	// Run the tests
	code := m.Run()
	
	// Cleanup code here if needed
	// For example: close database connections, cleanup files, etc.
	
	// Exit with the test result code
	os.Exit(code)
} 