package grpc

import (
	"context"
	"github.com/spencercotton/coupon-challenge/internal/core/entity"
	coupon "github.com/spencercotton/coupon-challenge/internal/core/service"
	"github.com/spencercotton/coupon-challenge/internal/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type CouponService struct {
	pb.UnimplementedCouponServiceServer
	svc *coupon.Service
}

func NewGRPCCouponService(svc *coupon.Service) *CouponService {
	return &CouponService{
		svc: svc,
	}
}

func (s *CouponService) CreateCoupon(ctx context.Context, req *pb.CreateCouponRequest) (*pb.Coupon, error) {
	c, err := s.svc.Create(ctx, req.GetName(), req.GetBrand(), req.GetValue(), req.GetExpiry())

	if err != nil {
		log.Printf("Error creating coupon.")
	}

	// Map model to pb response
	couponResponse := pb.Coupon{
		Id:        c.ID,
		Name:      c.Name,
		Brand:     c.Brand,
		Value:     c.Value,
		CreatedAt: c.CreatedAt,
		Expiry:    c.Expiry,
	}

	return &couponResponse, nil
}

func (s *CouponService) GetCoupons(ctx context.Context, req *pb.GetCouponsRequest) (*pb.GetCouponsResponse, error) {
	coupons, _ := s.svc.List(ctx)

	response := pb.GetCouponsResponse{}

	for i := range coupons {
		// Map models to pb response
		couponResponse := s.buildCouponResponse(ctx, coupons[i])
		response.Coupons = append(response.Coupons, couponResponse)
	}

	return &response, nil
}

func (s *CouponService) GetCoupon(ctx context.Context, req *pb.GetCouponRequest) (*pb.Coupon, error) {
	c, err := s.svc.Find(ctx, req.GetId())

	st := status.New(codes.OK, "ok")

	if err != nil {
		st = status.New(codes.NotFound, "coupon not found")
		log.Printf("Error finding coupon.")
		return nil, st.Err()
	}

	if c == nil {
		st = status.New(codes.NotFound, "coupon not found")
		return nil, st.Err()
	}

	response := s.buildCouponResponse(ctx, c)

	return response, nil
}

func (s *CouponService) DeleteCoupon(ctx context.Context, req *pb.DeleteCouponRequest) (*pb.DeleteCouponResponse, error) {
	err := s.svc.Delete(ctx, req.GetId())

	if err != nil {
		log.Fatal(err)
	}

	response := pb.DeleteCouponResponse{}

	return &response, nil
}

func (s *CouponService) buildCouponResponse(ctx context.Context, coupon *entity.Coupon) *pb.Coupon {
	return &pb.Coupon{
		Id:        coupon.ID,
		Name:      coupon.Name,
		Brand:     coupon.Brand,
		Value:     coupon.Value,
		CreatedAt: coupon.CreatedAt,
		Expiry:    coupon.Expiry,
	}
}
