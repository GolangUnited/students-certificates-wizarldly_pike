package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gus_certificates/app/server"
	certSPb "gus_certificates/transport/protobuf/certificate"
)

func main() {
	host := os.Getenv("CERTIFICATES_HOST")
	if host == "" {
		log.Fatal(fmt.Errorf("environment variable %q not set", "CERTIFICATES_HOST"))
	}

	rpcPort := os.Getenv("RPC_CERTIFICATES_PORT")
	if rpcPort == "" {
		log.Fatal(fmt.Errorf("environment variable %q not set", "RPC_CERTIFICATES_PORT"))
	}

	restPort := os.Getenv("REST_CERTIFICATES_PORT")
	if restPort == "" {
		log.Fatal(fmt.Errorf("environment variable %q not set", "REST_CERTIFICATES_PORT"))
	}

	rpcAddr, restAddr := net.JoinHostPort(host, rpcPort), net.JoinHostPort(host, restPort)
	if restServer := os.Getenv("REST_SERVER_ENABLE"); restServer == "true" {
		go runRestServer(rpcAddr, restAddr)
	}

	err := runRpcServer(rpcAddr)
	if err != nil {
		log.Fatal(err)
	}
}

func runRpcServer(restAddr string) error {
	listener, err := net.Listen("tcp", restAddr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	server, err := server.NewCertificateServer()
	if err != nil {
		return err
	}

	certSPb.RegisterCertificateServer(grpcServer, server)
	return grpcServer.Serve(listener)
}

func runRestServer(rpcAddr, restAddr string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := certSPb.RegisterCertificateHandlerFromEndpoint(ctx, mux, rpcAddr, opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(restAddr, mux); err != nil {
		log.Fatal(err)
	}
}
