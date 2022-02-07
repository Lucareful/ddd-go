package aggregate

import (
	"testing"
)

func TestNewCustomer(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name        string
		test        string
		expectedErr error
	}{
		{
			test:        "Empty Name validation",
			name:        "",
			expectedErr: ErrInvalidPerson,
		}, {
			test:        "Valid Name",
			name:        "Percy Bolmer",
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new customer
			_, err := NewCustomer(tt.name)
			// Check if the error matches the expected error
			if err != tt.expectedErr {
				t.Errorf("Expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}
