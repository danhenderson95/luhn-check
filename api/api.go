package api

import (
	"danhenderson95/luhncheck/luhn"
	"encoding/json"
	"fmt"
	"net/http"
)

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	value := r.URL.Query().Get("value")

	if value == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	passesLuhnCheck := luhn.LuhnCheck(value)

	returnValue := map[string]bool{
		"passes": passesLuhnCheck,
	}

	returnJSON, err := json.Marshal(returnValue)

	if err != nil {
		fmt.Printf("api-error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(returnJSON)
}
