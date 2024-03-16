package server

import (
	"database/sql"
	actor2 "github.com/lilpipidron/vk-godeveloper-task/api/server/handler/actor"
	film2 "github.com/lilpipidron/vk-godeveloper-task/api/server/handler/film"
	"github.com/lilpipidron/vk-godeveloper-task/db/actor"
	"github.com/lilpipidron/vk-godeveloper-task/db/film"
	"log"
	"net/http"
)

func Start(db *sql.DB) error {
	actorRepo := actor.Repository{Driver: db}
	filmRepo := film.Repository{Driver: db}

	mux := http.NewServeMux()
	actor2.AddActorInMux(mux, actorRepo)
	film2.AddFilmInMux(mux, filmRepo)

	log.Println("server listening on port 8080")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		return err
	}

	return nil
}
