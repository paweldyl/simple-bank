package db

import (
	"context"
	"testing"
	"time"

	"github.com/paweldyl/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, from Account, to Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: from.ID,
		ToAccountID:   to.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, arg.FromAccountID)
	require.Equal(t, transfer.ToAccountID, arg.ToAccountID)
	require.Equal(t, transfer.Amount, arg.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	from := createRandomAccount(t)
	to := createRandomAccount(t)
	createRandomTransfer(t, from, to)
}

func TestGetTransfer(t *testing.T) {
	from := createRandomAccount(t)
	to := createRandomAccount(t)
	transfer := createRandomTransfer(t, from, to)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer2.ID, transfer.ID)
	require.Equal(t, transfer2.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transfer2.ToAccountID, transfer.ToAccountID)
	require.Equal(t, transfer2.Amount, transfer.Amount)

	require.NotZero(t, transfer2.CreatedAt)
	require.WithinDuration(t, transfer.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	from := createRandomAccount(t)
	to := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, from, to)
	}

	arg := ListTransfersParams{
		FromAccountID: from.ID,
		ToAccountID:   to.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
