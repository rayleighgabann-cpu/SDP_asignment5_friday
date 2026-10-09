package main

import "testing"

// In Go, table-driven tests are the idiomatic equivalent of JUnit's @ParameterizedTest
func TestLoanStrategies_Parameterized(t *testing.T) {
	tests := []struct {
		name        string
		strategy    InterestStrategy
		principal   float64
		creditScore int
		collateral  float64
		wantMin     float64
		wantMax     float64
	}{
		{"Standard", &StandardInterest{}, 10000, 700, 0, 400, 600},
		{"Risk High", &RiskAdjustedInterest{}, 10000, 500, 0, 800, 1000},
		{"Risk Low", &RiskAdjustedInterest{}, 10000, 800, 0, 300, 500},
		{"Collateral High", &CollateralDiscountInterest{}, 10000, 700, 8000, 100, 300},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.strategy.CalculateInterest(tt.principal, tt.creditScore, tt.collateral)
			if got < tt.wantMin || got > tt.wantMax {
				t.Errorf("CalculateInterest() = %v, want between %v and %v", got, tt.wantMin, tt.wantMax)
			}
		})
	}
}

func TestInterestSelector(t *testing.T) {
	selector := &InterestStrategySelector{}

	if _, ok := selector.SelectStrategy("standard").(*StandardInterest); !ok {
		t.Error("Selector failed for standard")
	}
	if _, ok := selector.SelectStrategy("risk").(*RiskAdjustedInterest); !ok {
		t.Error("Selector failed for risk")
	}
	if _, ok := selector.SelectStrategy("collateral").(*CollateralDiscountInterest); !ok {
		t.Error("Selector failed for collateral")
	}
}

func TestLoanContext(t *testing.T) {
	ctx := NewLoanCalculator(&StandardInterest{})
	res1 := ctx.Execute(10000, 700, 0)

	ctx.SetStrategy(&RiskAdjustedInterest{})
	res2 := ctx.Execute(10000, 500, 0)

	if res1 == res2 {
		t.Error("Context failed to switch strategies at runtime")
	}
}
