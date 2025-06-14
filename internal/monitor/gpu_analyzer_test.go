package monitor

import (
	"testing"
)

func TestNewGPUAnalyzer(t *testing.T) {
	// Create GPU analyzer with nil clientset for basic structure test
	analyzer := NewGPUAnalyzer(nil)

	// Test that analyzer is created correctly
	if analyzer == nil {
		t.Error("Expected GPUAnalyzer to be created, got nil")
	}
}

func TestContains(t *testing.T) {
	testCases := []struct {
		slice    []string
		item     string
		expected bool
	}{
		{[]string{"a", "b", "c"}, "b", true},
		{[]string{"a", "b", "c"}, "d", false},
		{[]string{}, "a", false},
		{[]string{"test"}, "test", true},
	}

	for _, tc := range testCases {
		result := contains(tc.slice, tc.item)
		if result != tc.expected {
			t.Errorf("contains(%v, %s) = %v, expected %v", tc.slice, tc.item, result, tc.expected)
		}
	}
}
