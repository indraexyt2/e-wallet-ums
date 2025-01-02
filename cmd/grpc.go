package cmd

import (
	pb "e-wallet-ums/cmd/proto/tokenvalidation"
	"e-wallet-ums/helpers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func ServeGRPC() {
	dependency := dependencyInject()

	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", ""))
	if err != nil {
		log.Fatal("Failed to listen grpc port: ", err)
	}
	s := grpc.NewServer()

	// list method
	pb.RegisterTokenValidationServer(s, dependency.TokenValidationAPI)
	logrus.Info("start grpc server on port: ", helpers.GetEnv("GRPC_PORT", "5000"))
	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to serve grpc port: ", err)
	}
}
