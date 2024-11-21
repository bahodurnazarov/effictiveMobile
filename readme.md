before start app please set the APP_ENV by running this command:

// for running test environment
# export APP_ENV=test 

// for running prod environment
# export APP_ENV=prod 

for running server
# go run cmd/app/main.go

curl for creating song:
curl -X POST http://localhost:8081/api/song \
-H "Content-Type: application/json" \
-d '{
"group_name": "Muse",
"song_name": "Supermassive Black Hole",
"release_date": "2006-07-16T00:00:00Z",
"song_text": "Ooh baby, don’t you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight",
"link": "https://www.youtube.com/watch?v=Xsp3_a-PMTw"
}'


curl for soft deleting:
curl -X DELETE http://localhost:8081/api/song/1

curl for get song by ID:
curl -X GET "http://localhost:8081/song/123" -H "Accept: application/json"
