package actor

import (
	"encoding/json"
	"github.com/lilpipidron/vk-godeveloper-task/api/types/gender"
	"github.com/lilpipidron/vk-godeveloper-task/db/actor"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct{}

func AddActorInMux(mux *http.ServeMux, repository actor.Repository) {
	mux.HandleFunc("/actor", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetActorByNameAndSurname(w, r, &repository)
		} else if r.Method == http.MethodPut {
			AddNewActor(w, r, &repository)
		} else if r.Method == http.MethodDelete {
			DeleteActorByID(w, r, &repository)
		} else if r.Method == http.MethodPost {
			ChangeInformationAboutActor(w, r, &repository)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			log.Println("Method not allowed", r.Method)
		}
	})
}

func GetActorByNameAndSurname(w http.ResponseWriter, r *http.Request, repository *actor.Repository) {
	log.Println("request: get actor by name and surname")
	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	surname := queryParams.Get("surname")
	actor, err := repository.FindActorsByNameAndSurname(name, surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	actorJSON, err := json.Marshal(actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(actorJSON)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("request completed")

}

func AddNewActor(w http.ResponseWriter, r *http.Request, repository *actor.Repository) {
	log.Println("request: put new actor")
	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	surname := queryParams.Get("surname")
	actorGender := gender.Gender(queryParams.Get("gender"))
	dateOfBirth, err := time.Parse("2006-01-02", queryParams.Get("dateOfBirth"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	err = repository.AddNewActor(name, surname, actorGender, dateOfBirth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("request completed")
}

func DeleteActorByID(w http.ResponseWriter, r *http.Request, repository *actor.Repository) {
	log.Println("request: delete actor by id")
	queryParams := r.URL.Query()
	id, err := strconv.ParseInt(queryParams.Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	err = repository.DeleteActor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("request completed")
}

func ChangeInformationAboutActor(w http.ResponseWriter, r *http.Request, repository *actor.Repository) {
	log.Println("request: post information about actor")
	queryParams := r.URL.Query()
	id, err := strconv.ParseInt(queryParams.Get("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	name := queryParams.Get("name")
	surname := queryParams.Get("surname")
	actorGender := gender.Gender(queryParams.Get("gender"))
	dateOfBirthString := queryParams.Get("dateOfBirth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	if name != "" {
		err = repository.ChangeActorName(id, name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}
	}
	if surname != "" {
		err = repository.ChangeActorSurname(id, surname)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}
	}
	if actorGender != "" {
		err = repository.ChangeActorGender(id, actorGender)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}
	}
	if dateOfBirthString != "" {
		dateOfBirth, err := time.Parse("2006-01-02", dateOfBirthString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}
		err = repository.ChangeActorDateOfBirth(id, dateOfBirth)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println(err)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	log.Println("request completed")
}
