package application

import (
	"github.com/MarluxGitHub/instagram/pkg/domain/heartbeat/domain/model"
	"github.com/MarluxGitHub/instagram/pkg/domain/heartbeat/domain/service"
)

func GetHeartBeat() model.Heartbeat {
	service := service.NewHeartBeatService()

	return service.GetHeartBeat()
}