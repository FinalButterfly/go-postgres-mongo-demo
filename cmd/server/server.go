package main

import (
	"GoNews/pkg/api"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/mongo"
	"log"
	"net/http"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// Создаём объект сервера.
	var srv server

	// pass := os.Getenv("postgresdbpass")
	// cn := "postgresql://postgres:" + pass + "@localhost/gonews"
	// db, err := postgres.New(cn)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	db, err := mongo.New("mongodb://localhost:27017/")
	if err != nil {
		log.Fatal(err)
	}
	srv.db = db

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
	// Предаём серверу маршрутизатор запросов,
	// поэтому сервер будет все запросы отправлять на маршрутизатор.
	// Маршрутизатор будет выбирать нужный обработчик.
	http.ListenAndServe(":8080", srv.api.Router())
}
