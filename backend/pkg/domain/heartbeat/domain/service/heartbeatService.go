package service

import (
	"time"

	"github.com/MarluxGitHub/instagram/pkg/domain/heartbeat/domain/model"
)

type HeartbeatService interface {
	GetHeartBeat() model.Heartbeat
}

type HeartBeatServiceImpl struct {
}

func (service *HeartBeatServiceImpl) GetHeartBeat() model.Heartbeat {
	return model.Heartbeat{
		Time:   time.Now().String(),
		Message: "I'm alive!",
	}
}

func NewHeartBeatService() HeartbeatService {
	return &HeartBeatServiceImpl{}
}