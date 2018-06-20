package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Загрузка конфигурации
	config, err := LoadConfiguration("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	// Инициализация базы данных
	err = InitSQL(config.FileNameDB)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	// Первичная загрузка/обновление списка новостей
	UpdateRSS(config.RSS)

	// Запуск тикера для мониторинга новый новостей
	StartTickerFoUpdateRSS(config.RSS, config.UpdateMinutes)

	r := mux.NewRouter()

	r.HandleFunc("/", handlerNews).Methods("GET")
	r.HandleFunc("/{search}", handlerNews).Methods("GET")

	log.Fatal(http.ListenAndServe(config.Webport, r))
}

func handlerNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tmpl, err := template.ParseFiles("tmpl/default.html", "tmpl/news.html")
	if err != nil {
		log.Fatal(err)
	}
	items := struct {
		Title string
		News  []News
	}{
		Title: "Новости",
		News:  GetNews(vars["search"])}
	if err := tmpl.Execute(w, items); err != nil {
		log.Fatal(err)
	}
}
