package gocet

import (
	"math"
)

func CalculateCET(totalLoanAmount, interestRate, administrativeFee float64, numberOfInstallments int) float64 {
	effectiveRate := math.Pow(1+interestRate/100, 1/12.0) - 1
	totalAdministrativeFee := administrativeFee * float64(numberOfInstallments)
	totalPayment := totalLoanAmount + totalAdministrativeFee

	annuityFactor := (effectiveRate * math.Pow(1+effectiveRate, float64(numberOfInstallments))) / (math.Pow(1+effectiveRate, float64(numberOfInstallments)) - 1)
	installmentAmount := totalPayment * annuityFactor
	monthlyInterest := (installmentAmount - totalLoanAmount/float64(numberOfInstallments)) * float64(numberOfInstallments) / totalLoanAmount * 100

	cet := ((monthlyInterest / 100) + (administrativeFee / totalLoanAmount)) * 100

	return math.Round(cet*100) / 100
}
