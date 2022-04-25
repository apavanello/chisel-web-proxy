package server

import (
	"chiselwebproxy/internal/dataset"
	pb "chiselwebproxy/pkg/gatewaypb"
	"context"
	"fmt"
	"time"

	chisel "github.com/jpillora/chisel/client"
)

var connStats bool = false
var forceDisconnect bool = false
var connSettings struct {
	Environment string
	Host        string
	LocalPort   int32
}

func (s *Server) Connect(ctx context.Context, in *pb.ConnectRequest) (*pb.ConnectResponse, error) {
	fmt.Println("ConnectService")

	remote, cSrv := dataset.GetConnString(in.Environment, in.Host)
	fmt.Println("remote:", remote, "cSrv:", cSrv, "localPort:", in.LocalPort)

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

	fmt.Println("proxyStr:", proxyStr)

	// Connect to Chisel
	c := chisel.Config{
		Server:           cSrv,
		Remotes:          []string{proxyStr},
		MaxRetryCount:    3,
		MaxRetryInterval: time.Second * 1,
	}

	cli, err := chisel.NewClient(&c)
	if err != nil {
		return nil, err
	}

	ctx2 := context.Background()
	ctx2, cancel := context.WithCancel(ctx2)

	connStats = true
	err = cli.Start(ctx2)
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for connStats {
			//time.Sleep(1 * time.Second)
			if forceDisconnect {
				fmt.Println("Force Disconnect")
				cancel()
				if err != nil {
					fmt.Println(err)
				}

				<-ctx2.Done()
				forceDisconnect = false
				connStats = false
				fmt.Println("Disconnected")
			}
		}

	}()

	connSettings.Environment = in.GetEnvironment()
	connSettings.Host = in.GetHost()
	connSettings.LocalPort = in.GetLocalPort()

	fmt.Printf("Connected \n Environment: %s \n Host: %s:%v", connSettings.Environment, connSettings.Host, connSettings.LocalPort)

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

	forceDisconnect = true

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
