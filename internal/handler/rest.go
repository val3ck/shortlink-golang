package rest

import (
	"encoding/json"
	"fmt"
	environment "golang-api/config"
	structures "golang-api/internal/model"
	"net/http"
)

func createLink(w http.ResponseWriter, r *http.Request) {
	secretKey := "Bearer " + environment.EnvGet("BEARER_KEY")
	if userAuth := r.Header.Get("Authorization"); userAuth == "" || userAuth != secretKey {
		// fmt.Fprintf(w, "Bad auth %d", http.StatusUnauthorized)
		http.Error(w, "Authorization error", http.StatusUnauthorized)
		return
	}
	var link structures.Links
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&link)
	fmt.Print(err)
	if err != nil {
		// fmt.Fprintf(w, "Bad body %d", http.StatusBadRequest)
		http.Error(w, "Body error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// TODO: generate random link ID , set it's own id(maybe based on db)
	// TODO: send response (random link)
	// TODO: add to a (Postgre)
	// SUCCESS

}
