package actor

import (
	"database/sql"
	"github.com/lilpipidron/vk-godeveloper-task/api/types/actor"
	"github.com/lilpipidron/vk-godeveloper-task/api/types/gender"
	"log"
	"time"
)

type Repository struct {
	Driver *sql.DB
}

func (repository *Repository) FindActorByNameAndSurname(name, surname string) (*actor.Actor, error) {
	query := "SELECT * FROM actors where name = $1 and surname = $2"
	row := repository.Driver.QueryRow(query, name, surname)

	actor := &actor.Actor{}
	err := row.Scan(&actor.ID, &actor.Name, &actor.Surname, &actor.Gender, &actor.DateOfBirth)
	if err != nil {
		return nil, err
	}
	return actor, nil
}

func (repository *Repository) AddNewActor(name, surname string, gender gender.Gender, dateOfBirth time.Time) error {
	query := "INSERT INTO actors (name, surname, gender, date_of_birth) VALUES ($1, $2, $3, $4)"
	_, err := repository.Driver.Exec(query, name, surname, gender, dateOfBirth)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) DeleteActor(actorID int) error {
	query := "DELETE FROM actors where actor_id = $1"
	_, err := repository.Driver.Exec(query, actorID)
	log.Println(err)
	if err != nil {
		return err
	}
	return nil

}
