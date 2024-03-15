package actorFilm

import (
	"database/sql"
	"github.com/lilpipidron/vk-godeveloper-task/db/actor"
	"log"
	"strings"
)

type Repository struct {
	Driver *sql.DB
}

func (repository *Repository) AddNewFilmAndActors(filmID int64, actors []string) error {

	for _, actorInfo := range actors {
		splitActor := strings.Split(actorInfo, " ")
		actorRepository := &actor.Repository{Driver: repository.Driver}
		actors, err := actorRepository.FindActorsByNameAndSurname(splitActor[0], splitActor[1])
		if err != nil {
			return err
		}
		actorID := actors[0].ID

		query := "INSERT INTO actor_film (actor_id, film_id) VALUES ($1, $2)"
		_, err = repository.Driver.Exec(query, actorID, filmID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repository *Repository) GetAllActorsID(filmID int64) ([]int64, error) {
	query := "SELECT actor_id FROM actor_film WHERE film_id = $1"
	rows, err := repository.Driver.Query(query, filmID)
	if err != nil {
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
		err := rows.Scan(id)
		if err != nil {
			return nil, err
		}
		actorsID = append(actorsID, id)
	}
	return actorsID, nil
}
