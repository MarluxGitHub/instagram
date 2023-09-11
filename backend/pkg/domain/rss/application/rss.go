package application

import (
	"github.com/MarluxGitHub/instagram/pkg/domain/rss/domain/models"
	"github.com/MarluxGitHub/instagram/pkg/domain/rss/domain/service"
)

type RssApplication struct {
	feedService service.FeedService
}

func (rss *RssApplication) GetAllFeeds() []models.Rss {
	return rss.feedService.GetAllFeeds()
}

func (rss *RssApplication) GetAllFeedsOfProvider(providerId int) models.Rss {
	return rss.feedService.GetAllFeedsOfProvider(providerId)
}

func NewRssApplication(c models.Configuration) RssApplication {
	rssApp := RssApplication{
		feedService: service.NewFeedService(),
	}

	rssApp.feedService.Initialize(c)

	return rssApp
}