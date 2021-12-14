package coupon

import (
	"context"
	"database/sql"
	"github.com/davecgh/go-spew/spew"
	"github.com/spencercotton/coupon-challenge/internal/core/entity"
	"log"
)

type MySQLCouponRepository struct {
	Repository
	db *sql.DB
}

func NewMySQLCouponRepository(db *sql.DB) *MySQLCouponRepository {
	return &MySQLCouponRepository{
		db: db,
	}
}

func (r *MySQLCouponRepository) Create(ctx context.Context, couponID string, name string, brand string, value int64, expiry string) error {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO `coupons` (id, name, brand, value, expiry) values (?,?,?,?,?)",
		couponID, name, brand, value, expiry,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *MySQLCouponRepository) Find(ctx context.Context, couponID string) (*entity.Coupon, error) {
	row, err := r.db.QueryContext(
		ctx,
		"SELECT * FROM `coupons` where `id` = ?",
		couponID,
	)

	if err != nil {
		log.Fatal(err)
	}

	e := entity.Coupon{}
	row.Next()
	err = row.Scan(&e.ID, &e.Name, &e.Brand, &e.Value, &e.CreatedAt, &e.Expiry)

	if err != nil {
		spew.Dump(err)
		return nil, err
	}

	return &e, nil
}

func (r *MySQLCouponRepository) List(ctx context.Context) ([]*entity.Coupon, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT * FROM `coupons`",
	)

	if err != nil {
		log.Fatal(err)
	}

	var entities []*entity.Coupon

	for rows.Next() {
		e := entity.Coupon{}
		rows.Scan(&e.ID, &e.Name, &e.Brand, &e.Value, &e.CreatedAt, &e.Expiry)
		entities = append(entities, &e)
	}

	return entities, nil
}

func (r *MySQLCouponRepository) Delete(ctx context.Context, couponID string) error {
	_, err := r.db.ExecContext(
		ctx,
		"DELETE FROM `coupons` WHERE id = ?",
		couponID,
	)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
