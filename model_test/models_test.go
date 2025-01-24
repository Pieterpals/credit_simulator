package models_test

import (
	"credit-simulator/app/models"
	"testing"
)

func TestCarLoanValidation(t *testing.T) {
	testCases := []struct {
		name        string
		loan        models.CarLoan
		expectedErr bool
	}{
		{
			name: "Valid New Car Loan",
			loan: models.CarLoan{
				BaseLoan: models.BaseLoan{
					VehicleType:     "Mobil",
					Condition:       "Baru",
					Year:            2024,
					TotalLoanAmount: 100000,
					TenorYears:      3,
					DownPayment:     40000,
				},
			},
			expectedErr: false,
		},
		{
			name: "Invalid Down Payment",
			loan: models.CarLoan{
				BaseLoan: models.BaseLoan{
					VehicleType:     "Mobil",
					Condition:       "Baru",
					Year:            2024,
					TotalLoanAmount: 100000,
					TenorYears:      3,
					DownPayment:     20000,
				},
			},
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.loan.Validate()
			if (err != nil) != tc.expectedErr {
				t.Errorf("Validation error mismatch: got %v, want %v", err, tc.expectedErr)
			}
		})
	}
}

func TestMonthlyInstallmentCalculation(t *testing.T) {
	loan := models.CarLoan{
		BaseLoan: models.BaseLoan{
			TotalLoanAmount: 100000,
			TenorYears:      3,
			DownPayment:     35000,
			BaseRate:        0.08,
		},
	}

	installment, err := loan.CalculateMonthlyInstallment(1, 0.08)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if installment <= 0 {
		t.Error("Monthly installment should be positive")
	}
}
