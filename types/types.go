package types

type Summary struct {
	Beginning   string
	Ending      string
	Deposits    string
	Withdrawals string
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
	Summary     Summary
	Deposits    []Deposit
	Withdrawals []Withdrawal
	Checks      []Check
}

type PurchaseType string

type TestJson struct {
	ID     string
	Data   []int
	Title  string
	Labels []string
}

type GraphJson struct {
	ID     string
	Data   []float64
	Title  string
	Labels []string
}
