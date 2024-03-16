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

func (repository *Repository) AddNewActor(name, surname string, gender gender.Gender, dateOfBirth time.Time) error {
	query := "INSERT INTO actors (name, surname, gender, date_of_birth) VALUES ($1, $2, $3, $4)"
	_, err := repository.Driver.Exec(query, name, surname, gender, dateOfBirth)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) FindActorsByNameAndSurname(name, surname string) ([]*actor.Actor, error) {
	query := "SELECT * FROM actors WHERE name LIKE '%' || $1 || '%' AND surname LIKE '%' || $2 || '%'"
	rows, err := repository.Driver.Query(query, name, surname)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var actors []*actor.Actor
	for rows.Next() {
		a := &actor.Actor{}
		err := rows.Scan(&a.ID, &a.Name, &a.Surname, &a.Gender, &a.DateOfBirth)
		if err != nil {
			return nil, err
		}
		actors = append(actors, a)
	}

	return actors, nil
}

func (repository *Repository) DeleteActor(actorID int64) error {
	query := "DELETE FROM actors WHERE actor_id = $1"
	_, err := repository.Driver.Exec(query, actorID)
	if err != nil {
		return err
	}
	return nil

}

func (repository *Repository) ChangeActorName(actorID int64, name string) error {
	query := "UPDATE actors set name = $1 where actor_id = $2"
	_, err := repository.Driver.Exec(query, name, actorID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) ChangeActorSurname(actorID int64, surname string) error {
	query := "UPDATE actors set surname = $1 where actor_id = $2"
	_, err := repository.Driver.Exec(query, surname, actorID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) ChangeActorGender(actorID int64, gender gender.Gender) error {
	query := "UPDATE actors set gender = $1 where actor_id = $2"
	_, err := repository.Driver.Exec(query, gender, actorID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) ChangeActorDateOfBirth(actorID int64, dateOfBirth time.Time) error {
	query := "UPDATE actors set date_of_birth = $1 where actor_id = $2"
	_, err := repository.Driver.Exec(query, dateOfBirth, actorID)
	if err != nil {
		return err
	}
	return nil
}
