# Assignment 5: Strategy Pattern in Banking

## 1. Class Responsibilities Table

| Class | Responsibility |
|---|---|
| `InterestStrategy` | Defines contract for loan interest calculations. |
| `StandardInterest` | Calculates fixed percentage loan interest. |
| `RiskAdjustedInterest` | Calculates interest based on credit score tiers. |
| `CollateralDiscountInterest` | Calculates interest with collateral ratio discounts. |
| `LoanCalculator` | Delegates interest calculation to selected strategy. |
| `InterestStrategySelector` | Selects appropriate interest strategy by loan type. |
| `FeeStrategy` | Defines contract for transaction fee calculations. |
| `DomesticFee` | Calculates flat plus percentage domestic transfer fees. |
| `InternationalFee` | Calculates fees using external exchange rate data. |
| `CryptoTransferFee` | Calculates fees based on simulated network congestion. |
| `FeeProcessor` | Delegates fee calculation to selected strategy. |
| `FeeStrategySelector` | Selects appropriate fee strategy by transaction type. |
| `ExchangeRateProvider` | Defines contract for retrieving currency exchange rates. |
| `LiveExchangeRateProvider` | Provides simulated live exchange rate data. |

## 2. Why Strategy and not State or Template Method?

This solution strictly implements the **Strategy Pattern** because the primary goal is to define a family of interchangeable algorithms (interest calculations, fee calculations) and select them at runtime based on external factors (loan type, transaction type). 

It is **not the State Pattern** because the context (`LoanCalculator`, `FeeProcessor`) does not transition between internal states that dictate behavior. The behavior changes because the *client* or *selector* explicitly swaps the algorithm based on business rules, not because the object's internal state changed. The strategies do not hold references to the context to change its state.

It is **not the Template Method Pattern** because we use **composition over inheritance**. Template Method would require creating a base class with an abstract skeleton and forcing subclasses to override specific steps via inheritance. Strategy encapsulates the entire algorithm in separate classes and injects them into the context via an interface, allowing runtime swapping without altering the class hierarchy.

## 3. Git Refactoring History

*   **Commit 1 (Before Refactoring):** `bad_code_before_refactor.go` containing `MonolithicLoanCalculator` with hardcoded `switch` statements.
    *   *Link:* `[INSERT_LINK_TO_FIRST_COMMIT_HERE]`
*   **Commit 2 (After Refactoring):** Deleted the monolithic class and implemented the Strategy pattern with separate interfaces, contexts, and selectors.
    *   *Link:* `[INSERT_LINK_TO_SECOND_COMMIT_HERE]`
