package db

import (
	"context"
	"simple-bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	accountA := createRandomAccount(t)
	accountB := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccount: accountA.ID,
		ToAccount:   accountB.ID,
		Amount:      util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccount, transfer.FromAccount)
	require.Equal(t, arg.ToAccount, transfer.ToAccount)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotEmpty(t, transfer.ID)
	require.NotEmpty(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transferA := createRandomTransfer(t)
	transferB, err := testQueries.GetTransfer(context.Background(), transferA.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transferB)

	require.Equal(t, transferA.ID, transferB.ID)
	require.Equal(t, transferA.FromAccount, transferB.FromAccount)
	require.Equal(t, transferA.ToAccount, transferB.ToAccount)
	require.Equal(t, transferA.Amount, transferB.Amount)
	require.WithinDuration(t, transferA.CreatedAt, transferB.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransferParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
