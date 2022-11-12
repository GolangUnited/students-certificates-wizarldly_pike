package main

import (
	"gus_certificates/app/server"
	"log"
	"net"

	certSPb "gus_certificates/protobuf/transport/certificate"

	"google.golang.org/grpc"
)

func main() {
	host := "0.0.0.0"
	port := "1234"

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
