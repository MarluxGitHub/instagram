package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHeartBeat(t *testing.T) {
    service := NewHeartBeatService()

    heartbeat := service.GetHeartBeat()

    assert.NotEmpty(t, heartbeat.Time, "Expected non-empty time, but got empty string")
    assert.Equal(t, "I'm alive!", heartbeat.Message, "Expected message 'I'm alive!', but got '%s'", heartbeat.Message)
}