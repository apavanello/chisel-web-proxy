package server

import (
	pb "chisel-web-proxy/pkg/gatewaypb"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedProxyServiceServer
	pb.UnimplementedPreLoadServiceServer
}

func StartServer(s *grpc.Server, port int) {
	//	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	//	if err != nil {
	//		log.Fatalf("failed to listen: %v", err)
	//	}

	pb.RegisterProxyServiceServer(s, &Server{})
	pb.RegisterPreLoadServiceServer(s, &Server{})
	//	log.Printf("server listening at %v", lis.Addr())
	//	if err := s.Serve(lis); err != nil {
	//		log.Fatalf("failed to serve: %v", err)
	//	}
}
