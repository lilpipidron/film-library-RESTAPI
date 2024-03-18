package actor

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/lilpipidron/vk-godeveloper-task/api/types/gender"
	"github.com/lilpipidron/vk-godeveloper-task/db/actor"
	"github.com/lilpipidron/vk-godeveloper-task/db/actorFilm"
)

type Handler interface {
	Handler(w http.ResponseWriter, r *http.Request) int
	GetActorByNameAndSurname(w http.ResponseWriter, r *http.Request) int
	AddNewActor(w http.ResponseWriter, r *http.Request) int
	DeleteActorByID(w http.ResponseWriter, r *http.Request) int
	ChangeInformationAboutActor(w http.ResponseWriter, r *http.Request) int
}

type Repository struct {
	repository actor.Repository
}

func NewActorRepository(repository actor.Repository) *Repository {
	return &Repository{repository: repository}
}

func (actorRepository *Repository) Handler(w http.ResponseWriter, r *http.Request) int {
	switch r.Method {
	case http.MethodGet:
		return actorRepository.GetActorByNameAndSurname(w, r)
	case http.MethodPut:
		return actorRepository.AddNewActor(w, r)
	case http.MethodDelete:
		return actorRepository.DeleteActorByID(w, r)
	case http.MethodPost:
		return actorRepository.ChangeInformationAboutActor(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Println("Method not allowed", r.Method)

		return http.StatusMethodNotAllowed
	}
}

func (actorRepository *Repository) GetActorByNameAndSurname(w http.ResponseWriter, r *http.Request) int {
	repository := actorRepository.repository
	log.Println("request: get actor by name and surname")
	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	surname := queryParams.Get("surname")
	actor, err := repository.FindActorsByNameAndSurname(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}

	actorJSON, err := json.Marshal(actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(actorJSON)
	if err != nil {
		log.Println(err)
		return http.StatusBadRequest
	}

	log.Println("request completed")
	return http.StatusOK
}

func (actorRepository *Repository) AddNewActor(w http.ResponseWriter, r *http.Request) int {
	repository := actorRepository.repository
	log.Println("request: put new actor")
	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	surname := queryParams.Get("surname")
	actorGender := gender.Gender(queryParams.Get("gender"))
	dateOfBirth, err := time.Parse("2006-01-02", queryParams.Get("dateOfBirth"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}
	err = repository.AddNewActor(name, surname, actorGender, dateOfBirth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}
	w.WriteHeader(http.StatusOK)
	log.Println("request completed")
	return http.StatusOK
}

func (actorRepository *Repository) DeleteActorByID(w http.ResponseWriter, r *http.Request) int {
	repository := actorRepository.repository
	log.Println("request: delete actor by id")
	queryParams := r.URL.Query()
	id, err := strconv.ParseInt(queryParams.Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}
	actorFilmRepository := actorFilm.NewActorFilmRepository(repository.DB)
	err = actorFilmRepository.DeleteActor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}
	err = repository.DeleteActor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}
	w.WriteHeader(http.StatusOK)
	log.Println("request completed")
	return http.StatusOK
}

func (actorRepository *Repository) ChangeInformationAboutActor(w http.ResponseWriter, r *http.Request) int {
	repository := actorRepository.repository
	log.Println("request: post information about actor")
	queryParams := r.URL.Query()
	id, err := strconv.ParseInt(queryParams.Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}
	name := queryParams.Get("name")
	surname := queryParams.Get("surname")
	actorGender := gender.Gender(queryParams.Get("gender"))
	dateOfBirthString := queryParams.Get("dateOfBirth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return http.StatusBadRequest
	}
	if name != "" {
		err = repository.ChangeActorName(id, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return http.StatusBadRequest
		}
	}
	if surname != "" {
		err = repository.ChangeActorSurname(id, surname)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return http.StatusBadRequest
		}
	}
	if actorGender != "" {
		err = repository.ChangeActorGender(id, actorGender)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return http.StatusBadRequest
		}
	}
	if dateOfBirthString != "" {
		dateOfBirth, err := time.Parse("2006-01-02", dateOfBirthString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return http.StatusBadRequest
		}
		err = repository.ChangeActorDateOfBirth(id, dateOfBirth)
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
