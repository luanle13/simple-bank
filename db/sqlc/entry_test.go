package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/luanle13/simple-bank/util"
	"github.com/stretchr/testify/assert"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, entry)
	assert.Equal(t, arg.AccountID, entry.AccountID)
	assert.Equal(t, arg.Amount, entry.Amount)
	assert.NotZero(t, entry.ID)
	assert.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, entry2)
	assert.Equal(t, entry1.AccountID, entry2.AccountID)
	assert.Equal(t, entry1.Amount, entry2.Amount)
	assert.Equal(t, entry1.ID, entry2.ID)
	assert.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestUpdateEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	arg := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: util.RandomMoney(),
	}
	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	assert.NoError(t, err)
	assert.NotEmpty(t, entry2)
	assert.Equal(t, entry1.AccountID, entry2.AccountID)
	assert.Equal(t, arg.ID, entry2.ID)
	assert.Equal(t, arg.Amount, entry2.Amount)
	assert.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestDeleteEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	assert.NoError(t, err)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	assert.Error(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, entry2)
}