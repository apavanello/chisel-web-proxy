package server

import (
	"chisel-web-proxy/internal/dataset"
	pb "chisel-web-proxy/pkg/gatewaypb"
	"context"
	"fmt"
	"time"

	chisel "github.com/jpillora/chisel/client"
)

var connStats bool = false

var connSettings struct {
	Environment string
	Host        string
	LocalPort   int32
}

var ctx2 context.Context
var cancel context.CancelFunc

func (s *Server) Connect(ctx context.Context, in *pb.ConnectRequest) (*pb.ConnectResponse, error) {
	fmt.Println("ConnectService")

	ctx2, cancel = context.WithCancel(context.Background())
	remote, cSrv, proxy := dataset.GetConnString(in.Environment, in.Host)
	fmt.Println("remote:", remote, "cSrv:", cSrv, "localPort:", in.LocalPort, "proxy:", proxy)

	if in.Environment == "" || in.Host == "" {
		return &pb.ConnectResponse{
			HasConnected: false,
			Status:       "Error",
			Error:        "Not Allowed null values",
		}, nil
	}

	if connStats {
		return &pb.ConnectResponse{
			HasConnected: false,
			Status:       "Error",
			Error:        "Not Allowed two connections",
		}, nil
	}

	proxyStr := fmt.Sprintf("%d:%s", in.LocalPort, remote)

	// Connect to Chisel
	c := chisel.Config{
		Server:           cSrv,
		Remotes:          []string{proxyStr},
		MaxRetryCount:    3,
		MaxRetryInterval: time.Second * 1,
		Proxy:            proxy,
	}

	cli, err := chisel.NewClient(&c)
	if err != nil {
		return nil, err
	}

	cli.Debug = true

	connStats = true
	err = cli.Start(ctx2)
	if err != nil {
		fmt.Println(err)
	}

	connSettings.Environment = in.GetEnvironment()
	connSettings.Host = in.GetHost()
	connSettings.LocalPort = in.GetLocalPort()

	fmt.Printf("Connected")

	return &pb.ConnectResponse{
		HasConnected: true,
		Status:       fmt.Sprintf("Connected -> %s -> %s:%v", connSettings.Environment, connSettings.Host, connSettings.LocalPort),
		Error:        "",
	}, nil
}

func (s *Server) Disconnect(ctx context.Context, in *pb.DisconnectRequest) (*pb.DisconnectResponse, error) {
	fmt.Println("DisconnectService")

	if !connStats {
		return &pb.DisconnectResponse{
			Status: "Error",
			Error:  "Already Disconnected",
		}, nil
	}

	fmt.Println("Force Disconnect")
	cancel()
	<-ctx2.Done()
	connStats = false
	fmt.Println("Disconnected")

	return &pb.DisconnectResponse{
		Status: "Disconnected",
		Error:  "",
	}, nil
}

func (s *Server) Status(ctx context.Context, in *pb.StatusRequest) (*pb.StatusResponse, error) {
	fmt.Println("StatusService")

	return &pb.StatusResponse{
		Status: "Connected: " + fmt.Sprint(connStats),
	}, nil
}
