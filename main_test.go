package main

import (
	"testing"
	"net/http"
	"net/http/httptest"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGETtodo(t *testing.T) {
	// todo := todos{	
	// 		ID : 1,
	// 		Task: "Study",			
	// }
	//server := &todos{1,"study"}
	const DBName = "todos.db"
	var DB, err = gorm.Open(sqlite.Open(DBName), &gorm.Config{})
	if err != nil {
		t.Errorf("failed to connect to db %v", err)
	}
    server := Server{DB}
	t.Run("returns all todo items", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/todo", nil)
		response := httptest.NewRecorder()
        
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")

		server.GetTodos(response, request)
		assertResponseBody(t, response.Body.String(), "20")
	})

	
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}