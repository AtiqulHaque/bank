package db

import (
	"context"
	"simplebank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {

	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
 
	account, er := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, er)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Owner, account.Owner)

	return account
}

func createRandomEntry(t *testing.T) Entry{
	account := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	return entry
}


func TestCreateEntry(t *testing.T){
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T)  {
	entry := createRandomEntry(t)

	entry1 , err := testQueries.GetEnrty(context.Background(), entry.ID)
 
	require.NoError(t, err)
	require.NotEmpty(t, entry1)
	require.Equal(t, entry.ID, entry1.ID)
	require.Equal(t, entry.Amount, entry1.Amount)
}


func TestListEntry(t *testing.T){

	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	arg := ListEntryParams{
		Limit: 5,
		Offset: 5,
	}

	entries , err := testQueries.ListEntry(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5) 
 
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

}
