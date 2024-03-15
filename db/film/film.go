package film

import (
	"database/sql"
	"github.com/lilpipidron/vk-godeveloper-task/api/types/actor"
	"github.com/lilpipidron/vk-godeveloper-task/api/types/film"
	"github.com/lilpipidron/vk-godeveloper-task/db/actorFilm"
	"log"
	"time"
)

type Repository struct {
	Driver *sql.DB
}

func (repository *Repository) AddNewFilm(title, description string, releaseDate time.Time, rating float32, actors []string) error {
	query := "INSERT INTO films (title, description, release_date, rating) VALUES ($1,$2,$3,$4)"
	_, err := repository.Driver.Exec(query, title, description, releaseDate, rating)
	if err != nil {
		return err
	}

	filmAndActors := &actorFilm.Repository{Driver: repository.Driver}
	films, err := repository.FindFilmByTitleOrActorName(title, "", "")
	if err != nil {
		return err
	}
	err = filmAndActors.AddNewFilmAndActors(films[0].ID, actors)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) FindFilmByTitleOrActorName(title, name, surname string) ([]*film.Film, error) {
	query := "SELECT film_title, film_description, film_release_date, film_rating FROM actor_film_view WHERE film_title LIKE '%' || $1 || '%' OR actor_name LIKE '%' || $2 || '%' OR actor_surname LIKE '%' || $3 || '%' "
	rows, err := repository.Driver.Query(query, title, name, surname)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var films []*film.Film
	for rows.Next() {
		f := &film.Film{}
		err := rows.Scan(&f.ID, &f.Title, &f.Description, &f.ReleaseDate, &f.Rating)
		if err != nil {
			return nil, err
		}
		films = append(films, f)
	}
	return films, nil
}

func (repository *Repository) DeleteFilm(filmID int64) error {
	query := "DELETE FROM films WHERE film_id = $1"
	_, err := repository.Driver.Exec(query, filmID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) ChangeFilmTitle(filmID int64, title string) error {
	query := "UPDATE films set title = $1 where film_id = $2"
	_, err := repository.Driver.Exec(query, title, filmID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) ChangeFilmDescription(filmID int64, description string) error {
	query := "UPDATE films set description = $1 where film_id = $2"
	_, err := repository.Driver.Exec(query, description, filmID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) ChangeFilmReleaseDate(filmID int64, releaseDate time.Time) error {
	query := "UPDATE films set release_date = $1 where film_id = $2"
	_, err := repository.Driver.Exec(query, releaseDate, filmID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) ChangeFilmRating(filmID int64, rating float32) error {
	query := "UPDATE films set rating = $1 where film_id = $2"
	_, err := repository.Driver.Exec(query, rating, filmID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) FindAllActors(filmID int64) ([]*actor.ActorWithFilms, error) {
	filmAndActors := &actorFilm.Repository{Driver: repository.Driver}
	actorsID, err := filmAndActors.GetAllActorsID(filmID)
	if err != nil {
		return nil, err
	}
	var actorAndFilms []*actor.ActorWithFilms
	for i, actorID := range actorsID {
		query := "SELECT actor_id, actor_name, actor_surname, actor_gender, actor_date_of_birth FROM actor_film_view WHERE actor_id = $1"
		rows, err := repository.Driver.Query(query, actorID)
		if err != nil {
			return nil, err
		}
		a := &actor.ActorWithFilms{}
		err = rows.Scan(&a.ID, a.Name, a.Surname, a.Gender, a.DateOfBirth)
		if err != nil {
			return nil, err
		}
		actorAndFilms = append(actorAndFilms, a)

		query = "SELECT film_id, film_title, film_description, film_release_date, film_rating FROM actor_film_view WHERE actor_id = $1"
		rows, err = repository.Driver.Query(query, actorID)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			f := &film.Film{}
			err := rows.Scan(&f.ID, &f.Title, &f.Description, &f.ReleaseDate, &f.Rating)
			if err != nil {
				return nil, err
			}
			actorAndFilms[i].Films = append(actorAndFilms[i].Films, *f)
		}
	}
	return actorAndFilms, nil
}
