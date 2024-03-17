package film

import (
	"database/sql"
	"github.com/lilpipidron/vk-godeveloper-task/api/types/actor"
	"github.com/lilpipidron/vk-godeveloper-task/api/types/film"
	"github.com/lilpipidron/vk-godeveloper-task/db/actorFilm"
	"log"
	"strings"
	"time"
)

type FilmRepository interface {
	GetAllFilms() ([]*film.Film, error)
	AddNewFilm(title, description string, releaseDate time.Time, rating float32, actors []string) error
	FindFilmByTitleOrActorName(title, actor string) ([]*film.Film, error)
	DeleteFilm(filmID int64) error
	ChangeFilmTitle(filmID int64, title string) error
	ChangeFilmDescription(filmID int64, description string) error
	ChangeFilmReleaseDate(filmID int64, releaseDate time.Time) error
	ChangeFilmRating(filmID int64, rating float32) error
	FindAllActors(title string) ([]*actor.ActorWithFilms, error)
}
type Repository struct {
	DB *sql.DB
}

func NewFilmRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (repository *Repository) GetAllFilms() ([]*film.Film, error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return nil, err
	}
	query := "SELECT * FROM films"
	rows, err := repository.DB.Query(query)
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
	var films []*film.Film
	for rows.Next() {
		f := &film.Film{}
		err := rows.Scan(&f.ID, &f.Title, &f.Description, &f.ReleaseDate, &f.Rating)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return nil, err
			}
			return nil, err
		}
		films = append(films, f)
	}
	return films, nil
}

func (repository *Repository) AddNewFilm(title, description string, releaseDate time.Time, rating float32, actors []string) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	query := "INSERT INTO films (title, description, release_date, rating) VALUES ($1,$2,$3,$4)"
	_, err = repository.DB.Exec(query, title, description, releaseDate, rating)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	actorsAndFilm := actorFilm.NewActorFilmRepository(repository.DB)
	id, err := repository.findFilmIDByTitle(title)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	err = actorsAndFilm.AddNewFilmAndActors(id, actors)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func (repository *Repository) FindFilmByTitleOrActorName(title, actor string) ([]*film.Film, error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return nil, err
	}
	actorParam := strings.Split(actor, " ")
	query := "SELECT DISTINCT film_id, film_title, film_description, film_release_date, film_rating FROM actor_film_view WHERE film_title LIKE '%' || $1 || '%' OR (actor_name LIKE '%' || $2 || '%' AND actor_surname LIKE '%' || $3 || '%')"
	if actor == "" {
		actorParam = append(actorParam, "empty")
		actorParam[0] = "empty"
	}
	rows, err := repository.DB.Query(query, title, actorParam[0], actorParam[1])
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

	var films []*film.Film
	for rows.Next() {
		f := &film.Film{}
		err := rows.Scan(&f.ID, &f.Title, &f.Description, &f.ReleaseDate, &f.Rating)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return nil, err
			}
			return nil, err
		}
		films = append(films, f)
	}
	return films, nil
}

func (repository *Repository) findFilmIDByTitle(title string) (int64, error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return -1, err
	}
	query := "SELECT film_id FROM films WHERE title = $1"
	rows, err := repository.DB.Query(query, title)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return -1, err
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
	var id int64
	rows.Next()
	err = rows.Scan(&id)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return -1, err
	}
	return id, nil
}

func (repository *Repository) DeleteFilm(filmID int64) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	query := "DELETE FROM films WHERE film_id = $1"
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

func (repository *Repository) ChangeFilmTitle(filmID int64, title string) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE films set title = $1 where film_id = $2"
	_, err = repository.DB.Exec(query, title, filmID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func (repository *Repository) ChangeFilmDescription(filmID int64, description string) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE films set description = $1 where film_id = $2"
	_, err = repository.DB.Exec(query, description, filmID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func (repository *Repository) ChangeFilmReleaseDate(filmID int64, releaseDate time.Time) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE films set release_date = $1 where film_id = $2"
	_, err = repository.DB.Exec(query, releaseDate, filmID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func (repository *Repository) ChangeFilmRating(filmID int64, rating float32) error {
	tx, err := repository.DB.Begin()
	if err != nil {
		return err
	}
	query := "UPDATE films set rating = $1 where film_id = $2"
	_, err = repository.DB.Exec(query, rating, filmID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func (repository *Repository) FindAllActors(title string) ([]*actor.ActorWithFilms, error) {
	tx, err := repository.DB.Begin()
	if err != nil {
		return nil, err
	}
	filmID, err := repository.findFilmIDByTitle(title)
	actorsAndFilm := actorFilm.NewActorFilmRepository(repository.DB)
	actorsID, err := actorsAndFilm.GetAllActorsID(filmID)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	var actorAndFilms []*actor.ActorWithFilms
	for i, actorID := range actorsID {
		query := "SELECT actor_id, actor_name, actor_surname, actor_gender, actor_date_of_birth FROM actor_film_view WHERE actor_id = $1"
		rows, err := repository.DB.Query(query, actorID)
		for rows.Next() {
			if err != nil {
				err := tx.Rollback()
				if err != nil {
					return nil, err
				}
				return nil, err
			}

			a := &actor.ActorWithFilms{}
			err = rows.Scan(&a.ID, &a.Name, &a.Surname, &a.Gender, &a.DateOfBirth)
			if err != nil {
				err := tx.Rollback()
				if err != nil {
					return nil, err
				}
				return nil, err
			}
			actorAndFilms = append(actorAndFilms, a)

			query = "SELECT film_id, film_title, film_description, film_release_date, film_rating FROM actor_film_view WHERE actor_id = $1"
			rows, err = repository.DB.Query(query, actorID)
			if err != nil {
				err := tx.Rollback()
				if err != nil {
					return nil, err
				}
				return nil, err
			}

			for rows.Next() {
				f := &film.Film{}
				err := rows.Scan(&f.ID, &f.Title, &f.Description, &f.ReleaseDate, &f.Rating)
				if err != nil {
					err := tx.Rollback()
					if err != nil {
						return nil, err
					}
					return nil, err
				}

				actorAndFilms[i].Films = append(actorAndFilms[i].Films, *f)
			}
		}
	}
	return actorAndFilms, nil
}
