package store

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/HT0323/go_todo_app/clock"
	"github.com/HT0323/go_todo_app/entity"
	"github.com/HT0323/go_todo_app/testutil"
	"github.com/jmoiron/sqlx"
)

func TestRepository_Register(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	c := clock.FixedClocker{}
	var wantID int64 = 1
	okUser := &entity.User{
		Name:     "test",
		Password: "test",
		Role:     "test",
		Created:  c.Now(),
		Modified: c.Now(),
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = db.Close() })
	mock.ExpectExec(regexp.QuoteMeta(
		`INSERT INTO user (
		name, password, role, created, modified
		) VALUES (?,?,?,?,?)`),
	).WithArgs(okUser.Name, okUser.Password, okUser.Role, c.Now(), c.Now()).
		WillReturnResult(sqlmock.NewResult(wantID, 1))

	xdb := sqlx.NewDb(db, "mysql")
	r := &Repository{Clocker: c}
	if err := r.RegisterUser(ctx, xdb, okUser); err != nil {
		t.Errorf("want no error, but got %v", err)
	}
}

func TestRepository_GetUser(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	_, wantName := prepareUser(ctx, t, tx)

	c := clock.FixedClocker{}

	r := &Repository{Clocker: c}
	u, err := r.GetUser(ctx, tx, wantName)
	if err != nil {
		t.Errorf("want no error, but got %v", err)
	}
	if u.Name != wantName {
		t.Errorf("want %s, but %s", wantName, u.Name)
	}
}
