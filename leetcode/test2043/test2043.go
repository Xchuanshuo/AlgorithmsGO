package test2043

type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{balance: balance}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.isValid(account1) || !this.isValid(account2) ||
		this.balance[account1-1] < money {
		return false
	}
	this.balance[account1-1] -= money
	this.balance[account2-1] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.isValid(account) { return false }
	this.balance[account-1] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.isValid(account) || this.balance[account-1] < money {
		return false
	}
	this.balance[account-1] -= money
	return true
}

func (this *Bank) isValid(account int) bool {
	if account >= 1 && account <= len(this.balance) {
		return true
	}
	return false
}