package actor

import (
	"database/sql"
	"github.com/lilpipidron/vk-godeveloper-task/api/types/actor"
	"github.com/lilpipidron/vk-godeveloper-task/api/types/gender"
	"log"
	"time"
)

type ActorRepository interface {
	AddNewActor(name, surname string, gender gender.Gender, dateOfBirth time.Time) error
	FindActorsByNameAndSurname(name, surname string) ([]*actor.Actor, error)
	DeleteActor(actorID int64) error
	ChangeActorName(actorID int64, name string) error
	ChangeActorSurname(actorID int64, surname string) error
	ChangeActorGender(actorID int64, gender gender.Gender) error
	ChangeActorDateOfBirth(actorID int64, dateOfBirth time.Time) error
}

type Repository struct {
	DB *sql.DB
}

func NewActorRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (repository *Repository) AddNewActor(name, surname string, gender gender.Gender, dateOfBirth time.Time) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	query := "INSERT INTO actors (name, surname, gender, date_of_birth) VALUES ($1, $2, $3, $4)"
	_, err = repository.DB.Exec(query, name, surname, gender, dateOfBirth)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func (repository *Repository) FindActorsByNameAndSurname(name, surname string) ([]*actor.Actor, error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return nil, err
	}
	query := "SELECT * FROM actors WHERE name LIKE '%' || $1 || '%' AND surname LIKE '%' || $2 || '%'"
	rows, err := repository.DB.Query(query, name, surname)
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
			err := tx.Rollback()
			if err != nil {
				return
			}
			log.Println(err)
		}
	}(rows)

	var actors []*actor.Actor
	for rows.Next() {
		a := &actor.Actor{}
		err := rows.Scan(&a.ID, &a.Name, &a.Surname, &a.Gender, &a.DateOfBirth)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return nil, err
			}
			return nil, err
		}
		actors = append(actors, a)
	}

	return actors, nil
}

func (repository *Repository) DeleteActor(actorID int64) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	query := "DELETE FROM actors WHERE actor_id = $1"
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

func (repository *Repository) ChangeActorName(actorID int64, name string) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE actors set name = $1 where actor_id = $2"
	_, err = repository.DB.Exec(query, name, actorID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func (repository *Repository) ChangeActorSurname(actorID int64, surname string) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE actors set surname = $1 where actor_id = $2"
	_, err = repository.DB.Exec(query, surname, actorID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func (repository *Repository) ChangeActorGender(actorID int64, gender gender.Gender) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE actors set gender = $1 where actor_id = $2"
	_, err = repository.DB.Exec(query, gender, actorID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func (repository *Repository) ChangeActorDateOfBirth(actorID int64, dateOfBirth time.Time) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE actors set date_of_birth = $1 where actor_id = $2"
	_, err = repository.DB.Exec(query, dateOfBirth, actorID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	return nil
}
