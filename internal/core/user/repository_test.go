package user

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateUser(t *testing.T) {
	db,mock,err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	repo := NewUserRepo(db)

	userId := 1
	ctx := context.Background()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (id) VALUES ($1)")).
		WithArgs(userId).
		WillReturnResult(sqlmock.NewResult(1,1))

	err  = repo.CreateUser(ctx,userId)

	if err != nil {
		t.Errorf("createUser returned an error: %v",err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateUser_DBError(t *testing.T){
	db,mock,err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}

	defer db.Close()

	repo := NewUserRepo(db)

	userId := 1
	ctx := context.Background()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (id) VALUES ($1)")).
		WithArgs(userId). 
		WillReturnError(sql.ErrConnDone)
	
	err = repo.CreateUser(ctx,userId)

	if err == nil {
		t.Error(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepo(db)

	userId := 789

	
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	exists, err := repo.UserExists(userId)

	if err != nil {
		t.Errorf("UserExists returned an error: %v", err)
	}

	if !exists {
		t.Errorf("UserExists returned false, expected true")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUserExists_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepo(db)

	userId := 999

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`)).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	exists, err := repo.UserExists(userId)

	if err != nil {
		t.Errorf("UserExists returned an error: %v", err)
	}

	if exists {
		t.Errorf("UserExists returned true, expected false")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}