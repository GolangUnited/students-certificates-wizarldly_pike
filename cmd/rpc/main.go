package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"gus_certificates/app/server"
	certSPb "gus_certificates/protobuf/transport/certificate"
)

func main() {
	host := os.Getenv("CERTIFICATES_HOST")
	if host == "" {
		log.Fatal(fmt.Errorf("environment variable %q not set", "CERTIFICATES_HOST"))
	}

	port := os.Getenv("CERTIFICATES_PORT")
	if port == "" {
		log.Fatal(fmt.Errorf("environment variable %q not set", "CERTIFICATES_PORT"))
	}

	err := runRpcServer(net.JoinHostPort(host, port))
	if err != nil {
		log.Fatal(err)
	}
}

func runRpcServer(addr string) error {
	listener, err := net.Listen("tcp", addr)
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
