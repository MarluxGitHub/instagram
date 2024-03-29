package config

import (
	"github.com/MarluxGitHub/instagram/pkg/domain/rss/domain/models"
	"github.com/MarluxGitHub/instagram/pkg/lib/config/types"
)

type Configuration struct {
	Api types.Api
	DB  types.Db
	Rss models.Configuration
}

