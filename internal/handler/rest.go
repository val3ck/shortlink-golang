package rest

import (
	"encoding/json"
	"fmt"
	environment "golang-api/config"
	structures "golang-api/internal/model"
	encodeS "golang-api/internal/service"
	"net/http"
)

func createLink(w http.ResponseWriter, r *http.Request) { // TODO: invalid body link error
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
		http.Error(w, "Body error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	encoded := encodeS.EncodingURL(link.Link, 12)
	link.Code = encoded
	// TODO: ENCODE AND SEND
	// TODO: add to a (Postgre)
	// SUCCESS

}
