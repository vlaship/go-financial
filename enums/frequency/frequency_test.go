package frequency

import "testing"

func TestType_Value(t *testing.T) {
	tests := []struct {
		name     string
		freq     Type
		expected int
	}{
		{"Daily frequency", DAILY, 365},
		{"Weekly frequency", WEEKLY, 52},
		{"Monthly frequency", MONTHLY, 12},
		{"Annual frequency", ANNUALLY, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.freq.Value(); got != tt.expected {
				t.Errorf("Type.Value() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestType_ValuePointer(t *testing.T) {
	freq := MONTHLY
	if got := freq.Value(); got != 12 {
		t.Errorf("Type.Value() = %v, want %v", got, 12)
	}
}
