package db

const (
	NOFixedProfit = 1
	FixProfit     = 2
)

type investing interface {
}

type History struct {
	HistoryInvestings []int
}

type Counting struct {
	InvestedMoney int
	ProfitedMoney int
	History       History
	InvestingType int
}
type f struct {
}

type Person struct {
	id        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Username  string   `json:"username,omitempty"`
	Email     string   `json:"email,omitempty"`
	Counting  Counting `json:"counting,omitempty"`
	Age       int      `json:"age,omitempty"`
}
