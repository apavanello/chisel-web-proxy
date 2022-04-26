package main

import (
	"chisel-web-proxy/internal/web"
	"fmt"

	ds "chisel-web-proxy/internal/dataset"
	"chisel-web-proxy/internal/remoteprocedurecall"
	s "chisel-web-proxy/internal/server"
	"chisel-web-proxy/internal/versioninfo"
)

func main() {
	fmt.Println("Project: " + versioninfo.ProjectName)
	fmt.Println("Description: " + versioninfo.ProjectDescription)
	fmt.Println("Copyright: " + versioninfo.ProjectCopyright)
	fmt.Println("Version: " + versioninfo.Version)
	fmt.Println("Revision: " + versioninfo.Revision)
	fmt.Println("Branch: " + versioninfo.Branch)

	ds.LoadDataset()

	rpcServer := remoteprocedurecall.NewServer()
	s.StartServer(rpcServer.Grpc, 9001)

	// Starts the HTTP server
	web.Start(":9000", rpcServer)

}
