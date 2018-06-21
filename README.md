# news-aggregator

Простой агрегатор новостей

## Файл конфигурации

Настройка производится с помощью файла конфигурации "config.json" в формате JSON.

### Список полей

* Webport - Строка реквизитов в формате "IP:Port", на которых будет поднят веб-сервер, например, ":8080"
* FileNameDB - Строка с адресом расположения sqlite базы данных, например, "news.db"
* UpdateMinutes - Целочисленное значение количества минут между обновлениями новостей, например, 1
* RSS - Массив строк с ссылками на RSS ленты новостей, например, "https://meduza.io/rss/all"

### Пример файла конфигурации

```
{
    "Webport": ":8080",
    "FileNameDB": "news.db",
    "UpdateMinutes": 1,
    "RSS": [
        "https://www.e1.ru/news/rdf/full.xml",
        "https://meduza.io/rss/all",
        "http://static.feed.rbc.ru/rbc/logical/footer/news.rss"
    ]
}
```

## Создано с использованием

* [mux](http://github.com/gorilla/mux)
* [gorm](http://github.com/jinzhu/gorm)