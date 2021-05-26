package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/julienschmidt/httprouter"

	"github.com/kucinghitam13/go-openapi/model"
)

var (
	people = map[int64]model.Person{
		1: model.Person{
			ID:      1,
			Name:    "Ethan Winters",
			Age:     37,
			Address: "Depok",
		},
		2: model.Person{
			ID:      2,
			Name:    "Chris Redfield",
			Age:     48,
			Address: "Jakarta",
		},
		3: model.Person{
			ID:      3,
			Name:    "Leon Scott Kennedy",
			Age:     44,
			Address: "Bandung",
		},
	}
	idCounter int64 = int64(len(people))
	mutex     sync.Mutex
)

// GetPersons godoc
// @Tags persons
// @Summary Get list of persons.
// @Description Get list of persons saved in db.
// @Produce json
// @Success 200 {object} model.GetPersonsResponse "Success"
// @Failure 500 "Something went wrong"
// @Router /persons [get]
func GetPersons(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	resp := model.GetPersonsResponse{
		Total:   len(people),
		Persons: make([]model.Person, 0, len(people)),
	}
	for _, person := range people {
		resp.Persons = append(resp.Persons, person)
	}
	writeJSON(w, http.StatusCreated, resp)
}

// AddPerson godoc
// @Tags persons
// @Summary Add person.
// @Description Add person to db.
// @Accept json
// @Produce json
// @Param body body model.AddPersonRequest true "Request payload."
// @Success 201 {object} model.AddPersonResponse "Success"
// @Success 400 "Bad request"
// @Failure 500 "Something went wrong"
// @Router /persons [post]
func AddPerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var param model.AddPersonRequest
	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := getNextID()
	person := model.Person{
		ID:      id,
		Name:    param.Name,
		Age:     param.Age,
		Address: param.Address,
	}

	mutex.Lock()
	people[id] = person
	mutex.Unlock()

	resp := model.AddPersonResponse{
		PersonMutationResponse: model.PersonMutationResponse{
			Operation: "ADD",
			Success:   true,
			Person:    person,
		},
	}

	writeJSON(w, http.StatusOK, resp)
}

// EditPerson godoc
// @Tags persons
// @Summary Edit person.
// @Description Edit existing person in db.
// @Accept json
// @Produce json
// @Param id path string true "Person ID."
// @Param body body model.EditPersonRequest true "Request payload."
// @Success 201 {object} model.EditPersonResponse "Success"
// @Success 400 "Bad request"
// @Success 404 "Person not found"
// @Failure 500 "Something went wrong"
// @Router /persons/id/{id} [put]
func EditPerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var param model.EditPersonRequest
	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mutex.Lock()
	person, exists := people[id]
	if !exists {
		mutex.Unlock()
		w.WriteHeader(http.StatusNotFound)
		return
	}
	person.Name = param.Name
	person.Age = param.Age
	person.Address = param.Address
	people[id] = person
	mutex.Unlock()

	resp := model.AddPersonResponse{
		PersonMutationResponse: model.PersonMutationResponse{
			Operation: "EDIT",
			Success:   true,
			Person:    person,
		},
	}

	writeJSON(w, http.StatusOK, resp)
}

// DeletePerson godoc
// @Tags persons
// @Summary Delete person.
// @Description Delete existing person in db.
// @Produce json
// @Param id path string true "Person ID."
// @Success 201 {object} model.DeletePersonResponse "Success"
// @Success 400 "Bad request"
// @Success 404 "Person not found"
// @Failure 500 "Something went wrong"
// @Router /persons/id/{id} [delete]
func DeletePerson(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mutex.Lock()
	person, exists := people[id]
	if !exists {
		mutex.Unlock()
		w.WriteHeader(http.StatusNotFound)
		return
	}
	delete(people, id)
	mutex.Unlock()

	resp := model.AddPersonResponse{
		PersonMutationResponse: model.PersonMutationResponse{
			Operation: "DELETE",
			Success:   true,
			Person:    person,
		},
	}

	writeJSON(w, http.StatusOK, resp)
}

func getNextID() int64 {
	return atomic.AddInt64(&idCounter, 1)
}
