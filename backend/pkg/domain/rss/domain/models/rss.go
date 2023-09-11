package models

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Item        []Item `xml:"item"`
	Language	string `xml:"language" json:"language,omitempty"`
	Copyright	string `xml:"copyright" json:"copyright,omitempty"`
	ManagingEditor 	string `xml:"managingEditor" json:"managingEditor,omitempty"`
	WebMaster	string `xml:"webMaster" json:"webMaster,omitempty"`
	PubDate		string `xml:"pubDate" json:"pubDate,omitempty"`
	LastBuildDate	string `xml:"lastBuildDate" json:"lastBuildDate,omitempty"`
	Category	string `xml:"category" json:"category,omitempty"`
	Generator	string `xml:"generator" json:"generator,omitempty"`
	Docs		string `xml:"docs" json:"docs,omitempty"`
	Cloud		string `xml:"cloud" json:"cloud,omitempty"`
	Ttl			string `xml:"ttl" json:"ttl,omitempty"`
	Image		string `xml:"image" json:"image,omitempty"`
	Rating		string `xml:"rating" json:"rating,omitempty"`
	TextInput	string `xml:"textInput" json:"textInput,omitempty"`
	SkipHours	string `xml:"skipHours" json:"skipHours,omitempty"`
	SkipDays	string `xml:"skipDays" json:"skipDays,omitempty"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Author      string `xml:"author" json:"author,omitempty"`
	Category    string `xml:"category" json:"category,omitempty"`
	Comments    string `xml:"comments" json:"comments,omitempty"`
	Enclosure   string `xml:"enclosure" json:"enclosure,omitempty"`
	Guid        string `xml:"guid" json:"guid,omitempty"`
	PubDate     string `xml:"pubDate" json:"pubDate,omitempty"`
	Source      string `xml:"source" json:"source,omitempty"`
}