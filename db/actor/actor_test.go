package actor

import (
	"testing"
	"time"

	actor2 "github.com/lilpipidron/vk-godeveloper-task/api/types/actor"
	"github.com/lilpipidron/vk-godeveloper-task/api/types/gender"
	mock_actor "github.com/lilpipidron/vk-godeveloper-task/mocks/db/actor"
	"go.uber.org/mock/gomock"
)

func TestActorRepository_addNewActor(t *testing.T) {
	actor := actor2.Actor{
		Name:        "name",
		Surname:     "surname",
		Gender:      gender.Male,
		DateOfBirth: time.Now(),
	}

	c := gomock.NewController(t)
	defer c.Finish()

	db := mock_actor.NewMockActorRepository(c)

	db.EXPECT().AddNewActor(actor.Name, actor.Surname, actor.Gender, actor.DateOfBirth).Return(nil)

	err := db.AddNewActor(actor.Name, actor.Surname, actor.Gender, actor.DateOfBirth)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestActorRepository_addAndDeleteActor(t *testing.T) {
	actor := actor2.Actor{
		Name:        "name",
		Surname:     "surname",
		Gender:      gender.Male,
		DateOfBirth: time.Now(),
	}

	c := gomock.NewController(t)
	defer c.Finish()

	db := mock_actor.NewMockActorRepository(c)

	db.EXPECT().AddNewActor(actor.Name, actor.Surname, actor.Gender, actor.DateOfBirth).Return(nil)

	db.EXPECT().DeleteActor(gomock.Any()).Return(nil)

	err := db.AddNewActor(actor.Name, actor.Surname, actor.Gender, actor.DateOfBirth)
	if err != nil {
		t.Errorf("Failed to add actor: %v", err)
	}

	err = db.DeleteActor(1)
	if err != nil {
		t.Errorf("Failed to delete actor: %v", err)
	}
}

func TestActorRepository_addAndFindActor(t *testing.T) {
	actor := actor2.Actor{
		Name:        "name",
		Surname:     "surname",
		Gender:      gender.Male,
		DateOfBirth: time.Now(),
	}

	c := gomock.NewController(t)
	defer c.Finish()

	db := mock_actor.NewMockActorRepository(c)

	db.EXPECT().AddNewActor(actor.Name, actor.Surname, actor.Gender, actor.DateOfBirth).Return(nil)

	db.EXPECT().FindActorsByNameAndSurname(actor.Name, actor.Surname).
		Return([]*actor2.Actor{&actor}, nil)

	err := db.AddNewActor(actor.Name, actor.Surname, actor.Gender, actor.DateOfBirth)
	if err != nil {
		t.Errorf("Failed to add actor: %v", err)
	}

	actors, err := db.FindActorsByNameAndSurname(actor.Name, actor.Surname)
	if err != nil {
		t.Errorf("Failed to find actor: %v", err)
	}

	if len(actors) != 1 {
		t.Errorf("Expected to find 1 actor, found %d", len(actors))
	}
}

func TestActorRepository_addActorChangeNameAndFindActor(t *testing.T) {
	actor := actor2.Actor{
		Name:        "name",
		Surname:     "surname",
		Gender:      gender.Male,
		DateOfBirth: time.Now(),
	}
	newActor := actor2.Actor{
		Name:        "new",
		Surname:     "surname",
		Gender:      gender.Male,
		DateOfBirth: time.Now(),
	}
	newName := "new"
	c := gomock.NewController(t)
	defer c.Finish()

	db := mock_actor.NewMockActorRepository(c)

	db.EXPECT().AddNewActor(actor.Name, actor.Surname, actor.Gender, actor.DateOfBirth).Return(nil)

	db.EXPECT().ChangeActorName(int64(1), newName).Return(nil)

	db.EXPECT().FindActorsByNameAndSurname(newName, actor.Surname).
		Return([]*actor2.Actor{&newActor}, nil)

	err := db.AddNewActor(actor.Name, actor.Surname, actor.Gender, actor.DateOfBirth)
	if err != nil {
		t.Errorf("Failed to add actor: %v", err)
	}

	err = db.ChangeActorName(1, newName)
	if err != nil {
		t.Errorf("Failed to change actor's name: %v", err)
	}

	actors, err := db.FindActorsByNameAndSurname(newName, actor.Surname)
	if err != nil {
		t.Errorf("Failed to find actor: %v", err)
	}

	if len(actors) != 1 {
		t.Errorf("Expected to find 1 actor, found %d", len(actors))
	}
}
