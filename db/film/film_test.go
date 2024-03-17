package film

import (
	actorType "github.com/lilpipidron/vk-godeveloper-task/api/types/actor"
	filmType "github.com/lilpipidron/vk-godeveloper-task/api/types/film"
	"github.com/lilpipidron/vk-godeveloper-task/api/types/gender"
	mock_actor "github.com/lilpipidron/vk-godeveloper-task/mocks/db/actor"
	mock_film "github.com/lilpipidron/vk-godeveloper-task/mocks/db/film"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func TestFilmRepository_addNewFilm(t *testing.T) {
	film := filmType.Film{
		Title:       "title",
		Description: "description",
		ReleaseDate: time.Now(),
		Rating:      3.2,
	}

	c := gomock.NewController(t)
	defer c.Finish()

	db := mock_film.NewMockFilmRepository(c)

	db.EXPECT().AddNewFilm(film.Title, film.Description, film.ReleaseDate, film.Rating, gomock.Any()).Return(nil)

	err := db.AddNewFilm(film.Title, film.Description, film.ReleaseDate, film.Rating, []string{"John Doe"})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestFilmRepository_addFilmAndDeleteFilm(t *testing.T) {
	film := filmType.Film{
		Title:       "title",
		Description: "description",
		ReleaseDate: time.Now(),
		Rating:      3.2,
	}

	c := gomock.NewController(t)
	defer c.Finish()

	db := mock_film.NewMockFilmRepository(c)

	db.EXPECT().AddNewFilm(film.Title, film.Description, film.ReleaseDate, film.Rating, gomock.Any()).Return(nil)

	db.EXPECT().DeleteFilm(gomock.Any()).Return(nil)

	err := db.AddNewFilm(film.Title, film.Description, film.ReleaseDate, film.Rating, []string{"John Doe"})
	if err != nil {
		t.Errorf("Failed to add film: %v", err)
	}

	err = db.DeleteFilm(1)
	if err != nil {
		t.Errorf("Failed to delete film: %v", err)
	}
}

func TestFilmRepository_addAndFindFilm(t *testing.T) {
	film := filmType.Film{
		Title:       "title",
		Description: "description",
		ReleaseDate: time.Now(),
		Rating:      3.2,
	}
	c := gomock.NewController(t)
	defer c.Finish()

	db := mock_film.NewMockFilmRepository(c)

	db.EXPECT().AddNewFilm(film.Title, film.Description, film.ReleaseDate, film.Rating, gomock.Any()).Return(nil)

	db.EXPECT().FindFilmByTitleOrActorName(film.Title, "q q").Return([]*filmType.Film{&film}, nil)

	err := db.AddNewFilm(film.Title, film.Description, film.ReleaseDate, film.Rating, []string{"John Doe"})
	if err != nil {
		t.Errorf("Failed to add film: %v", err)
	}

	films, err := db.FindFilmByTitleOrActorName(film.Title, "q q")
	if err != nil {
		t.Errorf("Failed to find film: %v", err)
	}

	if len(films) != 1 {
		t.Errorf("Expected to find 1 actor, found %d", len(films))
	}
}

func TestFilmRepository_addFilmAddActorsAndFindAllActors(t *testing.T) {
	film := filmType.Film{
		Title:       "title",
		Description: "description",
		ReleaseDate: time.Now(),
		Rating:      3.2,
	}
	actor1 := actorType.Actor{
		Name:        "name1",
		Surname:     "surname1",
		Gender:      gender.Male,
		DateOfBirth: time.Now(),
	}
	actor2 := actorType.Actor{
		Name:        "name2",
		Surname:     "surname2",
		Gender:      gender.Male,
		DateOfBirth: time.Now(),
	}
	actor1WithFilm := actorType.ActorWithFilms{
		Name:        "name1",
		Surname:     "surname1",
		Gender:      gender.Male,
		DateOfBirth: time.Now(),
		Films:       []filmType.Film{film},
	}
	actor2WithFilm := actorType.ActorWithFilms{
		Name:        "name2",
		Surname:     "surname2",
		Gender:      gender.Male,
		DateOfBirth: time.Now(),
		Films:       []filmType.Film{film},
	}
	c := gomock.NewController(t)
	defer c.Finish()
	dbFilms := mock_film.NewMockFilmRepository(c)
	dbActors := mock_actor.NewMockActorRepository(c)

	dbActors.EXPECT().AddNewActor(actor1.Name, actor1.Surname, actor1.Gender, actor1.DateOfBirth).Return(nil)

	dbActors.EXPECT().AddNewActor(actor2.Name, actor2.Surname, actor2.Gender, actor2.DateOfBirth).Return(nil)

	dbFilms.EXPECT().AddNewFilm(film.Title, film.Description, film.ReleaseDate, film.Rating, gomock.Any()).Return(nil)

	dbFilms.EXPECT().FindAllActors(film.Title).Return([]*actorType.ActorWithFilms{&actor1WithFilm, &actor2WithFilm}, nil)

	err := dbActors.AddNewActor(actor1.Name, actor1.Surname, actor1.Gender, actor1.DateOfBirth)
	if err != nil {
		t.Errorf("Failed to add actor1: %v", err)
	}

	err = dbActors.AddNewActor(actor2.Name, actor2.Surname, actor2.Gender, actor2.DateOfBirth)
	if err != nil {
		t.Errorf("Failed to add actorType: %v", err)
	}

	err = dbFilms.AddNewFilm(film.Title, film.Description, film.ReleaseDate, film.Rating, []string{"name1 surname1", "name2 surname2"})
	if err != nil {
		t.Errorf("Failed to add film: %v", err)
	}

	actors, err := dbFilms.FindAllActors(film.Title)
	if err != nil {
		t.Errorf("Failed to find film's actors: %v", err)
	}

	if len(actors) != 2 {
		t.Errorf("Expected to find 2 actors, found %d", len(actors))
	}
}

func TestFilmRepository_addFilmChangeTitle(t *testing.T) {
	film := filmType.Film{
		Title:       "title",
		Description: "description",
		ReleaseDate: time.Now(),
		Rating:      3.2,
	}
	newFilm := filmType.Film{
		Title:       "new title",
		Description: "description",
		ReleaseDate: time.Now(),
		Rating:      3.2,
	}
	c := gomock.NewController(t)
	defer c.Finish()

	db := mock_film.NewMockFilmRepository(c)

	db.EXPECT().AddNewFilm(film.Title, film.Description, film.ReleaseDate, film.Rating, gomock.Any()).Return(nil)

	db.EXPECT().ChangeFilmTitle(int64(1), newFilm.Title).Return(nil)

	db.EXPECT().FindFilmByTitleOrActorName(newFilm.Title, "q q").Return([]*filmType.Film{&newFilm}, nil)

	err := db.AddNewFilm(film.Title, film.Description, film.ReleaseDate, film.Rating, []string{"John Doe"})
	if err != nil {
		t.Errorf("Failed to add film: %v", err)
	}

	err = db.ChangeFilmTitle(int64(1), newFilm.Title)
	if err != nil {
		t.Errorf("Failed to change film's title: %v", err)
	}

	films, err := db.FindFilmByTitleOrActorName(newFilm.Title, "q q")
	if err != nil {
		t.Errorf("Failed to find film with new title: %v", err)
	}

	if len(films) != 1 {
		t.Errorf("Expected to find 1 actor, found %d", len(films))
	}
}
