package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/midepeter/grpc-service/db"
	"github.com/midepeter/grpc-service/proto/userpb"
	"github.com/midepeter/grpc-service/server"
	"google.golang.org/grpc"
)

const (
	port = ":5000"
)

var (
	d db.Db
)

func main() {
	// var (
	// 	certFile string = "server.crt"
	// 	keyFile  string = "server.key"
	// )
	_, err := d.Setup(context.Background(), "")
	if err != nil {
		errors.Unwrap(err)
		return
	}

	// srv, err := setUpTLS(certFile, keyFile)
	// if err != nil {
	// 	panic(fmt.Errorf("Failed while setting up tls %v\n", err))
	// }

	srv := grpc.NewServer()

	userpb.RegisterUserServer(srv, &server.Server{})

	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(fmt.Errorf("Failed while listen on port %s with error %v\n", port, err))
	}

	fmt.Println("Spinning server on port ", port)
	if err := srv.Serve(lis); err != nil {
		panic(fmt.Errorf("Failed while serve on port %s with error %v\n", port, err))
	}

	var cancel context.CancelFunc
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	select {
	case <-ctx.Done():
		log.Println("Shutting down server........")
		cancel()
	}
}

// func setUpTLS(certFile, keyFile string) (*grpc.Server, error) {
// 	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
// 	if err != nil {
// 		return nil, fmt.Errorf("Unable to parse certificates: %v\n", err)

// 	}

// 	options := []grpc.ServerOption{
// 		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
// 	}

// 	srv := grpc.NewServer(options...)
// 	return srv, nil
// }