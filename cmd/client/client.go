package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/safeguardproperties/gocontentservice/pkg/contentservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func main() {
	tls := flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile := flag.String("ca_file", "testdata/ca.pem", "The file containning the CA root cert file")
	serverAddr := flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride := flag.String("server_host_override", "", "The server name use to verify the hostname returned by TLS handshake")

	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		var sn string
		if *serverHostOverride != "" {
			sn = *serverHostOverride
		}
		var creds credentials.TransportCredentials
		if *caFile != "" {
			var err error
			creds, err = credentials.NewClientTLSFromFile(*caFile, sn)
			if err != nil {
				grpclog.Fatalf("Failed to create TLS credentials %v", err)
			}
		} else {
			creds = credentials.NewClientTLSFromCert(nil, sn)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewContentServiceClient(conn)

	response, err := client.ListByOrderNumber(context.Background(), &pb.OrderNumberRequest{OrderNumber: 600016554})
	if err != nil {
		log.Fatalf("Error Pinging: %v", err)
	}
	fmt.Println("Response: ", response)
}
