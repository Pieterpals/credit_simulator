package models

import (
	"fmt"
)

// CarLoan represents a car loan
type CarLoan struct {
	BaseLoan
}

// Validate performs car-specific loan validation
func (c *CarLoan) Validate() error {
	// Validate base loan properties
	if err := c.ValidateYear(); err != nil {
		return err
	}

	if err := c.ValidateTenor(); err != nil {
		return err
	}

	// Down payment validation
	downPaymentRatio := c.DownPayment / c.TotalLoanAmount

	if c.IsNewVehicle() && downPaymentRatio < 0.35 {
		return fmt.Errorf("down payment for new cars must be at least 35%")
	}

	if !c.IsNewVehicle() && downPaymentRatio < 0.25 {
		return fmt.Errorf("down payment for used cars must be at least 25%")
	}

	return nil
}

// The CalculateInterestRate method remains largely the same
func (c *CarLoan) CalculateInterestRate(in int) float64 {
	if in == 1 {
		return c.BaseRate
	}

	if in % 2  == 0 {
		c.BaseRate += 0.001
	} else {
		c.BaseRate += 0.005
	}

	return c.BaseRate
}

// CalculateMonthlyInstallment method can stay the same
func (c *CarLoan) CalculateMonthlyInstallment(in int, rate float64) (float64, error) {
	if err := c.Validate(); err != nil {
		return 0, err
	}

	// Loan amount after down payment
	if in == 1 {
		c.Principal = c.TotalLoanAmount - c.DownPayment
	} 

	// Calculate monthly installment using standard amortization formula
	c.Principal = c.Principal * (1 + rate)
	monthLeft := 12*c.TenorYears - (in-1)*12
	monthlyInstallment := c.Principal / float64(monthLeft)
	yearlyInstallment := 12 * monthlyInstallment

	c.Principal = c.Principal - yearlyInstallment

	return monthlyInstallment, nil
}

// GetDetails returns loan details as string
func (c *CarLoan) GetDetails() string {
	return fmt.Sprintf("Car Loan: %s, Year: %d, Amount: %.2f",
		c.Condition, c.Year, c.TotalLoanAmount)
}

func (c *CarLoan) GetTenor() int {
	return c.TenorYears
}
