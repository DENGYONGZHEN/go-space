package db

import (
	"context"
	"simple-bank/util"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	arg := CreateUserParams{
		Username:       util.RandOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandOwner(),
		Email:          util.RandEmail(),
	}
	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.Time.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	user2, err := testStore.GetUser(context.Background(), user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user.HashedPassword, user2.HashedPassword)
	require.Equal(t, user.FullName, user2.FullName)
	require.Equal(t, user.Email, user2.Email)
	require.WithinDuration(t, user.PasswordChangedAt.Time, user2.PasswordChangedAt.Time, time.Second)
	require.WithinDuration(t, user.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}

func TestUpdateUserOnlyFullName(t *testing.T) {
	olderUser := createRandomUser(t)
	newFullName := util.RandOwner()
	newUsr, err := testStore.UpdateUser(context.Background(), UpdateUserParams{

		FullName: pgtype.Text{
			String: newFullName,
			Valid:  true,
		},
		Username: olderUser.Username,
	})
	require.NoError(t, err)
	require.NotEqual(t, olderUser.FullName, newUsr.FullName)
	require.Equal(t, newUsr.FullName, newFullName)
	require.Equal(t, newUsr.Email, olderUser.Email)
	require.Equal(t, newUsr.HashedPassword, olderUser.HashedPassword)
}

func TestUpdateUserOnlyEmail(t *testing.T) {
	olderUser := createRandomUser(t)
	newEmail := util.RandEmail()
	newUsr, err := testStore.UpdateUser(context.Background(), UpdateUserParams{

		Email: pgtype.Text{
			String: newEmail,
			Valid:  true,
		},
		Username: olderUser.Username,
	})
	require.NoError(t, err)
	require.NotEqual(t, olderUser.Email, newUsr.Email)
	require.Equal(t, newUsr.Email, newEmail)
	require.Equal(t, newUsr.FullName, olderUser.FullName)
	require.Equal(t, newUsr.HashedPassword, olderUser.HashedPassword)
}

func TestUpdateUserOnlyPassword(t *testing.T) {
	olderUser := createRandomUser(t)
	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)
	newUsr, err := testStore.UpdateUser(context.Background(), UpdateUserParams{

		HashedPassword: pgtype.Text{
			String: newHashedPassword,
			Valid:  true,
		},
		Username: olderUser.Username,
	})
	require.NoError(t, err)
	require.NotEqual(t, olderUser.HashedPassword, newUsr.HashedPassword)
	require.Equal(t, newUsr.HashedPassword, newHashedPassword)
	require.Equal(t, newUsr.FullName, olderUser.FullName)
	require.Equal(t, newUsr.Email, olderUser.Email)
}
