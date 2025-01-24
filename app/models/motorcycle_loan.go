package models

import (
	"fmt"
)

// MotorcycleLoan represents a motorcycle loan
type MotorcycleLoan struct {
	BaseLoan
}

// Validate performs motorcycle-specific loan validation
func (m *MotorcycleLoan) Validate() error {
	// Similar to car loan validation with different base rates
	if err := m.ValidateYear(); err != nil {
		return err
	}

	if err := m.ValidateTenor(); err != nil {
		return err
	}

	// Down payment validation
	downPaymentRatio := m.DownPayment / m.TotalLoanAmount

	if m.IsNewVehicle() && downPaymentRatio < 0.35 {
		return fmt.Errorf("down payment for new motorcycles must be at least 35%")
	}

	if !m.IsNewVehicle() && downPaymentRatio < 0.25 {
		return fmt.Errorf("down payment for used motorcycles must be at least 25%")
	}

	return nil
}

// CalculateInterestRate calculates motorcycle loan interest rate
func (m *MotorcycleLoan) CalculateInterestRate(year int) float64 {
	if year == 1 {
		return m.BaseRate
	}

	if year%2 == 0 {
		m.BaseRate += 0.001
	} else {
		m.BaseRate += 0.005
	}

	return m.BaseRate
}

// CalculateMonthlyInstallment calculates monthly loan installment
func (m *MotorcycleLoan) CalculateMonthlyInstallment(in int, rate float64) (float64, error) {
	if err := m.Validate(); err != nil {
		return 0, err
	}

	// Loan amount after down payment
	if in == 1 {
		m.Principal = m.TotalLoanAmount - m.DownPayment
	}

	// Calculate monthly installment using standard amortization formula
	m.Principal = m.Principal * (1 + rate)
	monthLeft := 12*m.TenorYears - (in-1)*12
	monthlyInstallment := m.Principal / float64(monthLeft)
	yearlyInstallment := 12 * monthlyInstallment

	m.Principal = m.Principal - yearlyInstallment

	return monthlyInstallment, nil
}

// GetDetails returns loan details as string
func (m *MotorcycleLoan) GetDetails() string {
	return fmt.Sprintf("Motorcycle Loan: %s, Year: %d, Amount: %.2f",
		m.Condition, m.Year, m.TotalLoanAmount)
}

func (m *MotorcycleLoan) GetTenor() int {
	return m.TenorYears
}
