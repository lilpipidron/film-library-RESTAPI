package film

import (
	"encoding/json"
	filmStruct "github.com/lilpipidron/vk-godeveloper-task/api/types/film"
	"github.com/lilpipidron/vk-godeveloper-task/db/actorFilm"
	"github.com/lilpipidron/vk-godeveloper-task/db/film"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request, repository film.Repository) {
	if r.Method == http.MethodGet {
		GetMediator(w, r, &repository)
	} else if r.Method == http.MethodPut {
		AddNewFilm(w, r, &repository)
	} else if r.Method == http.MethodDelete {
		DeleteFilmByID(w, r, &repository)
	} else if r.Method == http.MethodPost {

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Println("Method not allowed", r.Method)
	}
}
func GetMediator(w http.ResponseWriter, r *http.Request, repository *film.Repository) {
	queryParams := r.URL.Query()
	title := queryParams.Get("title")
	actor := queryParams.Get("actor")
	if title == "" && actor == "" {
		GetAllFilms(w, queryParams, repository)
		return
	}

	if actor == "all" {
		GetActors(w, repository, title)
		return
	}
	GetFilmByTitleOrActorNameAndSurname(w, repository, title, actor)
}
func GetFilmByTitleOrActorNameAndSurname(w http.ResponseWriter, repository *film.Repository, title, actor string) {
	log.Println("request: get film by title or actor name and surname")
	films, err := repository.FindFilmByTitleOrActorName(title, actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	filmsJSON, err := json.Marshal(films)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(filmsJSON)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("request completed")
}
func GetActors(w http.ResponseWriter, repository *film.Repository, title string) {
	log.Println("request: get film's actors")
	actors, err := repository.FindAllActors(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	actorsJSON, err := json.Marshal(actors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(actorsJSON)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("request completed")
}
func GetAllFilms(w http.ResponseWriter, queryParams url.Values, repository *film.Repository) {
	log.Println("request: get all films and sort")
	films, err := repository.GetAllFilms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	sortField := queryParams.Get("sortField")
	sortType := queryParams.Get("sortType")
	switch sortField {
	case "ID":
		if sortType == "asc" {
			filmStruct.ByIDAsc(films)
		} else {
			filmStruct.ByIDDesc(films)
		}
	case "title":
		if sortType == "asc" {
			filmStruct.ByTitleAsc(films)
		} else {
			filmStruct.ByTitleDesc(films)
		}
	case "releaseDate":
		if sortType == "asc" {
			filmStruct.ByReleaseDateAsc(films)
		} else {
			filmStruct.ByReleaseDateDesc(films)
		}
	case "rating":
		if sortType == "asc" {
			filmStruct.ByRatingAsc(films)
		} else {
			filmStruct.ByRatingDesc(films)
		}
	default:
		filmStruct.ByRatingDesc(films)
	}
	filmsJSON, err := json.Marshal(films)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(filmsJSON)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("request completed")
}
func AddNewFilm(w http.ResponseWriter, r *http.Request, repository *film.Repository) {
	log.Println("request: put new film")
	queryParams := r.URL.Query()
	title := queryParams.Get("title")
	description := queryParams.Get("description")
	releaseDate, err := time.Parse("2006-01-02", queryParams.Get("releaseDate"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	rating, err := strconv.ParseFloat(queryParams.Get("rating"), 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	actors := strings.Split(queryParams.Get("actors"), ",")
	err = repository.AddNewFilm(title, description, releaseDate, float32(rating), actors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("request completed")
}

func DeleteFilmByID(w http.ResponseWriter, r *http.Request, repository *film.Repository) {
	log.Println("request: delete film by ID")
	queryParams := r.URL.Query()
	id, err := strconv.ParseInt(queryParams.Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	actorFilmRepository := &actorFilm.Repository{Driver: repository.Driver}
	err = actorFilmRepository.DeleteFilm(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	err = repository.DeleteFilm(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("request completed")
}
