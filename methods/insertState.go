package methods

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"gamificacion/db"
)

// insertState from body
/*       FORMATO:
{
        "UserID": 0,
        "CourseID": 0,
        "Status": "",
        "Complete": false,
        "ActualPoints": 0
}
*/

func InsertState(w http.ResponseWriter, r *http.Request) {
	if (r.Method != "POST"){
		log.Println("400: Bad request in InsertState.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	if (reqBody == nil){
		log.Println("204: No content in InsertState.")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var Result db.States
	json.Unmarshal(reqBody, &Result)
	db.DB.Create(&Result);
	log.Println("200: OK in InsertState.")
	w.WriteHeader(http.StatusOK)
}
