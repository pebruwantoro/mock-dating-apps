package repository

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	request  interface{}
	response interface{}
	mockFunc func(m sqlmock.Sqlmock)
	err      error
}

func TestCreateUser(t *testing.T) {
	cases := []testCase{
		{
			name: "success",
			request: User{
				UUID:      "uuid",
				Email:     "email",
				Username:  "username",
				Password:  "password",
				IsPremium: true,
			},
			response: User{
				UUID: "uuid",
			},
			mockFunc: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(`
				INSERT INTO users (uuid, email, username, password, is_premium, created_at)
					VALUES ($1, $2, $3, $4, $5, NOW())
					returning uuid;
				`)).WithArgs("uuid", "email", "username", "password", true).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("uuid"))
			},
			err: nil,
		},
		{
			name: "error",
			request: User{
				UUID:      "uuid",
				Email:     "email",
				Username:  "username",
				Password:  "password",
				IsPremium: true,
			},
			response: User{},
			mockFunc: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(`
				INSERT INTO users (uuid, email, username, password, is_premium, created_at)
					VALUES ($1, $2, $3, $4, $5, NOW())
					returning uuid;
				`)).WithArgs("uuid", "email", "username", "password", true).
					WillReturnError(errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, tc := range cases {
		db, mock, _ := sqlmock.New()
		defer db.Close()

		repo := &Repository{
			Db: db,
		}

		tc.mockFunc(mock)

		res, err := repo.CreateUser(context.Background(), tc.request.(User))
		assert.Equal(t, res, tc.response.(User))
		assert.Equal(t, err, tc.err)
	}
}

func TestGetUserByID(t *testing.T) {
	timeNow := time.Now()

	cases := []testCase{
		{
			name:    "success",
			request: "uuid",
			response: User{
				UUID:      "uuid",
				Email:     "email",
				Username:  "username",
				Password:  "password",
				IsPremium: true,
				CreatedAt: timeNow,
			},
			mockFunc: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(`
				SELECT uuid, email, username, password, is_premium, created_at
					FROM users
					WHERE uuid = $1;`)).WithArgs("uuid").
					WillReturnRows(sqlmock.NewRows([]string{"uuid", "email", "username", "password", "is_premium", "created_at"}).
						AddRow("uuid", "email", "username", "password", true, timeNow))
			},
			err: nil,
		},
		{
			name:     "error",
			request:  "uuid",
			response: User{},
			mockFunc: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(`
				SELECT uuid, email, username, password, is_premium, created_at
					FROM users
					WHERE uuid = $1;`)).WithArgs("uuid").WillReturnError(errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, tc := range cases {
		db, mock, _ := sqlmock.New()
		defer db.Close()

		repo := &Repository{
			Db: db,
		}

		tc.mockFunc(mock)

		res, err := repo.GetUserByID(context.Background(), tc.request.(string))
		assert.Equal(t, res, tc.response.(User))
		assert.Equal(t, err, tc.err)
	}
}

func TestGetUserByEmail(t *testing.T) {
	timeNow := time.Now()
	cases := []testCase{
		{
			name:    "success",
			request: "email@gmail.com",
			response: User{
				UUID:      "uuid",
				Email:     "email@gmail.com",
				Username:  "username",
				Password:  "password",
				IsPremium: true,
				CreatedAt: timeNow,
			},
			mockFunc: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(`
				SELECT uuid, email, username, password, is_premium, created_at
					FROM users
					WHERE email = $1;`)).WithArgs("email@gmail.com").
					WillReturnRows(sqlmock.NewRows([]string{"uuid", "email", "username", "password", "is_premium", "created_at"}).
						AddRow("uuid", "email@gmail.com", "username", "password", true, timeNow))
			},
			err: nil,
		},
		{
			name:     "error",
			request:  "email@gmail.com",
			response: User{},
			mockFunc: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(`
				SELECT uuid, email, username, password, is_premium, created_at
					FROM users
					WHERE email = $1;`)).WithArgs("email@gmail.com").WillReturnError(errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, tc := range cases {
		db, mock, _ := sqlmock.New()
		defer db.Close()

		repo := &Repository{
			Db: db,
		}

		tc.mockFunc(mock)

		res, err := repo.GetUserByEmail(context.Background(), tc.request.(string))
		assert.Equal(t, res, tc.response.(User))
		assert.Equal(t, err, tc.err)
	}
}

func TestGetUserByUsername(t *testing.T) {
	timeNow := time.Now()
	cases := []testCase{
		{
			name:    "success",
			request: "username",
			response: User{
				UUID:      "uuid",
				Email:     "email@gmail.com",
				Username:  "username",
				Password:  "password",
				IsPremium: true,
				CreatedAt: timeNow,
			},
			mockFunc: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(`
				SELECT uuid, email, username, password, is_premium, created_at
					FROM users
					WHERE username = $1;`)).WithArgs("username").
					WillReturnRows(sqlmock.NewRows([]string{"uuid", "email", "username", "password", "is_premium", "created_at"}).
						AddRow("uuid", "email@gmail.com", "username", "password", true, timeNow))
			},
			err: nil,
		},
		{
			name:     "error",
			request:  "username",
			response: User{},
			mockFunc: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(`
				SELECT uuid, email, username, password, is_premium, created_at
					FROM users
					WHERE username = $1;`)).WithArgs("username").WillReturnError(errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, tc := range cases {
		db, mock, _ := sqlmock.New()
		defer db.Close()

		repo := &Repository{
			Db: db,
		}

		tc.mockFunc(mock)

		res, err := repo.GetUserByUsername(context.Background(), tc.request.(string))
		assert.Equal(t, res, tc.response.(User))
		assert.Equal(t, err, tc.err)
	}
}

func TestCreateSwipe(t *testing.T) {
	cases := []testCase{
		{
			name: "success",
			request: Swipe{
				UUID:      "uuid",
				UserID:    "user_id",
				Direction: "direction",
				TargetID:  "target_id",
			},
			response: Swipe{
				UUID: "uuid",
			},
			mockFunc: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(`
				INSERT INTO swipes (uuid, user_id, target_id, direction, created_at, updated_at)
					VALUES ($1, $2, $3, $4, NOW(), NOW())
					RETURNING uuid;
				`)).WithArgs("uuid", "user_id", "target_id", "direction").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("uuid"))
			},
			err: nil,
		},
		{
			name: "error",
			request: Swipe{
				UUID:      "uuid",
				UserID:    "user_id",
				Direction: "direction",
				TargetID:  "target_id",
			},
			response: Swipe{},
			mockFunc: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(`
				INSERT INTO swipes (uuid, user_id, target_id, direction, created_at, updated_at)
					VALUES ($1, $2, $3, $4, NOW(), NOW())
					RETURNING uuid;
				`)).WithArgs("uuid", "user_id", "target_id", "direction").
					WillReturnError(errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, tc := range cases {
		db, mock, _ := sqlmock.New()
		defer db.Close()

		repo := &Repository{
			Db: db,
		}

		tc.mockFunc(mock)

		res, err := repo.CreateSwipe(context.Background(), tc.request.(Swipe))
		assert.Equal(t, res, tc.response.(Swipe))
		assert.Equal(t, err, tc.err)
	}
}

func TestUpdateSwipe(t *testing.T) {
	cases := []testCase{
		{
			name: "success",
			request: Swipe{
				UUID:      "uuid",
				UserID:    "user_id",
				Direction: "direction",
				TargetID:  "target_id",
			},
			response: Swipe{
				UUID:      "uuid",
				UserID:    "user_id",
				Direction: "direction",
				TargetID:  "target_id",
			},
			mockFunc: func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(`
				UPDATE swipes
					SET direction = $1,
					updated_at = NOW()
					WHERE uuid = $2;
				`)).WithArgs("direction", "uuid").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			err: nil,
		},
		{
			name: "error",
			request: Swipe{
				UUID:      "uuid",
				UserID:    "user_id",
				Direction: "direction",
				TargetID:  "target_id",
			},
			response: Swipe{},
			mockFunc: func(m sqlmock.Sqlmock) {
				m.ExpectExec(regexp.QuoteMeta(`
				UPDATE swipes
					SET direction = $1,
					updated_at = NOW()
					WHERE uuid = $2;
				`)).WithArgs("direction", "uuid").
					WillReturnError(errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, tc := range cases {
		db, mock, _ := sqlmock.New()
		defer db.Close()

		repo := &Repository{
			Db: db,
		}

		tc.mockFunc(mock)

		res, err := repo.UpdateSwipe(context.Background(), tc.request.(Swipe))
		assert.Equal(t, res, tc.response.(Swipe))
		assert.Equal(t, err, tc.err)
	}
}

func TestGetSwipeByUserIdAndTargetId(t *testing.T) {
	timeNow := time.Now()
	cases := []testCase{
		{
			name:    "success",
			request: "uuid-user,uuid-target",
			response: Swipe{
				UUID:      "uuid",
				UserID:    "uuid-user",
				Direction: "direction",
				TargetID:  "uuid-target",
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
			},
			mockFunc: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(`
				SELECT uuid, user_id, target_id, direction, created_at, updated_at
					FROM swipes
					WHERE user_id = $1
					AND target_id = $2;`)).WithArgs("uuid-user", "uuid-target").
					WillReturnRows(sqlmock.NewRows([]string{"uuid", "user_id", "target_id", "direction", "created_at", "updated_at"}).
						AddRow("uuid", "uuid-user", "uuid-target", "direction", timeNow, timeNow))
			},
			err: nil,
		},
		{
			name:     "error",
			request:  "uuid-user,uuid-target",
			response: Swipe{},
			mockFunc: func(m sqlmock.Sqlmock) {
				m.ExpectQuery(regexp.QuoteMeta(`
				SELECT uuid, user_id, target_id, direction, created_at, updated_at
					FROM swipes
					WHERE user_id = $1
					AND target_id = $2;`)).WithArgs("uuid-user", "uuid-target").WillReturnError(errors.New("error"))
			},
			err: errors.New("error"),
		},
	}

	for _, tc := range cases {
		db, mock, _ := sqlmock.New()
		defer db.Close()

		repo := &Repository{
			Db: db,
		}

		tc.mockFunc(mock)

		arrReq := strings.Split(tc.request.(string), ",")

		res, err := repo.GetSwipeByUserIdAndTargetId(context.Background(), arrReq[0], arrReq[1])
		assert.Equal(t, res, tc.response.(Swipe))
		assert.Equal(t, err, tc.err)
	}
}
