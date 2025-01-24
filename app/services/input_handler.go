package services

import (
	"bufio"
	"credit-simulator/app/models"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// InputHandler manages input processing
type InputHandler struct {
	reader *bufio.Reader
}

// NewInputHandler creates a new input handler
func NewInputHandler() *InputHandler {
	return &InputHandler{
		reader: bufio.NewReader(os.Stdin),
	}
}

// ProcessConsoleInput handles interactive console input
func (h *InputHandler) ProcessConsoleInput() (models.VehicleLoan, error) {
	fmt.Println("Credit Simulator - Interactive Input")

	// Vehicle Type Input
	vehicleType := h.promptInput("Enter Vehicle Type (Motor/Mobil): ",
		[]string{"Motor", "Mobil"}, true)

	// Vehicle Condition Input
	condition := h.promptInput("Enter Vehicle Condition (Baru/Bekas): ",
		[]string{"Baru", "Bekas"}, true)

	// Vehicle Year Input
	year, err := h.promptNumericInput("Enter Vehicle Year (4 digits): ",
		func(val int) bool { return val >= 1900 && val <= 2024 })
	if err != nil {
		return nil, err
	}

	// Total Loan Amount Input
	loanAmount, err := h.promptNumericInput("Enter Total Loan Amount (max 1 million): ",
		func(val int) bool { return val > 0 && val <= 100000000 })
	if err != nil {
		return nil, err
	}

	// Tenor Input
	tenor, err := h.promptNumericInput("Enter Loan Tenor (1-6 years): ",
		func(val int) bool { return val >= 1 && val <= 6 })
	if err != nil {
		return nil, err
	}

	// Down Payment Input
	downPayment, err := h.promptNumericInput("Enter Down Payment Amount: ",
		func(val int) bool { return val > 0 })
	if err != nil {
		return nil, err
	}

	var baseRate float64

	switch vehicleType {
	case "Mobil", "mobil":
		baseRate = 0.08
	case "Motor", "motor":
		baseRate = 0.09
	}

	// Create appropriate loan based on vehicle type
	baseLoan := models.BaseLoan{
		VehicleType:     vehicleType,
		Condition:       condition,
		Year:            year,
		TotalLoanAmount: float64(loanAmount),
		TenorYears:      tenor,
		DownPayment:     float64(downPayment),
		BaseRate:        baseRate,
	}

	var loan models.VehicleLoan
	if vehicleType == "Mobil" {
		loan = &models.CarLoan{BaseLoan: baseLoan}
	} else {
		loan = &models.MotorcycleLoan{BaseLoan: baseLoan}
	}

	return loan, nil
}

// Helper methods (need to be implemented)
func (h *InputHandler) promptInput(prompt string, validOptions []string, caseInsensitive bool) string {
	for {
		fmt.Print(prompt)
		input, _ := h.reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Validate input
		for _, option := range validOptions {
			if caseInsensitive {
				if strings.EqualFold(input, option) {
					return option
				}
			} else {
				if input == option {
					return input
				}
			}
		}

		fmt.Println("Invalid input. Please try again.")
	}
}

func (h *InputHandler) promptNumericInput(prompt string, validator func(int) bool) (int, error) {
	for {
		fmt.Print(prompt)
		input, _ := h.reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Convert to integer
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		// Validate input
		if validator(value) {
			return value, nil
		}

		fmt.Println("Invalid input. Please try again.")
	}
}

// ProcessFileInput handles file-based input
func (h *InputHandler) ProcessFileInput(filename string) (models.VehicleLoan, error) {
	// Read JSON file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var baseLoan models.BaseLoan
	err = json.Unmarshal(data, &baseLoan)
	if err != nil {
		return nil, err
	}

	// Create appropriate loan type
	var loan models.VehicleLoan
	if baseLoan.VehicleType == "Mobil" {
		loan = &models.CarLoan{BaseLoan: baseLoan}
	} else {
		loan = &models.MotorcycleLoan{BaseLoan: baseLoan}
	}

	return loan, nil
}
