package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/murawakimitsuhiro/go-simple-RESTful-api/driver"
	"github.com/murawakimitsuhiro/go-simple-RESTful-api/models"
	"github.com/murawakimitsuhiro/go-simple-RESTful-api/repository"
	"github.com/murawakimitsuhiro/go-simple-RESTful-api/repository/note"
)

func NewNoteHandler(db *driver.DB) *Note {
	return &Note{
		repo: note.NewGormRepo(db.DB),
	}
}

type Note struct {
	repo repository.NoteRepo
}

func (n *Note) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := n.repo.Fetch(r.Context(), 100)

	respondWithJSON(w, http.StatusOK, payload)
}

func (n *Note) Create(w http.ResponseWriter, r *http.Request) {
	note := models.Note{}
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		respondWithError(w, http.StatusBadRequest, "Request Error")
		return
	}

	if _, err := n.repo.Create(r.Context(), &note); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

func (n *Note) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data := models.Note{}
	data.ID = uint(id)

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		respondWithError(w, http.StatusBadRequest, "Request Error")
		return
	}

	payload, err := n.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	respondWithJSON(w, http.StatusOK, payload)
}

func (n *Note) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := n.repo.GetByID(r.Context(), uint(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
		return
	}

	respondWithJSON(w, http.StatusOK, payload)
}

func (n *Note) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := n.repo.Delete(r.Context(), uint(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	respondWithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"message": msg})
}
