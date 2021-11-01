package methods

import (
	"encoding/json"
	"log"
	"net/http"
	"gamificacion/db"
	"gamificacion/structures"
)

func GameRanking(w http.ResponseWriter, r *http.Request) { //GET: Game_id como parametro
	if (r.Method != "GET"){
		log.Println("400: Bad request in GameRanking.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	game, ok := r.URL.Query()["game"]
	if (!ok || len(game) < 1){
		log.Println("204: No content in GameRanking.")
		w.WriteHeader(http.StatusNoContent)
		return
	} else if (len(game) > 1){
		log.Println("414: Request uri too long in GameRanking.")
		w.WriteHeader(http.StatusRequestURITooLong)
		return
	}
	var result []structures.GetPointsResult
	db.DB.Model(&db.States{}).Select("user_id, sum(actual_points) as total_points").Where("game = ?", game[0]).Group("user_id").Find(&result)
	if (len(result) == 0){
		log.Println("204: No content in GameRanking.")
		w.WriteHeader(http.StatusNoContent)
		return
	} 
	w.Header().Set("Content-Type", "application/json")
	log.Println("200: OK in GameRanking.")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&result)	
}
