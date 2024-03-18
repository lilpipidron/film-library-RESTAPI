package login

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/lilpipidron/vk-godeveloper-task/db/actor"
	"github.com/lilpipidron/vk-godeveloper-task/db/film"
	actorHandler "github.com/lilpipidron/vk-godeveloper-task/server/handler/actor"
	filmHandler "github.com/lilpipidron/vk-godeveloper-task/server/handler/film"
)

type Application struct {
	Auth struct {
		Username string
		Password string
	}
}

func (app *Application) AdminAuth(db *sql.DB, mux *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			if username == app.Auth.Username && password == app.Auth.Username {
				app.handleAdminRequest(db, mux)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func (app *Application) UserAuth(db *sql.DB, mux *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.handleUserRequest(w, r, db, mux)
	}
}

func (app *Application) handleUserRequest(w http.ResponseWriter, r *http.Request, db *sql.DB, mux *http.ServeMux) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Println("Method not allowed", r.Method)
		return
	}
	actorRepository := actor.NewActorRepository(db)
	filmRepository := film.NewFilmRepository(db)

	actorHandlerRepository := actorHandler.NewActorRepository(*actorRepository)
	mux.HandleFunc("/user/actor", func(w http.ResponseWriter, r *http.Request) { actorHandlerRepository.Handler(w, r) })

	filmHandlerRepository := filmHandler.NewFilmRepository(*filmRepository)
	mux.HandleFunc("/user/film", func(w http.ResponseWriter, r *http.Request) { filmHandlerRepository.Handler(w, r) })

	log.Println("Authorized access for user")
}

func (*Application) handleAdminRequest(db *sql.DB, mux *http.ServeMux) {
	actorRepository := actor.NewActorRepository(db)
	filmRepository := film.NewFilmRepository(db)

	actorHandlerRepository := actorHandler.NewActorRepository(*actorRepository)
	mux.HandleFunc("/admin/actor", func(w http.ResponseWriter, r *http.Request) { actorHandlerRepository.Handler(w, r) })

	filmHandlerRepository := filmHandler.NewFilmRepository(*filmRepository)
	mux.HandleFunc("/admin/film", func(w http.ResponseWriter, r *http.Request) { filmHandlerRepository.Handler(w, r) })

	log.Println("Authorized access for admin")
}
