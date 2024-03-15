package film

import (
	"database/sql"
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
	films, err := repository.FindFilmByTitle(title)
	if err != nil {
		return err
	}
	err = filmAndActors.AddNewFilmAndActors(films[0].ID, actors)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) FindFilmByTitle(title string) ([]*film.Film, error) {
	query := "SELECT * FROM films WHERE title LIKE '%' || $1 || '%'"
	rows, err := repository.Driver.Query(query, title)
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

func (repository *Repository) DeleteFilm(filmID int) error {
	query := "DELETE FROM films WHERE film_id = $1"
	_, err := repository.Driver.Exec(query, filmID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) ChangeFilmTitle(filmID int, title string) error {
	query := "UPDATE films set title = $1 where film_id = $2"
	_, err := repository.Driver.Exec(query, title, filmID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) ChangeFilmDescription(filmID int, description string) error {
	query := "UPDATE films set description = $1 where film_id = $2"
	_, err := repository.Driver.Exec(query, description, filmID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) ChangeFilmReleaseDate(filmID int, releaseDate time.Time) error {
	query := "UPDATE films set release_date = $1 where film_id = $2"
	_, err := repository.Driver.Exec(query, releaseDate, filmID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) ChangeFilmRating(filmID int, rating float32) error {
	query := "UPDATE films set rating = $1 where film_id = $2"
	_, err := repository.Driver.Exec(query, rating, filmID)
	if err != nil {
		return err
	}
	return nil
}
