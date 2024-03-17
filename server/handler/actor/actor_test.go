package actor

import (
	mock_actor_request "github.com/lilpipidron/vk-godeveloper-task/mocks/server/handler/actor"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestActorHandler_GetActor(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	req := httptest.NewRequest(http.MethodGet, "/actor?name=test&surname=test", nil)
	w := httptest.NewRecorder()
	actorRequest := mock_actor_request.NewMockActorHandler(c)

	actorRequest.EXPECT().Handler(w, req).Return(http.StatusOK)

	status := actorRequest.Handler(w, req)
	if status != http.StatusOK {
		t.Errorf("Expected status ok, got %v", http.StatusText(status))
	}
}

func TestActorHandler_deleteActor(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	req := httptest.NewRequest(http.MethodDelete, "/actor?id=1", nil)
	w := httptest.NewRecorder()
	actorRequest := mock_actor_request.NewMockActorHandler(c)

	actorRequest.EXPECT().Handler(w, req).Return(http.StatusOK)

	status := actorRequest.Handler(w, req)
	if status != http.StatusOK {
		t.Errorf("Expected status ok, got %v", http.StatusText(status))
	}
}

func TestActorHandler_methodNotAllowed(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	req := httptest.NewRequest(http.MethodConnect, "/actor?id=1", nil)
	w := httptest.NewRecorder()
	actorRequest := mock_actor_request.NewMockActorHandler(c)

	actorRequest.EXPECT().Handler(w, req).Return(http.StatusMethodNotAllowed)

	status := actorRequest.Handler(w, req)
	if status != http.StatusMethodNotAllowed {
		t.Errorf("Expected status method not allowed, got %v", http.StatusText(status))
	}
}

func TestActorHandler_badRequest(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	req := httptest.NewRequest(http.MethodDelete, "/actor?id=w", nil)
	w := httptest.NewRecorder()
	actorRequest := mock_actor_request.NewMockActorHandler(c)

	actorRequest.EXPECT().Handler(w, req).Return(http.StatusBadRequest)

	status := actorRequest.Handler(w, req)
	if status != http.StatusBadRequest {
		t.Errorf("Expected status bad request, got %v", http.StatusText(status))
	}
}
