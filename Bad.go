package main

import "fmt"

// BAD: Monolithic class with switch statements. Violates Open/Closed Principle.
type MonolithicLoanCalculator struct{}

func (c *MonolithicLoanCalculator) CalculateInterest(principal float64, loanType string, creditScore int, collateral float64) float64 {
	switch loanType {
	case "standard":
		return principal * 0.05
	case "risk":
		rate := 0.05
		if creditScore < 600 {
			rate += 0.03
		}
		return principal * rate
	case "collateral":
		discount := 0.0
		if collateral > principal*0.5 {
			discount = 0.02
		}
		return principal * (0.05 - discount)
	default:
		return 0
	}
}

func main() {
	calc := &MonolithicLoanCalculator{}
	fmt.Println(calc.CalculateInterest(10000, "standard", 700, 0))
}
