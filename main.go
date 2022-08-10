package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/thedevsaddam/renderer"
)

const DBName = "todos.db"
var rnd *renderer.Render
//var DB, err = gorm.Open(sqlite.Open(DBName), &gorm.Config{})

type todos struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	//Done bool `json:"done"`
}

type Server struct {
	DB *gorm.DB
}

func (s *Server) CreateTodo(w http.ResponseWriter, r *http.Request) { //POST
	w.Header().Set("Content-Type", "application/json")
	var new todos
	json.NewDecoder(r.Body).Decode(&new)
	s.DB.Create(&new)
	json.NewEncoder(w).Encode(new)
	w.WriteHeader(http.StatusOK)

}

func (s *Server) GetTodos(w http.ResponseWriter, r *http.Request) { //GET all
	_, err := s.DB.Model(&todos{}).Where("ID > ?", "0").Rows()

	if err != nil {
		panic("error parsing data")
	}
	w.Header().Set("Content-Type", "application/json")

	var new []todos
	s.DB.Find(&new)
	json.NewEncoder(w).Encode(new)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) GetTodoByID(w http.ResponseWriter, r *http.Request) { //GET by ID
	vars := mux.Vars(r)
	id := vars["id"]
	var res todos
	out := s.DB.First(&res, id)
	if out.Error != nil {
		http.Error(w, out.Error.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - can't find this task"))
		return
	}
	data, err := json.Marshal(res)
	if err != nil {
		http.Error(w, out.Error.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Error can't find this task"))
		return
	}
	w.Write(data)
	w.Write([]byte("\n"))

}

func (s *Server) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	var new todos
	res := s.DB.First(&new, param["ID"])
	json.NewDecoder(r.Body).Decode(&new)

	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - can't find this task"))
		return
	}

	s.DB.Save(&new)
	json.NewEncoder(w).Encode(&new)
	w.WriteHeader(http.StatusOK)

}

func (s *Server) DeleteTodo(w http.ResponseWriter, r *http.Request) { //DELETE
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["taskId"])

	var new = todos{ID: id}
	s.DB.Find(&new)
	res := s.DB.First(&new, "ID")
	s.DB.Delete(&new)
	if res.Error != nil {
		http.Error(w, res.Error.Error(), http.StatusNotFound)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - can't find this task"))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Task deleted successfully"))
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the MY TODO List!")
	fmt.Println("Endpoint Hit: homePage")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := rnd.Template(w, http.StatusOK, []string{"client/app.tpl"}, nil)
	fmt.Errorf(err.Error())
}

func initialize() (s Server){
	rnd = renderer.New()
	var DB, err = gorm.Open(sqlite.Open(DBName), &gorm.Config{})
	if err != nil {
		fmt.Errorf("can not connect to db")
	}
	server := Server{DB}
	return server
}

func handleRequests() {
	//http.HandleFunc("/", homePage)
	var DB, err = gorm.Open(sqlite.Open(DBName), &gorm.Config{})
	server := Server{DB}
	if !(DB.Migrator().HasTable(&todos{})) {
		log.Println("table { todos } created")
		DB.Migrator().CreateTable(&todos{})
	}
	if err != nil {
		panic("can't connect to DB")
	}
	router := mux.NewRouter()
    http.HandleFunc("/", homeHandler)
	router.HandleFunc("/todo", server.GetTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", server.GetTodoByID).Methods("GET")
	router.HandleFunc("/todo", server.CreateTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", server.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo/{taskId}", server.DeleteTodo).Methods("DELETE")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	  handleRequests()
}
