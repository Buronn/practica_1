package evaluate

import (
	"gamificacion/db"
	"log"
	"time"

	"gamificacion/functions"
	"gamificacion/structures"
)

func All(body_values chan structures.Body, state_values chan db.States) {
	for {
		body := <-body_values
		state := <-state_values

		var result db.Results
		result.StatesID = state.ID
		result.Key = body.Key
		result.Value = body.Value

		t0 := time.Now()
		name, level := functions.ParseFlight(body.Key)
		log.Println("waiting" + name)
		var answer bool
		answer = false
		var bleft int
		bleft = -1
		if name == "puzzle" {
			log.Println("Entra a Puzzle")
			answer = Puzzle(body)
		} else if name == "maze" || name == "bird" || name == "music" || name == "pond-tutor" || name == "turtle" {
			log.Println("Entra a Others")
			answer, bleft = Others(body)
		} else if name == "movie" {
			log.Println("Entra a Movie")
			if level != "10" {
				answer = Movie(body)
			}
		}
		t1 := time.Now()
		if answer {
			log.Println("Correcto en", t1.Sub(t0))
			if bleft == -1 {
				result.Points = 10 + 100/(1+int(t1.Sub(t0))/1000000000)
			} else {
				result.Points = 10 + 10*bleft
			}

		} else {
			log.Println("Incorrecto en", t1.Sub(t0))
			// si falla por timeout seria bueno buscar la forma de volver a encolar
		}

		db.DB.Create(&result)

		if answer {
			var progress int
			subQuery := db.DB.Select("key").Where("states_id = ? and points > 0", state.ID).Table("results").Group("key")
			db.DB.Select("count(key) as progress").Table("(?) as completes_keys", subQuery).Find(&progress)
			state.Progress = progress
			if (state.Progress == 10 || body.Key == "puzzle1") && !state.Complete {
				state.Complete = true
				state.CompleteTime = time.Now()
			}

			var actual_points int
			subQuery = db.DB.Select("max(points) as points").Where("states_id = ?", state.ID).Table("results").Group("key")
			db.DB.Select("sum(points) as actual_points").Table("(?) as keys_points", subQuery).Find(&actual_points)
			state.ActualPoints = actual_points

			db.DB.Save(&state)
		}
	}
}
