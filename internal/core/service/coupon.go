package coupon

import (
	"context"
	"github.com/google/uuid"
	"github.com/spencercotton/coupon-challenge/internal/core/entity"
	"github.com/spencercotton/coupon-challenge/internal/repositories/coupon"
	"log"
)

type Service struct {
	repo coupon.Repository
}

func (c *Service) Create(ctx context.Context, name string, brand string, value int64, expiry string) (*entity.Coupon, error) {
	newID := uuid.NewString()

	err := c.repo.Create(ctx, newID, name, brand, value, expiry)

	if err != nil {
		return nil, err
	}

	log.Println("Created core.")

	coupon, err := c.Find(ctx, newID)

	if err != nil {
		return nil, err
	}

	return coupon, nil
}

func (c *Service) List(ctx context.Context) ([]*entity.Coupon, error) {
	coupons, err := c.repo.List(ctx)

	if err != nil {
		return nil, err
	}

	return coupons, nil
}

func (c *Service) Find(ctx context.Context, couponID string) (*entity.Coupon, error) {
	coupon, err := c.repo.Find(ctx, couponID)

	if err != nil {
		return nil, err
	}

	return coupon, nil
}

func (c *Service) Delete(ctx context.Context, couponID string) error {
	err := c.repo.Delete(ctx, couponID)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func NewCouponService(repo coupon.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
