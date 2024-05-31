package db

import (
	"context"
	"simplebank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAccountForTransfer(t *testing.T) Account {

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

func createRandomTransfer(t *testing.T) Transfer{
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	return transfer
}


func TestCreateTransfer(t *testing.T){
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T)  {
	transfer := createRandomTransfer(t)

	transfer2 , err := testQueries.GetTransfer(context.Background(), transfer.ID)
 
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer2.Amount, transfer2.Amount)
}


func TestListTransfer(t *testing.T){

	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransferParams{
		Limit: 5,
		Offset: 5,
	}

	transfer , err := testQueries.ListTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfer, 5) 
 
	for _, entry := range transfer {
		require.NotEmpty(t, entry)
	}

}
