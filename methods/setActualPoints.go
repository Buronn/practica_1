package methods

import (
	"log"
	"net/http"
	"gamificacion/db"
	"strconv"
)

//PATCH: user_id, course_id y points como parametro 
func SetActualPoints(w http.ResponseWriter, r *http.Request) {
	if (r.Method != "GET"){
		log.Println("400: Bad request in SetActualPoints.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user_id, ok := r.URL.Query()["user_id"]
	course_id, ok1 := r.URL.Query()["course_id"]
	points, ok2 := r.URL.Query()["points"]
	if (!ok || len(user_id) < 1 || !ok1 || len(course_id) < 1 || !ok2 || len(points) < 1) {
		log.Println("204: No content in SetActualPoints.")
		w.WriteHeader(http.StatusNoContent)
		return
	} else if (len(user_id) > 1 || len(course_id) > 1 || len(points) > 1) {
		log.Println("414: Request uri too long in SetActualPoints.")
		w.WriteHeader(http.StatusRequestURITooLong)
		return
	}
	var state db.States
	db.DB.First(&state, "user_id = ? AND course_id = ?", user_id[0], course_id[0])
	param, _ := strconv.Atoi(points[0])
	if state.ActualPoints < param{ // solo si lo que recibe es mayor a lo que estaba lo actualiza
		state.ActualPoints = param
		db.DB.Save(&state)
	}else if state.ID == 0{
		param2, _ := strconv.Atoi(user_id[0])
		param3, _ := strconv.Atoi(course_id[0])
		state.UserID = uint(param2)
		state.CourseID = uint(param3)
		db.DB.Create(&state)
	}else{
		log.Println("400: Bad request in SetActualPoints.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("200: OK in SetActualPoints.")
	w.WriteHeader(http.StatusOK)
}
