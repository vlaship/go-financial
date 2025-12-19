package interesttype

import "testing"

func TestType_String(t *testing.T) {
	tests := []struct {
		name     string
		itype    Type
		expected string
	}{
		{"Flat interest type", FLAT, "flat"},
		{"Reducing interest type", REDUCING, "reducing"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.itype.String(); got != tt.expected {
				t.Errorf("Type.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestType_StringZeroValue(t *testing.T) {
	var itype Type
	if got := itype.String(); got != "" {
		t.Errorf("Type.String() for zero value = %v, want empty string", got)
	}
}
