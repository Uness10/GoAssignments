package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

type Course struct {
	Code   string
	Name   string
	Credit int
}

type Student struct {
	Id      int
	Name    string
	Age     int
	Major   string
	Address Address
	Course  []Course
}

type StudentRepository interface {
	Create(student Student) (Student, error)
	GetByID(id int) (Student, error)
	Update(id int, student Student) (Student, error)
	Delete(id int) error
	GetAll() ([]Student, error)
}

type StudentsData struct {
	students map[int]Student
	maxID    int
}

func repoInit() *StudentsData {
	return &StudentsData{
		students: make(map[int]Student),
		maxID:    1,
	}
}

var repo = repoInit()

func (repo *StudentsData) Create(student Student) (Student, error) {
	student.Id = repo.maxID
	repo.maxID++
	repo.students[student.Id] = student
	return student, nil
}

func (repo *StudentsData) GetByID(id int) (Student, error) {
	student, exists := repo.students[id]
	if !exists {
		return Student{}, errors.New("student not found")
	}
	return student, nil
}

func (repo *StudentsData) Update(id int, student Student) (Student, error) {
	_, exists := repo.students[id]
	if !exists {
		return Student{}, errors.New("student not found")
	}
	student.Id = id
	repo.students[id] = student
	return student, nil
}

func (repo *StudentsData) Delete(id int) error {
	_, exists := repo.students[id]
	if !exists {
		return errors.New("student not found")
	}
	delete(repo.students, id)
	return nil
}

func (repo *StudentsData) GetAll() ([]Student, error) {
	var data []Student
	for _, student := range repo.students {
		data = append(data, student)
	}
	return data, nil
}

func addStudent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user Student
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid input"})
		fmt.Println(err)
		return
	}
	newst, err := repo.Create(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error creating student"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newst)
}

func getStudentByid(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid student ID"})
		return
	}

	student, err := repo.GetByID(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Student not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func updateStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid student ID"})
		return
	}

	var upd Student
	err = json.NewDecoder(r.Body).Decode(&upd)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid input"})
		return
	}

	student, err := repo.Update(id, upd)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Student not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

func deleteStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid student ID"})
		return
	}

	err = repo.Delete(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Student not found"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	students, err := repo.GetAll()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error retrieving students"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func main() {
	router := httprouter.New()
	router.POST("/students", addStudent)
	router.GET("/students/:id", getStudentByid)
	router.PUT("/students/:id", updateStudent)
	router.DELETE("/students/:id", deleteStudent)
	router.GET("/students", getAll)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error serving:", err)
	}
}
