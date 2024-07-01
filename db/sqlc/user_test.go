package db

import (
	"context"
	"testing"
	"time"

	"github.com/Kiminoto1412/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	// user := createRandomUser(t)

	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	User, err := testQueries.CreateUser(context.Background(), arg)
	// User, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, User)

	require.Equal(t, arg.Username, User.Username)
	require.Equal(t, arg.HashedPassword, User.HashedPassword)
	require.Equal(t, arg.FullName, User.FullName)
	require.Equal(t, arg.Email, User.Email)

	require.NotZero(t, User.PasswordChangedAt.IsZero())
	require.NotZero(t, User.CreatedAt)

	return User
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	// user2, err := testStore.GetUser(context.Background(), user1.Username)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

// func TestUpdateUser(t *testing.T) {
// 	user1 := createRandomUser(t)

// 	arg := UpdateUserParams{
// 		ID:      user1.ID,
// 		Balance: util.RandomMoney(),
// 	}

// 	// user2, err := testStore.UpdateUser(context.Background(), arg)
// 	user2, err := testQueries.UpdateUser(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, user2)

// 	require.Equal(t, user1.ID, user2.ID)
// 	require.Equal(t, user1.Owner, user2.Owner)
// 	require.Equal(t, arg.Balance, user2.Balance)
// 	require.Equal(t, user1.Currency, user2.Currency)
// 	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
// }

// func TestDeleteUser(t *testing.T) {
// 	user1 := createRandomUser(t)
// 	// err := testStore.DeleteUser(context.Background(), user1.ID)
// 	err := testQueries.DeleteUser(context.Background(), user1.ID)
// 	require.NoError(t, err)

// 	// user2, err := testStore.GetUser(context.Background(), user1.ID)
// 	user2, err := testQueries.GetUser(context.Background(), user1.ID)
// 	require.Error(t, err)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	// require.EqualError(t, err, ErrRecordNotFound.Error())
// 	require.Empty(t, user2)
// }

// func TestListUsers(t *testing.T) {
// 	// var lastUser User
// 	for i := 0; i < 10; i++ {
// 		// lastUser = createRandomUser(t)
// 		createRandomUser(t)
// 	}

// 	arg := ListUsersParams{
// 		// Owner:  lastUser.Owner,
// 		Limit:  5,
// 		Offset: 0,
// 	}

// 	// Users, err := testStore.ListUsers(context.Background(), arg)
// 	Users, err := testQueries.ListUsers(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, Users)

// 	for _, User := range Users {
// 		require.NotEmpty(t, User)
// 		// require.Equal(t, lastUser.Owner, User.Owner)
// 	}
// }
