package actorFilm

import (
	"database/sql"
	"log"
	"strings"

	"github.com/lilpipidron/vk-godeveloper-task/db/actor"
)

type Repository struct {
	DB *sql.DB
}

func NewActorFilmRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (repository *Repository) AddNewFilmAndActors(filmID int64, actors []string) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	for _, actorInfo := range actors {
		splitActor := strings.Split(actorInfo, " ")
		actorRepository := actor.NewActorRepository(repository.DB)
		actors, err := actorRepository.FindActorsByNameAndSurname(splitActor[0], splitActor[1])
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return err
			}
			return err
		}
		actorID := actors[0].ID

		query := "INSERT INTO actor_film (actor_id, film_id) VALUES ($1, $2)"
		_, err = repository.DB.Exec(query, actorID, filmID)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return err
			}
			return err
		}
	}
	return nil
}

func (repository *Repository) GetAllActorsID(filmID int64) ([]int64, error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return nil, err
	}
	query := "SELECT actor_id FROM actor_film WHERE film_id = $1"
	rows, err := repository.DB.Query(query, filmID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)
	var actorsID []int64
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return nil, err
			}
			return nil, err
		}
		actorsID = append(actorsID, id)
	}
	return actorsID, nil
}

func (repository *Repository) DeleteFilm(filmID int64) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	query := "DELETE FROM actor_film WHERE film_id = $1"
	_, err = repository.DB.Exec(query, filmID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func (repository *Repository) DeleteActor(actorID int64) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	query := "DELETE FROM actor_film WHERE actor_id = $1"
	_, err = repository.DB.Exec(query, actorID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	return nil
}
