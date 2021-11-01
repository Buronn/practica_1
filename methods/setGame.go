package methods

import (
	"log"
	"net/http"
	"gamificacion/db"
	"strconv"
)

 //PATCH: user_id, course_id y game como parametro 
func SetGame(w http.ResponseWriter, r *http.Request) {
	if (r.Method != "GET") {
		log.Println("400: Bad request in SetGame.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user_id, ok := r.URL.Query()["user_id"]
	course_id, ok1 := r.URL.Query()["course_id"]
	game, ok2 := r.URL.Query()["game"]
	if (!ok || len(user_id) < 1 || !ok1 || len(course_id) < 1 || !ok2 || len(game) < 1){
		log.Println("204: No content in SetGame.")
		w.WriteHeader(http.StatusNoContent)
		return
	} else if (len(user_id) > 1 || len(course_id) > 1 || len(game) > 1){
		log.Println("414: Request uri too long in SetGame.")
		w.WriteHeader(http.StatusRequestURITooLong)
		return
	}
	var state db.States
	db.DB.First(&state, "user_id = ? AND course_id = ?", user_id[0], course_id[0])
	state.Game = game[0]
	if state.ID == 0{
		param2, _ := strconv.Atoi(user_id[0])
		param3, _ := strconv.Atoi(course_id[0])
		state.UserID = uint(param2)
		state.CourseID = uint(param3)
		db.DB.Create(&state)
	}else{
		db.DB.Save(&state)
	}
	log.Println("200: OK in SetGame.")
	w.WriteHeader(http.StatusOK)
}
