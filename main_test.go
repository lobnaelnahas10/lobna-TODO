package todo

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTodo(t *testing.T) {  //POST
	var server Server
	newTodo := todos{
		ID:   12,
		Task: "task",
	}
	server = initialize()
	jsonValue, _ := json.Marshal(newTodo)

	t.Run("create new todo", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/todo", bytes.NewBuffer(jsonValue))
		response := httptest.NewRecorder()
	
		server.CreateTodo(response, request)

		assertStatus(t, response.Code, http.StatusOK)

	})
}

func TestGETtodo(t *testing.T) {
	server := initialize()
	t.Run("returns all todo items", func(t *testing.T) { //GET ALL
		request := httptest.NewRequest(http.MethodGet, "/todo", nil)
		response := httptest.NewRecorder()
        
		server.GetTodos(response, request)
		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("return todo items by id", func(t *testing.T) { //GET BY ID
		request := httptest.NewRequest(http.MethodGet, "/todo/12", nil)
		response := httptest.NewRecorder()
        
		server.GetTodoByID(response, request)
		assertStatus(t, response.Code, http.StatusOK)
	})
}

func TestDeleteTodo(t *testing.T) {  //DELETE
	var server Server
	server = initialize()

	t.Run("delete existing/non-existing todo by id", func(t *testing.T) { 
		request, _ := http.NewRequest(http.MethodPost, "/todo/12", nil)
		response := httptest.NewRecorder()
	
		server.DeleteTodo(response, request)
        if(response.Code == 200){assertStatus(t, response.Code, http.StatusOK) //id exits
		}else { //id does not exist
			assertStatus(t, response.Code, http.StatusNotFound)
		}

	})
}

func TestUpdateTodo(t *testing.T) {  //PUT
	var server Server
	newTodo := todos{
		ID:   12,
		Task: "task",
	}
	server = initialize()
	jsonValue, _ := json.Marshal(newTodo)

	t.Run("create new todo", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/todo/12", bytes.NewBuffer(jsonValue))
		response := httptest.NewRecorder()
	
		server.CreateTodo(response, request)

		assertStatus(t, response.Code, http.StatusOK)

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