package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	pb "github.com/divyag9/gocontentservice/pkg/contentservice"
	"github.com/divyag9/gocontentservice/pkg/list"
	"golang.org/x/net/context"
	ora "gopkg.in/rana/ora.v4"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

// Server contains information required by server
type Server struct {
	Db              *ora.Ses
	MemClient       *memcache.Client
	SecondsToExpiry int32
}

//ListByImageIds retrieves imagedetails for given imageids
func (s *Server) ListByImageIds(ctx context.Context, in *pb.ImageIdsRequest) (*pb.ListResponse, error) {
	//TO be implemented
	return nil, nil
}

//ListByOrderNumber retrieves imagedetails for given ordernumber
func (s *Server) ListByOrderNumber(ctx context.Context, in *pb.OrderNumberRequest) (*pb.ListResponse, error) {
	listInfo := &list.Info{}
	listInfo.OrderNumber = in.OrderNumber
	listInfo.Db = s.Db
	listInfo.MemClient = s.MemClient
	listInfo.SecondsToExpiry = s.SecondsToExpiry
	imageDetailList := &list.ImageDetailList{}
	imageDetailList.DatabaseCaller = listInfo
	imageDetailList.CacheCaller = listInfo
	imageDetails, err := imageDetailList.GetImageDetails()
	if err != nil {
		return nil, err
	}
	listResponse := &pb.ListResponse{}
	listResponse.ImageDetails = imageDetails

	return listResponse, nil
}

func main() {
	tls := flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile := flag.String("cert_file", "testdata/server1.pem", "The TLS cert file")
	keyFile := flag.String("key_file", "testdata/server1.key", "The TLS key file")
	port := flag.Int("port", 10000, "The server port")

	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			grpclog.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)

	dsn := os.Getenv("GO_OCI8_CONNECT_STRING")
	env, srv, ses, err := ora.NewEnvSrvSes(dsn)
	if err != nil {
		fmt.Println(err)
	}
	defer env.Close()
	defer srv.Close()
	defer ses.Close()

	mc := getMemcacheClient()
	expiryTime, _ := strconv.Atoi(os.Getenv("EXPIRY_TIME"))

	server := &Server{}
	server.Db = ses
	server.MemClient = mc
	server.SecondsToExpiry = int32(expiryTime)
	pb.RegisterContentServiceServer(grpcServer, server)
	if err := grpcServer.Serve(listen); err != nil {
		fmt.Println("Failed to serve: ", err)
	}
}

func getMemcacheClient() *memcache.Client {
	servers := os.Getenv("MEMCACHE_SERVERS")
	memcacheServers := strings.Split(servers, ",")
	mc := memcache.New(memcacheServers...)

	return mc
}
