package main

import (
	heartbeat "github.com/MarluxGitHub/instagram/pkg/domain/heartbeat/interfaces"
	"github.com/MarluxGitHub/instagram/pkg/lib/api/rest"
)

func main() {
	initInterfaces()

}

func initInterfaces() {
	restApi := rest.NewRestApi()
	restApi.Init()

	heartbeat := heartbeat.NewHeartbeatRestEndpoint()
	restApi.AddEndpoint(&heartbeat)

	restApi.Start()
}
