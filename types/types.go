package types

type Summary struct {
	Balance struct {
		Starting string
		Ending   string
	}
	Deposits    string
	Checks      string
	Withdrawals struct {
		Debit      string
		Electronic string
	}
}

type Deposit struct {
	Date        string
	Amount      string
	Description string
}

type Withdrawal struct {
	Date        string
	Amount      string
	Description string
}

type Check struct {
	ID     string
	Date   string
	Amount string
}

type BankJson struct {
	CheckingSummary Summary
	Deposits        []Deposit
	Withdrawals     struct {
		Debit      []Withdrawal
		Electronic []Withdrawal
	}
	Checks []Check
}
