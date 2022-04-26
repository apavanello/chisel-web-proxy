package server

import (
	"chiselwebproxy/internal/dataset"
	pb "chiselwebproxy/pkg/gatewaypb"
	"context"
	"fmt"
)

func (s *Server) PreLoad(ctx context.Context, in *pb.PreLoadRequest) (*pb.PreLoadResponse, error) {
	fmt.Println("PreLoadService")

	data := dataset.GetEnv()
	var envs []string

	for k := range data {
		if k != "" {
			envs = append(envs, k)
		}
	}

	if connStats {
		return &pb.PreLoadResponse{
			Status: "Connected",
			ConnectRequest: &pb.ConnectRequest{
				Host:        connSettings.Host,
				LocalPort:   connSettings.LocalPort,
				Environment: connSettings.Environment,
			},
			Env: envs,
		}, nil
	}
	return &pb.PreLoadResponse{
		Status: "Disconnected",
		Env:    envs,
	}, nil
}

func (s *Server) GetHosts(ctx context.Context, in *pb.HostsRequest) (*pb.HostsResponse, error) {
	fmt.Println("GetHostsService")

	data := dataset.GetHosts(in.Env)
	var hosts []string

	for k := range data {
		if k != "" {
			hosts = append(hosts, k)
		}
	}

	return &pb.HostsResponse{
		Hosts: hosts,
	}, nil
}
