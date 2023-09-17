package db

import (
	"simple-bank/util"
	"testing"
)

func TestCreateTransfer(t *testing.T) {
  accountA := createRandomAccount(t)
  accountB := createRandomAccount(t)

  arg := CreateTransferParams {
    FromAccount: accountA.ID,
    ToAccount: accountB.ID,
    Amount: util.RandomMoney(),
  }
}
