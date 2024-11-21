# Music API Service

## Установка окружения

- Тестовая среда:
  export APP_ENV=test

- Продакшен среда:
  export APP_ENV=prod

## Запуск приложения

- go run cmd/app/main.go


## Эндпоинты и cURL

### Создать песню
curl -X POST http://localhost:8081/api/song \
-H "Content-Type: application/json" \
-d '{
"group_name": "Muse",
"song_name": "Supermassive Black Hole",
"release_date": "2006-07-16T00:00:00Z",
"song_text": "Ooh baby, don’t you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
"link": "https://www.youtube.com/watch?v=Xsp3_a-PMTw"
}'


### Удалить песню
curl -X DELETE http://localhost:8081/api/song/1

### Получить песню по ID
curl -X GET "http://localhost:8081/api/song/1" -H "Accept: application/json"

### Обновить песню
curl -X PUT http://localhost:8081/api/song/1 \
-H "Content-Type: application/json" \
-d '{
"group_name": "Muse",
"song_name": "Supermassive Black Hole",
"release_date": "2006-07-16T00:00:00Z",
"song_text": "Updated song text",
"link": "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
"deleted": false
}'


### Получить все песни
curl -X GET "http://localhost:8081/api/songs?page=2&limit=5" -H "accept: application/json"


### Фильтр песен
curl -X 'GET'
'http://localhost:8081/api/songs/filter'
-H 'accept: application/json'

### Фильтры:
- `group_name` — Фильтр по имени группы
- `song_name` — Фильтр по названию песни
- `release_date` — Фильтр по дате (формат: YYYY-MM-DD)
- `song_text` — Фильтр по тексту песни
- `link` — Фильтр по ссылке