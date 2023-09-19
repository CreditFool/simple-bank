package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	accountA := createRandomAccount(t)
	accountB := createRandomAccount(t)
	fmt.Println(">> Before: ", accountA.Balance, accountB.Balance)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: accountA.ID,
				ToAccountID:   accountB.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, accountA.ID, transfer.FromAccount)
		require.Equal(t, accountB.ID, transfer.ToAccount)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, accountA.ID, fromEntry.Account)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, accountB.ID, toEntry.Account)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, accountA.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, accountB.ID, toAccount.ID)

		fmt.Println(">> Tx: ", fromAccount.Balance, toAccount.Balance)
		diffA := accountA.Balance - fromAccount.Balance
		diffB := toAccount.Balance - accountB.Balance
		require.Equal(t, diffA, diffB)
		require.True(t, diffA > 0)
		require.True(t, diffA%amount == 0)

		k := int(diffA / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	updatedAccountA, err := testQueries.GetAccount(context.Background(), accountA.ID)
	require.NoError(t, err)

	updatedAccountB, err := testQueries.GetAccount(context.Background(), accountB.ID)
	require.NoError(t, err)

	fmt.Println(">> After: ", updatedAccountA.Balance, updatedAccountB.Balance)
	require.Equal(t, accountA.Balance-(int64(n)*amount), updatedAccountA.Balance)
	require.Equal(t, accountB.Balance+(int64(n)*amount), updatedAccountB.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	accountA := createRandomAccount(t)
	accountB := createRandomAccount(t)
	fmt.Println(">> Before: ", accountA.Balance, accountB.Balance)

	n := 10
	amount := int64(10)
	errs := make(chan error)

	for i := 0; i < n; i++ {

		fromAccount := accountA.ID
		toAccount := accountB.ID

		if i%2 == 1 {
			fromAccount = accountB.ID
			toAccount = accountA.ID
		}

		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccount,
				ToAccountID:   toAccount,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	updatedAccountA, err := testQueries.GetAccount(context.Background(), accountA.ID)
	require.NoError(t, err)

	updatedAccountB, err := testQueries.GetAccount(context.Background(), accountB.ID)
	require.NoError(t, err)

	fmt.Println(">> After: ", updatedAccountA.Balance, updatedAccountB.Balance)
	require.Equal(t, accountA.Balance, updatedAccountA.Balance)
	require.Equal(t, accountB.Balance, updatedAccountB.Balance)
}
