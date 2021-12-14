package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/spencercotton/coupon-challenge/internal/core/service"
	interalGRPC "github.com/spencercotton/coupon-challenge/internal/grpc"
	"github.com/spencercotton/coupon-challenge/internal/grpc/pb"
	couponRepository "github.com/spencercotton/coupon-challenge/internal/repositories/coupon"
	"github.com/spencercotton/coupon-challenge/internal/storage"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	svc *coupon.Service
	db  *sql.DB
	dsn string
)

func main() {
	dsn = "local:local1234@tcp(db:3306)/coupons"
	ctx := context.Background()

	if err := bootstrap(ctx); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err := runGRPC(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func bootstrap(ctx context.Context) error {
	// Configure and ping DB
	connector, err := storage.SQLConnector(dsn)
	if err != nil {
		return err
	}
	db = sql.OpenDB(connector)

	err = storage.Migrate(db)
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	repo := couponRepository.NewMySQLCouponRepository(db)
	svc = coupon.NewCouponService(repo)
	return nil
}

func runGRPC(ctx context.Context) error {
	s := grpc.NewServer()

	pb.RegisterCouponServiceServer(
		s,
		interalGRPC.NewGRPCCouponService(svc),
	)

	lis, err := net.Listen("tcp", ":50051")

	log.Printf("server listening at %v", lis.Addr())

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return err
	}

	return nil
}
