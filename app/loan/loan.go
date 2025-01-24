package loan

import (
	"fmt"
	"time"
)

// VehicleLoan interface defines common loan methods
type VehicleLoan interface {
	Validate() error
	CalculateMonthlyInstallment() (float64, error)
	CalculateInterestRate() float64
	GetDetails() string
}

// BaseLoan contains common loan properties
type BaseLoan struct {
	VehicleType     string // Motor/Mobil
	Condition       string // Baru/Bekas
	Year            int
	TotalLoanAmount float64
	TenorYears      int
	DownPayment     float64
}

// IsNewVehicle checks if the vehicle is new
func (b *BaseLoan) IsNewVehicle() bool {
	return b.Condition == "Baru"
}

// ValidateYear checks vehicle year validity
func (b *BaseLoan) ValidateYear() error {
	currentYear := time.Now().Year()

	if b.IsNewVehicle() && b.Year < currentYear-1 {
		return fmt.Errorf("new vehicles cannot be older than current year")
	}

	return nil
}

// ValidateTenor checks loan tenor
func (b *BaseLoan) ValidateTenor() error {
	if b.TenorYears < 1 || b.TenorYears > 6 {
		return fmt.Errorf("loan tenor must be between 1 and 6 years")
	}
	return nil
}
