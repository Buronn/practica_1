package methods

import (
	"log"
	"net/http"
	"gamificacion/db"
	"strconv"
)

 //PATCH: user_id, course_id y complete como parametro 
func SetComplete(w http.ResponseWriter, r *http.Request) {
	if (r.Method != "GET") {
		log.Println("400: Bad request in SetComplete.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user_id, ok := r.URL.Query()["user_id"]
	course_id, ok1 := r.URL.Query()["course_id"]
	complete, ok2 := r.URL.Query()["complete"]
	if (!ok || len(user_id) < 1 || !ok1 || len(course_id) < 1 || !ok2 || len(complete) < 1){
		log.Println("204: No content in SetComplete.")
		w.WriteHeader(http.StatusNoContent)
		return
	} else if (len(user_id) > 1 || len(course_id) > 1 || len(complete) > 1){
		log.Println("414: Request uri too long in SetComplete.")
		w.WriteHeader(http.StatusRequestURITooLong)
		return
	}
	var state db.States
	db.DB.First(&state, "user_id = ? AND course_id = ?", user_id[0], course_id[0])
	state.Complete, _ = strconv.ParseBool(complete[0])
	if state.ID == 0{
		param2, _ := strconv.Atoi(user_id[0])
		param3, _ := strconv.Atoi(course_id[0])
		state.UserID = uint(param2)
		state.CourseID = uint(param3)
		db.DB.Create(&state)
	}else{
		db.DB.Save(&state)
	}
	log.Println("200: OK in SetComplete.")
	w.WriteHeader(http.StatusOK)
}
