package methods

import (
	"encoding/json"
	"log"
	"net/http"
	"gamificacion/db"
	"gamificacion/structures"
)

func CourseRanking(w http.ResponseWriter, r *http.Request) { //GET: course_id como parametro
	if (r.Method != "GET"){
		log.Println("400: Bad request in CourseRanking.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	course_id, ok := r.URL.Query()["course_id"]
	if (!ok || len(course_id) < 1){
		log.Println("204: No content in CourseRanking.")
		w.WriteHeader(http.StatusNoContent)
		return
	} else if (len(course_id) > 1){
		log.Println("414: Request uri too long in CourseRanking.")
		w.WriteHeader(http.StatusRequestURITooLong)
		return
	}
	var result []structures.GetPointsResult
	db.DB.Model(&db.States{}).Select("user_id, sum(actual_points) as total_points").Where("course_id = ?", course_id[0]).Group("user_id").Find(&result)
	if (len(result) == 0){
		log.Println("204: No content in CourseRanking.")
		w.WriteHeader(http.StatusNoContent)
		return
	} 
	w.Header().Set("Content-Type", "application/json")
	log.Println("200: OK in CourseRanking.")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)	
}
