package bank

import (
	"fmt"

	"github.com/cucumber/godog"
)

var testAccount *account

func iHaveABankAccountWith(balance int) error {
	testAccount = &account{balance: balance}
	return nil
}

func iDeposit(amount int) error {
	testAccount.Deposit(amount)
	return nil
}

func iWithdraw(amount int) error {
	testAccount.Withdraw(amount)
	return nil
}

func itShouldHaveABalanceOf(balance int) error {
	if testAccount.balance == balance {
		return nil
	}
	return fmt.Errorf("Incorrect account balance, I wanted %d, I got %d", testAccount.balance, balance)
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	/*
		//With the new version of GoDog, it looks the state cleaning is not needed anymore
		func InitializeScenario(sc *godog.ScenarioContext) {
		sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
			testAccount = nil // clean the state before every scenario
			return ctx, nil
		})
		ctx.Step(`^I deposit (\d+)\$$`, iDeposit)
	*/

	ctx.Step(`^I deposit (\d+)\$$`, iDeposit)
	//ctx.Step(`^I have a bank account with (\d+)\$$`, iHaveABankAccountWith)
	ctx.Step(`^I withdraw (\d+)\$$`, iWithdraw)
	//ctx.Step(`^it should have a balance of (\d+)\$$`, itShouldHaveABalanceOf)

	// Add negative tests, thx to https://github.com/cucumber/godog/issues/257
	ctx.Step(`^I have a bank account with (\-*\d+)\$$`, iHaveABankAccountWith)
	ctx.Step(`^it should have a balance of (\-*\d+)\$$`, itShouldHaveABalanceOf)
}
