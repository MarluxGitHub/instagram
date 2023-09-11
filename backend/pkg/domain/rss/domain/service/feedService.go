package service

import "github.com/MarluxGitHub/instagram/pkg/domain/rss/domain/models"


type FeedService struct{
	providerService ProviderService
	rssReader RssReader
}

func (feedService *FeedService) GetAllFeeds() []models.Rss {
	feeds := make([]models.Rss, len(feedService.providerService.GetAll()))

	for i, provider := range feedService.providerService.GetAll() {
		feeds[i] = feedService.getFeedsOfProvider(&provider)
	}

	return feeds
}

func (feedService *FeedService) GetAllFeedsOfProvider(providerId int) models.Rss{
	feeds := feedService.getFeedsOfProvider(feedService.providerService.Get(providerId))

	return feeds
}

func (feedService *FeedService) GetFeedOfProvider(providerId int, id int) models.Item {
	feeds := feedService.getFeedsOfProvider(feedService.providerService.Get(providerId))

	return feeds.Channel.Item[id]
}

func (feedService *FeedService) getFeedsOfProvider(provider *models.Provider) models.Rss {
	return feedService.rssReader.Read(provider)
}

func (feedService *FeedService) Add(provider models.Provider) models.Provider {
	return feedService.providerService.Add(provider)
}

func NewFeedService() FeedService {
	return FeedService{
		providerService: NewProviderService(),
		rssReader: RssReader{},
	}
}

func (FeedService *FeedService) Initialize(c models.Configuration) {
	for _, provider := range c.Providers {
		FeedService.providerService.Add(models.Provider{
			URI: provider,
		})
	}
}


