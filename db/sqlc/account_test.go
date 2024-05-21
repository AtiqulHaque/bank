package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
)


func TestCreateAccount(t *testing.T){

	arg := CreateAccountParams{
		Owner: "Atik",
		Balance: 1234,
		Currency: "USD",
	}

	account, er := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, er)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Owner, account.Owner)
}