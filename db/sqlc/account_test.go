package db

import (
	"context"
	"simplebank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func getRandomAccount(t *testing.T) Account {

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


func TestCreateAccount(t *testing.T){
	getRandomAccount(t)
}

func TestGetAccount(t *testing.T)  {
	account1 :=getRandomAccount(t)

	account2 , err := testQueries.GetAccount(context.Background(), account1.ID)
 
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
}


func TestUpdateAccount(t *testing.T)  {
	account1 :=getRandomAccount(t)

	arg := UpdateAccountParams{
		ID : account1.ID,
		Balance: util.RandomMoney(),
	}

	account2 , err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err) 
	require.NotEmpty(t, account2)
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
}


func TestDeleteAccount(t *testing.T){

	account1 :=getRandomAccount(t)
	err  := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	acc, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, acc)
}

func TestListAccount(t *testing.T){

	for i := 0; i < 10; i++ {
		getRandomAccount(t)
	}

	arg := ListAccountParams{
		Limit: 5,
		Offset: 5,
	}

	accounts , err := testQueries.ListAccount(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5) 
 
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}