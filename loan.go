package main

// --- Strategy Interface ---
type InterestStrategy interface {
	CalculateInterest(principal float64, creditScore int, collateralValue float64) float64
}

// --- Concrete Strategies ---
type StandardInterest struct{}

func (s *StandardInterest) CalculateInterest(principal float64, creditScore int, collateralValue float64) float64 {
	return principal * 0.05 // Fixed 5%
}

type RiskAdjustedInterest struct{}

func (s *RiskAdjustedInterest) CalculateInterest(principal float64, creditScore int, collateralValue float64) float64 {
	rate := 0.05
	if creditScore < 600 {
		rate += 0.04 // High risk penalty
	} else if creditScore > 750 {
		rate -= 0.01 // Good credit discount
	}
	return principal * rate
}

type CollateralDiscountInterest struct{}

func (s *CollateralDiscountInterest) CalculateInterest(principal float64, creditScore int, collateralValue float64) float64 {
	baseRate := 0.05
	ratio := collateralValue / principal
	discount := ratio * 0.02 // Discount scales with collateral ratio
	if discount > 0.03 {
		discount = 0.03 // Max discount cap
	}
	return principal * (baseRate - discount)
}

// --- Context ---
type LoanCalculator struct {
	strategy InterestStrategy
}

func NewLoanCalculator(strategy InterestStrategy) *LoanCalculator {
	return &LoanCalculator{strategy: strategy}
}

func (c *LoanCalculator) SetStrategy(strategy InterestStrategy) {
	c.strategy = strategy
}

func (c *LoanCalculator) Execute(principal float64, creditScore int, collateralValue float64) float64 {
	return c.strategy.CalculateInterest(principal, creditScore, collateralValue)
}

// --- Selector ---
type InterestStrategySelector struct{}

func (s *InterestStrategySelector) SelectStrategy(loanType string) InterestStrategy {
	switch loanType {
	case "standard":
		return &StandardInterest{}
	case "risk":
		return &RiskAdjustedInterest{}
	case "collateral":
		return &CollateralDiscountInterest{}
	default:
		return &StandardInterest{}
	}
}
