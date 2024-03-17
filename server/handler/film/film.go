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

type FilmHandler interface {
	Handler(w http.ResponseWriter, r *http.Request) int
	GetMediator(w http.ResponseWriter, r *http.Request) int
	GetFilmByTitleOrActorNameAndSurname(w http.ResponseWriter, title, actor string) int
	GetActors(w http.ResponseWriter, title string) int
	GetAllFilms(w http.ResponseWriter, queryParams url.Values) int
	AddNewFilm(w http.ResponseWriter, r *http.Request) int
	DeleteFilmByID(w http.ResponseWriter, r *http.Request) int
	ChangeInformationAboutFilm(w http.ResponseWriter, r *http.Request) int
}

type FilmRepository struct {
	repository film.Repository
}

func NewFilmRepository(repository film.Repository) *FilmRepository {
	return &FilmRepository{repository: repository}
}

func (filmRepository *FilmRepository) Handler(w http.ResponseWriter, r *http.Request) int {
	if r.Method == http.MethodGet {
		return filmRepository.GetMediator(w, r)
	} else if r.Method == http.MethodPut {
		return filmRepository.AddNewFilm(w, r)
	} else if r.Method == http.MethodDelete {
		return filmRepository.DeleteFilmByID(w, r)
	} else if r.Method == http.MethodPost {
		return filmRepository.ChangeInformationAboutFilm(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Println("Method not allowed", r.Method)
		return http.StatusMethodNotAllowed
	}
}

func (filmRepository *FilmRepository) GetMediator(w http.ResponseWriter, r *http.Request) int {
	queryParams := r.URL.Query()
	title := queryParams.Get("title")
	actor := queryParams.Get("actor")
	if title == "" && actor == "" {
		return filmRepository.GetAllFilms(w, queryParams)

	}

	if actor == "all" {
		return filmRepository.GetActors(w, title)
	}
	return filmRepository.GetFilmByTitleOrActorNameAndSurname(w, title, actor)
}

func (filmRepository *FilmRepository) GetFilmByTitleOrActorNameAndSurname(w http.ResponseWriter, title, actor string) int {
	repository := filmRepository.repository
	log.Println("request: get film by title or actor name and surname")
	films, err := repository.FindFilmByTitleOrActorName(title, actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}
	filmsJSON, err := json.Marshal(films)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(filmsJSON)
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest
	}
	log.Println("request completed")
	return http.StatusOK
}

func (filmRepository *FilmRepository) GetActors(w http.ResponseWriter, title string) int {
	repository := filmRepository.repository
	log.Println("request: get film's actors")
	actors, err := repository.FindAllActors(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}
	actorsJSON, err := json.Marshal(actors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(actorsJSON)
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest
	}

	log.Println("request completed")
	return http.StatusOK
}

func (filmRepository *FilmRepository) GetAllFilms(w http.ResponseWriter, queryParams url.Values) int {
	repository := filmRepository.repository
	log.Println("request: get all films and sort")
	films, err := repository.GetAllFilms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
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
		return http.StatusBadRequest
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(filmsJSON)
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest
	}

	log.Println("request completed")
	return http.StatusOK
}

func (filmRepository *FilmRepository) AddNewFilm(w http.ResponseWriter, r *http.Request) int {
	repository := filmRepository.repository
	log.Println("request: put new film")
	queryParams := r.URL.Query()
	title := queryParams.Get("title")
	description := queryParams.Get("description")
	releaseDate, err := time.Parse("2006-01-02", queryParams.Get("releaseDate"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}
	rating, err := strconv.ParseFloat(queryParams.Get("rating"), 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}
	actors := strings.Split(queryParams.Get("actors"), ",")
	err = repository.AddNewFilm(title, description, releaseDate, float32(rating), actors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}
	w.WriteHeader(http.StatusOK)
	log.Println("request completed")
	return http.StatusOK
}

func (filmRepository *FilmRepository) DeleteFilmByID(w http.ResponseWriter, r *http.Request) int {
	repository := filmRepository.repository
	log.Println("request: delete film by ID")
	queryParams := r.URL.Query()
	id, err := strconv.ParseInt(queryParams.Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}

	actorFilmRepository := actorFilm.NewActorFilmRepository(repository.DB)
	err = actorFilmRepository.DeleteFilm(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}

	err = repository.DeleteFilm(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}

	w.WriteHeader(http.StatusOK)
	log.Println("request completed")
	return http.StatusOK
}

func (filmRepository *FilmRepository) ChangeInformationAboutFilm(w http.ResponseWriter, r *http.Request) int {
	repository := filmRepository.repository
	log.Println("request: post information about film")
	queryParams := r.URL.Query()
	id, err := strconv.ParseInt(queryParams.Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}
	title := queryParams.Get("title")
	description := queryParams.Get("description")
	rating, err := strconv.ParseFloat(queryParams.Get("rating"), 32)
	releaseDate := queryParams.Get("releaseDate")
	if err != nil {
		rating = -1
	}
	if title != "" {
		err = repository.ChangeFilmTitle(id, title)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return http.StatusBadRequest
		}
	}
	if description != "" {
		err = repository.ChangeFilmDescription(id, description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return http.StatusBadRequest
		}
	}
	if rating != -1 {
		err = repository.ChangeFilmRating(id, float32(rating))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return http.StatusBadRequest
		}
	}
	if releaseDate != "" {
		releaseDateTime, err := time.Parse("2006-01-02", releaseDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return http.StatusBadRequest
		}
		err = repository.ChangeFilmReleaseDate(id, releaseDateTime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return http.StatusBadRequest
		}
	}
	w.WriteHeader(http.StatusOK)
	log.Println("request completed")
	return http.StatusOK
}
