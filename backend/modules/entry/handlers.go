package entry

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"net/http"
)

var UserId = "fe1433f8-8576-4e04-87df-031778028bd5"

func GetEntries(w http.ResponseWriter, r *http.Request) {
	entries, err := findEntries(r.Context(), UserId)

	if err != nil {
		fmt.Errorf("Unable to query entries %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(entries)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func GetEntry(w http.ResponseWriter, r *http.Request) {
	entryId := chi.URLParam(r, "entryId")
	entry, err := findEntryById(r.Context(), entryId, UserId)

	if err != nil {
		fmt.Errorf("Unable to query entry %w", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(entry)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func CreateEntry(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please send a valid json."))
		return
	}

	var entry Entry
	err = json.Unmarshal(body, &entry)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please send a valid json."))
		return
	}

	err = insertEntry(r.Context(), entry.Content, UserId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateEntry(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	entryId := chi.URLParam(r, "entryId")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please send a valid json."))
		return
	}

	var entry Entry
	err = json.Unmarshal(body, &entry)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please send a valid json."))
		return
	}

	err = updateEntry(r.Context(), entryId, UserId, entry.Content)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
