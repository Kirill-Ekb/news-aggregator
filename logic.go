package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// LoadConfiguration функция загрузки конфигурации
func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

// GetXMLByURL функция получения XML по GET запросу к URL
func GetXMLByURL(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	xml, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Read body: %v", err)
	}

	return xml, nil
}

// ParseRSSByXML функция парсинга XML RSS
func ParseRSSByXML(data []byte) ([]News, error) {
	rss := RSS{}
	err := xml.Unmarshal(data, &rss)
	if err != nil {
		return rss.News, fmt.Errorf("Unmarshal error: %v", err)
	}
	// Парсинг даты публикации
	for i, item := range rss.News {
		tmp, err := time.Parse("Mon, 2 Jan 2006 15:04:05 -0700", item.PubDateStr)
		if err != nil {
			return rss.News, fmt.Errorf("Time parse error: %v", err)
		}
		rss.News[i].PubDate = tmp
		rss.News[i].LogoLink = rss.ImageURL
	}
	return rss.News, nil
}

// UpdateRSS функция обновления списка новостей
func UpdateRSS(urls []string) {
	for _, url := range urls {
		data, err := GetXMLByURL(url)
		if err != nil {
			log.Println(err)
		}

		rss, err := ParseRSSByXML(data)
		if err != nil {
			log.Println(err)
		}

		err = SaveRSSToSQL(rss)
		if err != nil {
			log.Println(err)
		}
	}
}

// StartTickerFoUpdateRSS функция запуска счётчика тиков для обновления новостей
func StartTickerFoUpdateRSS(urls []string, UpdateMinutes int) {
	ticker := time.NewTicker(time.Minute * time.Duration(UpdateMinutes))
	go func() {
		for range ticker.C {
			UpdateRSS(urls)
		}
	}()
}
