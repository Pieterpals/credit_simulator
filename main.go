package main

import (
	"fmt"
	"os"

	"credit-simulator/app/models"
	"credit-simulator/app/services"
)

func main() {
	// Create input handler
	var loan models.VehicleLoan
	// var err error

	inputHandler := services.NewInputHandler()

	loan, _ = inputHandler.ProcessConsoleInput()

	tenorYears := loan.GetTenor()

	// Display results
	for year := 1; year <= tenorYears; year++ {
		interestRate := loan.CalculateInterestRate(year)

		// Simulate changing interest rate for different years
		monthlyInstallment, err := loan.CalculateMonthlyInstallment(year, interestRate)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Printf("tahun %d : Rp. %.2f/bln , Suku Bunga : %.1f%%\n", year, monthlyInstallment, interestRate*100)
	}

	// Explicit pause
	fmt.Println("Press Enter to exit...")
	fmt.Scanln() // Waits for the user to press Enter
}
