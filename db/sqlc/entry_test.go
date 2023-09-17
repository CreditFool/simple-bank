package db

import (
	"context"
	"simple-bank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
  account := createRandomAccount(t)

  arg := CreateEntryParams {
    Account: account.ID,
    Amount: util.RandomMoney(),
  }

  entry, err := testQueries.CreateEntry(context.Background(), arg)
  require.NoError(t, err)
  require.NotEmpty(t, entry)
  
  require.Equal(t, arg.Account, entry.Account)
  require.Equal(t, arg.Amount, entry.Amount)

  require.NotZero(t, entry.ID)
  require.NotZero(t, entry.CreatedAt)

  return entry
}

func TestCreateEntry(t *testing.T) {
  createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
  entryA := createRandomEntry(t)
  entryB, err := testQueries.GetEntry(context.Background(), entryA.ID)
  require.NoError(t, err)
  require.NotEmpty(t, entryB)
  
  require.Equal(t, entryA.ID, entryB.ID)
  require.Equal(t, entryA.Account, entryB.Account)
  require.Equal(t, entryA.Amount, entryB.Amount)
  require.Equal(t, entryA.CreatedAt, entryB.CreatedAt)
}

func TestListEntry(t *testing.T) {
  for i := 0; i < 10; i++ {
    createRandomEntry(t)
  }
  
  arg := ListEntryParams {
    Limit: 5,
    Offset: 5,
  }
  
  entries, err := testQueries.ListEntry(context.Background(), arg)
  require.NoError(t, err)
  require.Len(t, entries, 5)
  
  for _, entry := range entries {
    require.NotEmpty(t, entry)
  }
}
