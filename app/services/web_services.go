package services

// LoanWebService handles web service interactions
type LoanWebService struct {
	BaseURL string
}

// LoanCalculation represents existing loan calculations
type LoanCalculation struct {
	VehicleType        string  `json:"vehicle_type"`
	MonthlyInstallment float64 `json:"monthly_installment"`
	InterestRate       float64 `json:"interest"`
}
