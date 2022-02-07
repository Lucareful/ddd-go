package aggregate_test

import (
	"github.com/luenci/ddd-go/aggregate"
	"testing"
)

func TestNewProduct(t *testing.T) {
	tests := []struct {
		test        string
		name        string
		description string
		price       float64
		expectedErr error
	}{
		{
			test:        "should return error if name is empty",
			name:        "",
			expectedErr: aggregate.ErrMissingValues,
		},
		{
			test:        "validvalues",
			name:        "test",
			description: "test",
			price:       1.0,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := aggregate.NewProduct(tt.name, tt.description, tt.price)
			if err != tt.expectedErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectedErr, err)
			}
		})
	}
}
