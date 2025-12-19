package paymentperiod

import "testing"

func TestType_Value(t *testing.T) {
	tests := []struct {
		name     string
		period   Type
		expected int64
	}{
		{"Beginning payment period", BEGINNING, 1},
		{"Ending payment period", ENDING, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.period.Value(); got != tt.expected {
				t.Errorf("Type.Value() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestType_ValueZeroValue(t *testing.T) {
	var period Type
	if got := period.Value(); got != 0 {
		t.Errorf("Type.Value() for zero value = %v, want 0", got)
	}
}
