package methods

import (
	"log"
	"net/http"
	"gamificacion/db"
)

//DELETE: id como parametro
func DeleteByID(w http.ResponseWriter, r *http.Request) {
	if (r.Method != "GET"){
		log.Println("400: Bad request in DelectByID.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, ok := r.URL.Query()["id"]
	if (!ok || len(id) < 1){
		log.Println("204: No content in DelectByID.")
		w.WriteHeader(http.StatusNoContent)
		return
	} else if (len(id) > 1){
		log.Println("414: Request uri too long in DelectByID.")
		w.WriteHeader(http.StatusRequestURITooLong)
		return
	}
	var state db.States;
	db.DB.First(&state, "id = ?", id[0])
	db.DB.Delete(&state)
	log.Println("200: OK in DelectByID.")
	w.WriteHeader(http.StatusOK)
}
