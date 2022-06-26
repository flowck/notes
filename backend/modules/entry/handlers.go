package entry

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetEntries(w http.ResponseWriter, r *http.Request) {
	entries, err := findEntries(r.Context(), "fe1433f8-8576-4e04-87df-031778028bd5")

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

func CreateEntry(w http.ResponseWriter, r *http.Request) {

}
