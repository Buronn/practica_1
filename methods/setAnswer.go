package methods

import (
	"encoding/json"
	"gamificacion/db"
	"gamificacion/functions"
	"gamificacion/judge"
	"gamificacion/structures"
	"io/ioutil"
	"log"
	"net/http"
)

//POST: user_id, course_id, key y value en el body
func SetAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("400: Bad request in SetAnswer.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	if reqBody == nil {
		log.Println("204: No content in SetAnswer.")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var body structures.Body
	json.Unmarshal(reqBody, &body)

	var state db.States
	game, _ := functions.ParseFlight(body.Key)
	db.DB.First(&state, "user_id = ? AND course_id = ? AND game = ?", body.UserID, body.CourseID, game) // busca el estado del juego "game" de un "usuario" en el curso "x"

	if state.ID != 0 {
		go func() {
			judge.Body_values <- body
			judge.State_values <- state
		}()

		log.Println("200: OK in SetAnswer.")
		w.WriteHeader(http.StatusOK)
	} else {
		log.Println("400: Bad Request in SetAnswer.")
		w.WriteHeader(http.StatusBadRequest)
	}
}
