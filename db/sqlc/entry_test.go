package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/paweldyl/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)

	require.NotZero(t, entry.AccountID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	acc := createRandomAccount(t)
	createRandomEntry(t, acc)
}

func TestGetEntry(t *testing.T) {
	acc := createRandomAccount(t)
	entry := createRandomEntry(t, acc)

	fmt.Println(entry.AccountID)
	fmt.Println(entry.Amount)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	fmt.Println(entry.AccountID)
	fmt.Println(entry.Amount)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry2.AccountID, entry.AccountID)
	require.Equal(t, entry2.Amount, entry.Amount)
	require.Equal(t, entry2.ID, entry.ID)

	require.NotZero(t, entry2.AccountID)
	require.NotZero(t, entry2.CreatedAt)
	require.WithinDuration(t, entry.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for range 10 {
		createRandomEntry(t, account)
	}
	entriesParams := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}
	entries, err := testQueries.ListEntries(context.Background(), entriesParams)
	require.NoError(t, err)

	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
