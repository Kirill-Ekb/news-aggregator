package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB глобальный доступ к базе данных
var DB *gorm.DB

// InitSQL инициализация базы данных
func InitSQL(FileNameDB string) error {
	var err error
	DB, err = gorm.Open("sqlite3", FileNameDB)
	if err != nil {
		return err
	}
	// Включение логирования обращения к базе данных
	// DB.LogMode(true)

	DB.AutoMigrate(&News{})
	return nil
}

// SaveRSSToSQL функция сохранения новостей в базу данных
func SaveRSSToSQL(rss []News) error {
	for _, item := range rss {
		// Проверяем уникальность новости по её ссылке
		if DB.Where("Link = ?", item.Link).First(&News{}).RecordNotFound() {
			DB.Save(&item)
		}
	}

	return nil
}

// GetNews функция получения новостей из базы данных
func GetNews(search string) []News {
	var result []News
	if len(search) == 0 {
		DB.Limit(50).Order("pub_date desc").Find(&result)
	} else {
		DB.Order("pub_date desc").Where("title LIKE ?", "%"+search+"%").Find(&result)
	}
	return result
}
