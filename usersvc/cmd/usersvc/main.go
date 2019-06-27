package main

import (
	"log"
	"net"
	"os"

	"github.com/Buzzvil/stella/usersvc/internal/app/usersrv"

	"github.com/Buzzvil/stella/usersvc/internal/pkg/user"
	"github.com/Buzzvil/stella/usersvc/internal/pkg/user/pgrepo"
	pb "github.com/Buzzvil/stella/usersvc/pkg/proto"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	r := pgrepo.New(db)
	u := user.NewUsecase(r)
	srv := usersrv.New(u)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := newGrpcServer()

	pb.RegisterUserServiceServer(grpcServer, srv)

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	return
}

func newGrpcServer() *grpc.Server {
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	opts := []grpc_logrus.Option{}
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)
	return grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry, opts...),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.StreamServerInterceptor(logrusEntry, opts...),
		),
	)
}