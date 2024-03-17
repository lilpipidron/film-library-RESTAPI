package login

import (
	"database/sql"
	"github.com/lilpipidron/vk-godeveloper-task/db/actor"
	"github.com/lilpipidron/vk-godeveloper-task/db/film"
	actorHandler "github.com/lilpipidron/vk-godeveloper-task/server/handler/actor"
	filmHandler "github.com/lilpipidron/vk-godeveloper-task/server/handler/film"
	"log"
	"net/http"
)

type Application struct {
	Auth struct {
		Username string
		Password string
	}
}

func (app *Application) AdminAuth(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			if username == app.Auth.Username && password == app.Auth.Username {
				app.handleAdminRequest(w, r, db)
				return
			}
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
func (app *Application) UserAuth(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.handleUserRequest(w, r, db)
	}
}

func (app *Application) handleUserRequest(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Println("Method not allowed", r.Method)
		return
	}
	actorRepository := actor.NewActorRepository(db)
	filmRepository := film.NewFilmRepository(db)
	actorHandler.Handler(w, r, *actorRepository)
	filmHandler.Handler(w, r, *filmRepository)
	log.Println("Authorized access for user")
}
func (*Application) handleAdminRequest(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	actorRepository := actor.NewActorRepository(db)
	filmRepository := film.NewFilmRepository(db)
	actorHandler.Handler(w, r, *actorRepository)
	filmHandler.Handler(w, r, *filmRepository)
	log.Println("Authorized access for admin")
}
