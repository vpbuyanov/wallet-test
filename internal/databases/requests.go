package databases

const (
	FindWallet        = `select number, code, amount from public."Wallet" where number = $1`
	FirstDeposit      = `insert into public."Wallet" (number, code, amount) values ($1, $2, $3)`
	Deposit           = `update public."Wallet" set amount = amount + $3 where number = $1 and code = $2`
	Withdraw          = `update public."Wallet" set amount = amount - $3 where number = $1 and code = $2`
	CheckStatusWallet = `select number, status from public."Wallet" where number = $1`
)
