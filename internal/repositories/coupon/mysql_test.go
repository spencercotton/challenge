package coupon_test

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	repository "github.com/spencercotton/coupon-challenge/internal/repositories/coupon"
	"regexp"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("unexpected sqlmock err: %s", err)
	}
	defer db.Close()

	stmt := "INSERT INTO `coupons` (id, name, brand, value, expiry) values (?,?,?,?,?)"

	mock.ExpectExec(regexp.QuoteMeta(stmt)).WithArgs("1234", "Test Name", "Test Brand", 10, "Tomorrow").WillReturnResult(sqlmock.NewResult(0, 1))

	repo := repository.NewMySQLCouponRepository(db)
	err = repo.Create(context.Background(), "1234", "Test Name", "Test Brand", 10, "Tomorrow")

	if err != nil {
		t.Error(err)
	}
}

func TestList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("unexpected sqlmock err: %s", err)
	}
	defer db.Close()

	stmt := "SELECT * FROM `coupons`"
	rows := sqlmock.NewRows([]string{"id", "name", "brand", "value", "created_at", "expiry"}).
		AddRow("1234", "Test Name", "Test Brand", 10, time.Now().Format("2006-01-02 15:04:05"), "Tomorrow").
		AddRow("5678", "Test Name 2", "Test Brand 2", 10, time.Now().Format("2006-01-02 15:04:05"), "Next Week")

	mock.ExpectQuery(regexp.QuoteMeta(stmt)).WillReturnRows(rows).RowsWillBeClosed()

	repo := repository.NewMySQLCouponRepository(db)

	_, err = repo.List(context.Background())

	if err != nil {
		t.Error(err)
	}
}

func TestFind(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("unexpected sqlmock err: %s", err)
	}
	defer db.Close()

	stmt := "SELECT * FROM `coupons` where `id` = ?"
	rows := sqlmock.NewRows([]string{"id", "name", "brand", "value", "created_at", "expiry"}).
		AddRow("1234", "Test Name", "Test Brand", 10, time.Now().Format("2006-01-02 15:04:05"), "Tomorrow")
	mock.ExpectQuery(regexp.QuoteMeta(stmt)).WithArgs("1234").WillReturnRows(rows).RowsWillBeClosed()

	repo := repository.NewMySQLCouponRepository(db)

	_, err = repo.Find(context.Background(), "1234")

	if err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("unexpected sqlmock err: %s", err)
	}
	defer db.Close()

	stmt := "DELETE FROM `coupons` WHERE id ?"

	mock.ExpectExec(stmt).WithArgs("1234").WillReturnResult(sqlmock.NewResult(0, 1))

	repo := repository.NewMySQLCouponRepository(db)

	err = repo.Delete(context.Background(), "1234")

	if err != nil {
		t.Error(err)
	}
}
