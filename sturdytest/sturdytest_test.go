package sturdytest_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	s "github.com/xplosunn/sturdy/shortsyntax"
	"github.com/xplosunn/sturdy/sturdytest"
	"testing"
)

func TestSturdyTestPropertyAfterCall(t *testing.T) {
	moneyOnAliceAccount := 100
	moneyOnBobAccount := 0
	totalMoneyOnBankPropertyFailed := false
	var totalMoneyOnBankWhenPropertyFailed int
	sturdytest.SturdyTestProperty(t, sturdytest.ErrorConstAfterCall(), func(t *testing.T) {
		totalMoneyOnBank := moneyOnAliceAccount + moneyOnBobAccount
		if totalMoneyOnBank != 100 {
			totalMoneyOnBankPropertyFailed = true
			totalMoneyOnBankWhenPropertyFailed = totalMoneyOnBank
		}
	}, func(ctx context.Context) error {
		if totalMoneyOnBankPropertyFailed {
			return nil
		}
		moneyWithdrawn, err := s.S1(ctx, func() (int, error) {
			money := moneyOnAliceAccount
			moneyOnAliceAccount = 0
			return money, nil
		})
		if err != nil {
			return err
		}
		panic("reached unexpected code point")
		// not reachable but here for sake of the finished example
		err = s.S0(ctx, func() error {
			moneyOnBobAccount += moneyWithdrawn
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
	assert.True(t, totalMoneyOnBankPropertyFailed)
	assert.Equal(t, 0, totalMoneyOnBankWhenPropertyFailed)
}

func TestSturdyTestEventualSuccess(t *testing.T) {
	var err error
	errCount := 0
	sturdytest.SturdyTest(t, sturdytest.ErrorIncreaseAfterCall(), func(ctx context.Context) error {
		err = s.S0(ctx, func() error {
			return nil
		})
		if err != nil {
			errCount += 1
			return err
		}
		err = s.S0(ctx, func() error {
			return nil
		})
		if err != nil {
			errCount += 1
			return err
		}
		return nil
	})
	assert.NoError(t, err)
	assert.Equal(t, 2, errCount)
}
