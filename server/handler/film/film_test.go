package film

import (
	mock_film_request "github.com/lilpipidron/vk-godeveloper-task/mocks/server/handler/film"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFilmHandler_GetFilmByTitle(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	req := httptest.NewRequest(http.MethodGet, "/film?title=test", nil)
	w := httptest.NewRecorder()
	filmRequest := mock_film_request.NewMockFilmHandler(c)

	filmRequest.EXPECT().Handler(w, req).Return(http.StatusOK)

	status := filmRequest.Handler(w, req)
	if status != http.StatusOK {
		t.Errorf("Expected status ok, got %v", http.StatusText(status))
	}
}

func TestFilmHandler_deleteFilm(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	req := httptest.NewRequest(http.MethodDelete, "/film?id=1", nil)
	w := httptest.NewRecorder()
	filmRequest := mock_film_request.NewMockFilmHandler(c)

	filmRequest.EXPECT().Handler(w, req).Return(http.StatusOK)

	status := filmRequest.Handler(w, req)
	if status != http.StatusOK {
		t.Errorf("Expected status ok, got %v", http.StatusText(status))
	}
}

func TestFilmHandler_methodNotAllowed(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	req := httptest.NewRequest(http.MethodConnect, "/film?id=1", nil)
	w := httptest.NewRecorder()
	filmRequest := mock_film_request.NewMockFilmHandler(c)

	filmRequest.EXPECT().Handler(w, req).Return(http.StatusMethodNotAllowed)

	status := filmRequest.Handler(w, req)
	if status != http.StatusMethodNotAllowed {
		t.Errorf("Expected status method not allowed, got %v", http.StatusText(status))
	}
}

func TestFilmHandler_badRequest(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	req := httptest.NewRequest(http.MethodDelete, "/film?id=w", nil)
	w := httptest.NewRecorder()
	filmRequest := mock_film_request.NewMockFilmHandler(c)

	filmRequest.EXPECT().Handler(w, req).Return(http.StatusBadRequest)

	status := filmRequest.Handler(w, req)
	if status != http.StatusBadRequest {
		t.Errorf("Expected status bad request, got %v", http.StatusText(status))
	}
}
