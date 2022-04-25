package main

import (
	"chiselwebproxy/internal/web"
	"fmt"

	ds "chiselwebproxy/internal/dataset"
	"chiselwebproxy/internal/remoteprocedurecall"
	s "chiselwebproxy/internal/server"
	"chiselwebproxy/internal/versioninfo"
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
