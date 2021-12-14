package coupon

import (
	"context"
	"github.com/spencercotton/coupon-challenge/internal/core/entity"
)

type Repository interface {
	Create(ctx context.Context, couponID string, name string, brand string, value int64, expiry string) error
	List(ctx context.Context) ([]*entity.Coupon, error)
	Find(ctx context.Context, couponID string) (*entity.Coupon, error)
	Delete(ctx context.Context, couponID string) error
}
