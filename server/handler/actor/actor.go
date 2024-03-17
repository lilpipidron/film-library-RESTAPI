package actor

import (
	"encoding/json"
	"github.com/lilpipidron/vk-godeveloper-task/api/types/gender"
	"github.com/lilpipidron/vk-godeveloper-task/db/actor"
	"github.com/lilpipidron/vk-godeveloper-task/db/actorFilm"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ActorHandler interface {
	Handler(w http.ResponseWriter, r *http.Request) int
	GetActorByNameAndSurname(w http.ResponseWriter, r *http.Request) int
	AddNewActor(w http.ResponseWriter, r *http.Request) int
	DeleteActorByID(w http.ResponseWriter, r *http.Request) int
	ChangeInformationAboutActor(w http.ResponseWriter, r *http.Request) int
}

type ActorRepository struct {
	repository actor.Repository
}

func NewActorRepository(repository actor.Repository) *ActorRepository {
	return &ActorRepository{repository: repository}
}

func (actorRepository *ActorRepository) Handler(w http.ResponseWriter, r *http.Request) int {
	if r.Method == http.MethodGet {
		return actorRepository.GetActorByNameAndSurname(w, r)
	} else if r.Method == http.MethodPut {
		return actorRepository.AddNewActor(w, r)
	} else if r.Method == http.MethodDelete {
		return actorRepository.DeleteActorByID(w, r)
	} else if r.Method == http.MethodPost {
		return actorRepository.ChangeInformationAboutActor(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Println("Method not allowed", r.Method)
		return http.StatusMethodNotAllowed
	}
}

func (actorRepository *ActorRepository) GetActorByNameAndSurname(w http.ResponseWriter, r *http.Request) int {
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

func (actorRepository *ActorRepository) AddNewActor(w http.ResponseWriter, r *http.Request) int {
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

func (actorRepository *ActorRepository) DeleteActorByID(w http.ResponseWriter, r *http.Request) int {
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

func (actorRepository *ActorRepository) ChangeInformationAboutActor(w http.ResponseWriter, r *http.Request) int {
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
