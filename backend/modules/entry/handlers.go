package entry

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"notes/infra"
)

var UserId = "fe1433f8-8576-4e04-87df-031778028bd5"

func GetEntries(w http.ResponseWriter, r *http.Request) {
	entries, err := findEntries(r.Context(), UserId)

	if err != nil {
		log.Printf("Unable to query entries %s", err)
		infra.SendResponse(w, "Unable to get entries", http.StatusInternalServerError)
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
	var entry Entry
	err := infra.ReadJSON(w, r, &entry)

	if err != nil {
		return
	}

	err = insertEntry(r.Context(), entry.Content, UserId)

	if err != nil {
		log.Printf("Error creating entry %s", err)
		infra.SendResponse(w, "Unable to create entry", http.StatusInternalServerError)
		return
	}

	infra.SendResponse(w, "Entry created successfully", http.StatusCreated)
}

func UpdateEntry(w http.ResponseWriter, r *http.Request) {
	entryId := chi.URLParam(r, "entryId")
	var entry Entry

	err := infra.ReadJSON(w, r, entry)

	if err != nil {
		return
	}

	err = updateEntry(r.Context(), entryId, UserId, entry.Content)

	if err != nil {
		infra.SendResponse(w, "Unable to update entry", http.StatusInternalServerError)
		return
	}

	infra.SendResponse(w, "Entry updated successfully", http.StatusOK)
}
