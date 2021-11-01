package main

import (
	"log"
	"net/http"

	"gamificacion/methods"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("/app/blockly-games/"))) //TEST ENSEÑA
	/* http.Handle("/", http.FileServer(http.Dir("./blockly-games/"))) */ //LOCAL

	http.HandleFunc("/completetime", methods.CompleteTime)       //Recibe {"user_id", "course_id"}
	http.HandleFunc("/setprogress", methods.SetProgress)         //Recibe{"user_id", "course_id" y "progress"}
	http.HandleFunc("/setcomplete", methods.SetComplete)         //Recibe{"user_id", "course_id" y "complete"}
	http.HandleFunc("/setgame", methods.SetGame)                 //Recibe{"user_id", "course_id" y "game"}
	http.HandleFunc("/deletestates", methods.DeleteByID)         //Recibe{"id"}
	http.HandleFunc("/setactualpoints", methods.SetActualPoints) //Recibe{"user_id", "course_id" y "points"}
	http.HandleFunc("/insertstate", methods.InsertState)         //Recibe un body formato JSON
	http.HandleFunc("/courseranking", methods.CourseRanking)     //Recibe{"course_id"}
	http.HandleFunc("/gameranking", methods.GameRanking)
	http.HandleFunc("/setanswer", methods.SetAnswer)           // Recibe un body formato JSON
	http.HandleFunc("/shareinstagram", methods.MovieInstagram) //Recibe Body

	log.Fatal(http.ListenAndServe(":3000", nil))
}
