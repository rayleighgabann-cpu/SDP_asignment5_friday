package main

import "fmt"

func main() {
	fmt.Println("=== BANKING STRATEGY PATTERN DEMO ===\n")

	// --- Family 1: Loan Interest ---
	fmt.Println("--- 1. Loan Interest Calculation ---")
	loanSelector := &InterestStrategySelector{}

	// Manual choice
	calc := NewLoanCalculator(&StandardInterest{})
	fmt.Printf("Manual (Standard): $%.2f\n", calc.Execute(10000, 700, 0))

	// Automatic choice via Selector
	strategy := loanSelector.SelectStrategy("risk")
	calc.SetStrategy(strategy)
	fmt.Printf("Auto (Risk, Score 500): $%.2f\n", calc.Execute(10000, 500, 0))

	// Runtime switching
	calc.SetStrategy(loanSelector.SelectStrategy("collateral"))
	fmt.Printf("Runtime Switch (Collateral): $%.2f\n\n", calc.Execute(10000, 700, 8000))

	// --- Family 2: Transaction Fees ---
	fmt.Println("--- 2. Transaction Fee Calculation ---")
	feeSelector := &FeeStrategySelector{}

	// Using real (simulated) external data source
	realProvider := &LiveExchangeRateProvider{}
	feeProc := NewFeeProcessor(feeSelector.SelectStrategy("domestic", realProvider))
	fmt.Printf("Domestic Fee: $%.2f\n", feeProc.Execute(1000, "USD"))

	feeProc.SetStrategy(feeSelector.SelectStrategy("international", realProvider))
	fmt.Printf("International Fee (EUR->USD): $%.2f\n", feeProc.Execute(1000, "EUR"))

	feeProc.SetStrategy(feeSelector.SelectStrategy("crypto", realProvider))
	fmt.Printf("Crypto Fee (Large): $%.2f\n", feeProc.Execute(60000, "BTC"))
}
