package actorFilm

import (
	"database/sql"
	"github.com/lilpipidron/vk-godeveloper-task/db/actor"
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
