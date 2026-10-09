package main

// --- External Dependency Interface (Clean Code Ch. 10) ---
type ExchangeRateProvider interface {
	GetRate(fromCurrency, toCurrency string) float64
}

type LiveExchangeRateProvider struct{}

func (p *LiveExchangeRateProvider) GetRate(fromCurrency, toCurrency string) float64 {
	// Simulated live API call
	return 1.10 // e.g., EUR to USD
}

// --- Strategy Interface ---
type FeeStrategy interface {
	CalculateFee(amount float64, currency string) float64
}

// --- Concrete Strategies ---
type DomesticFee struct{}

func (s *DomesticFee) CalculateFee(amount float64, currency string) float64 {
	return 5.0 + (amount * 0.001) // Flat $5 + 0.1%
}

type InternationalFee struct {
	rateProvider ExchangeRateProvider // Dependency Injection
}

func (s *InternationalFee) CalculateFee(amount float64, currency string) float64 {
	rate := s.rateProvider.GetRate(currency, "USD")
	convertedAmount := amount * rate
	return convertedAmount * 0.02 // 2% of converted amount
}

type CryptoTransferFee struct{}

func (s *CryptoTransferFee) CalculateFee(amount float64, currency string) float64 {
	// Simulated network congestion logic
	baseFee := 10.0
	if amount > 50000 {
		baseFee += 25.0 // High priority fee for large amounts
	}
	return baseFee
}

// --- Context ---
type FeeProcessor struct {
	strategy FeeStrategy
}

func NewFeeProcessor(strategy FeeStrategy) *FeeProcessor {
	return &FeeProcessor{strategy: strategy}
}

func (c *FeeProcessor) SetStrategy(strategy FeeStrategy) {
	c.strategy = strategy
}

func (c *FeeProcessor) Execute(amount float64, currency string) float64 {
	return c.strategy.CalculateFee(amount, currency)
}

// --- Selector ---
type FeeStrategySelector struct{}

func (s *FeeStrategySelector) SelectStrategy(txType string, provider ExchangeRateProvider) FeeStrategy {
	switch txType {
	case "domestic":
		return &DomesticFee{}
	case "international":
		return &InternationalFee{rateProvider: provider}
	case "crypto":
		return &CryptoTransferFee{}
	default:
		return &DomesticFee{}
	}
}
