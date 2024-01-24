package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/luanle13/simple-bank/util"
	"github.com/stretchr/testify/assert"
)

func createRandomTransfer(t *testing.T) Transfer {
	fromAccount := createRandomAccount(t)
	toAccount :=createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID: toAccount.ID,
		Amount: util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, transfer)
	assert.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	assert.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	assert.Equal(t, arg.Amount, transfer.Amount)
	assert.NotZero(t, transfer.ID)
	assert.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, transfer2)
	assert.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	assert.Equal(t, transfer1.Amount, transfer2.Amount)
	assert.Equal(t, transfer1.ID, transfer2.ID)
	assert.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	assert.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	arg := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: util.RandomMoney(),
	}
	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, transfer2)
	assert.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	assert.Equal(t, arg.ID, transfer2.ID)
	assert.Equal(t, arg.Amount, transfer2.Amount)
	assert.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	assert.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	assert.NoError(t, err)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	assert.Error(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, transfer2)
}