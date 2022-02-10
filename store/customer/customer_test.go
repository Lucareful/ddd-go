package customer

import (
	"github.com/google/uuid"
	"github.com/luenci/ddd-go/domain/customer"
	"testing"
)

func TestMemoryRepository_Add(t *testing.T) {
	tests := []struct {
		name        string
		cust        string
		expectedErr error
	}{
		{
			name:        "Add Customer",
			cust:        "Percy",
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := MemoryRepository{
				customers: map[uuid.UUID]customer.Customer{},
			}

			cust, err := customer.NewCustomer(tt.cust)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(cust)
			if err != tt.expectedErr {
				t.Errorf("Expected error %v, got %v", tt.expectedErr, err)
			}

			found, err := repo.Get(cust.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != cust.GetID() {
				t.Errorf("Expected %v, got %v", cust.GetID(), found.GetID())
			}
		})
	}
}

func TestMemoryRepository_Get(t *testing.T) {
	// Create a fake customer to add to repository
	cust, err := customer.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}
	id := cust.GetID()
	// Create the repo to use, and add some test Data to it for testing
	// Skip Factory for this
	repo := MemoryRepository{
		customers: map[uuid.UUID]customer.Customer{
			id: cust,
		},
	}

	tests := []struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}{
		{
			name:        "No Customer By ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		}, {
			name:        "Customer By ID",
			id:          id,
			expectedErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repo.Get(tt.id)
			if err != tt.expectedErr {
				t.Errorf("Expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}
