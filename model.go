package main

import (
	"encoding/xml"
	"time"
)

// Config структура JSON файла конфигурации
type Config struct {
	Webport       string   `json:"webport"`
	FileNameDB    string   `json:"FileNameDB"`
	UpdateMinutes int      `json:"UpdateMinutes"`
	RSS           []string `json:"RSS"`
}

// RSS структура для парсинга RSS
type RSS struct {
	Root     xml.Name `xml:"rss"`
	News     []News   `xml:"channel>item"`
	ImageURL string   `xml:"channel>image>url"`
}

// News структура одной новости
type News struct {
	ID         uint   `gorm:"primary_key"`
	Title      string `xml:"title"`
	Link       string `xml:"link" gorm:"unique_index"`
	PubDate    time.Time
	PubDateStr string `xml:"pubDate" gorm:"-"`
	LogoLink   string
}
