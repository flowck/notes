package http

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"notes/internal/common"
	entry "notes/internal/entry/domain"
	"notes/internal/entry/service"
)

var UserId = "fe1433f8-8576-4e04-87df-031778028bd5"

type EntryHttpHandler struct {
	service *service.EntryService
}

func NewEntryHttpHandler(service *service.EntryService) EntryHttpHandler {
	return EntryHttpHandler{service}
}

func (h *EntryHttpHandler) GetEntries(w http.ResponseWriter, r *http.Request) {
	entries, err := h.service.QueryEntries(r.Context(), UserId)

	if err != nil {
		log.Printf("Unable to query entries %s", err)
		common.SendResponse(w, "Unable to get entries", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(entries)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func (h *EntryHttpHandler) GetEntry(w http.ResponseWriter, r *http.Request) {
	entryId := chi.URLParam(r, "entryId")
	entry, err := h.service.QueryEntry(r.Context(), UserId, entryId)

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

func (h *EntryHttpHandler) CreateEntry(w http.ResponseWriter, r *http.Request) {
	var entry entry.Entry
	err := common.ReadJSON(w, r, &entry)

	if err != nil {
		return
	}

	err = h.service.AddNewEntry(r.Context(), UserId, entry.Content)

	if err != nil {
		log.Printf("Error creating entry %s", err)
		common.SendResponse(w, "Unable to create entry", http.StatusInternalServerError)
		return
	}

	common.SendResponse(w, "Entry created successfully", http.StatusCreated)
}

/* func (h *EntryHttpHandler) UpdateEntry(w http.ResponseWriter, r *http.Request) {
	entryId := chi.URLParam(r, "entryId")
	var entry domain.Entry

	err := infra.ReadJSON(w, r, entry)

	if err != nil {
		return
	}

	err = entry.updateEntry(r.Context(), entryId, UserId, entry.Content)

	if err != nil {
		infra.SendResponse(w, "Unable to update entry", http.StatusInternalServerError)
		return
	}

	infra.SendResponse(w, "Entry updated successfully", http.StatusOK)
}*/
