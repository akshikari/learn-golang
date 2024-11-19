package arraysandslicesgenerics

import "testing"

func TestBadBank(t *testing.T) {
	transactions := []Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}

	AssertEqual(t, BalanceFor(transactions, "Riya"), 100.0)
	AssertEqual(t, BalanceFor(transactions, "Chris"), -75.0)
	AssertEqual(t, BalanceFor(transactions, "Adil"), -25.0)

}
