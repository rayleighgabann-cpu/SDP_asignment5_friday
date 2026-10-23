package main

import "testing"

// Mock for external data source (Clean Code requirement)
type MockExchangeRateProvider struct {
	Rate float64
}

func (m *MockExchangeRateProvider) GetRate(from, to string) float64 {
	return m.Rate
}

func TestFeeStrategies_Parameterized(t *testing.T) {
	mockProvider := &MockExchangeRateProvider{Rate: 1.0} // 1:1 for easy math

	tests := []struct {
		name     string
		strategy FeeStrategy
		amount   float64
		currency string
		want     float64
	}{
		{"Domestic", &DomesticFee{}, 1000, "USD", 6.0},                                      // 5 + 1
		{"International", &InternationalFee{rateProvider: mockProvider}, 1000, "EUR", 20.0}, // 1000 * 0.02
		{"Crypto Small", &CryptoTransferFee{}, 1000, "BTC", 10.0},
		{"Crypto Large", &CryptoTransferFee{}, 60000, "BTC", 35.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.strategy.CalculateFee(tt.amount, tt.currency)
			if got != tt.want {
				t.Errorf("CalculateFee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFeeSelector(t *testing.T) {
	selector := &FeeStrategySelector{}
	mockProvider := &MockExchangeRateProvider{Rate: 1.0}

	if _, ok := selector.SelectStrategy("domestic", mockProvider).(*DomesticFee); !ok {
		t.Error("Selector failed for domestic")
	}
	if _, ok := selector.SelectStrategy("international", mockProvider).(*InternationalFee); !ok {
		t.Error("Selector failed for international")
	}
}
