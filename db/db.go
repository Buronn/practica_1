package db

import(
    "os"
	"log"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "time"
)

var DB *gorm.DB

func init(){
    var dsn string
    dsn = "host=172.18.0.2 user=postgres password=postgres DB.name=postgres port=5432 sslmode=disable" // os.Getenv("db")
    db, err := gorm.Open(postgres.New(postgres.Config{
        DSN: dsn,
        PreferSimpleProtocol: true,
      }), &gorm.Config{})
    if (err != nil){
        log.Fatal("Error",err)
        os.Exit(1)
    }
	db.AutoMigrate(&States{}, &Results{})
    // test data:
   
    
	tests := []States{
        {UserID: 1, CourseID: 1, Complete: false, Progress: 0, ActualPoints: 0, Game: "puzzle", StartTime: time.Now()},
        {UserID: 1, CourseID: 1, Complete: false, Progress: 0, ActualPoints: 0, Game: "maze", StartTime: time.Now()},
        {UserID: 1, CourseID: 1, Complete: false, Progress: 0, ActualPoints: 0, Game: "bird", StartTime: time.Now()},
        {UserID: 1, CourseID: 1, Complete: false, Progress: 0, ActualPoints: 0, Game: "turtle", StartTime: time.Now()},
        {UserID: 1, CourseID: 1, Complete: false, Progress: 0, ActualPoints: 0, Game: "movie", StartTime: time.Now()},
        {UserID: 1, CourseID: 1, Complete: false, Progress: 0, ActualPoints: 0, Game: "music", StartTime: time.Now()},
        {UserID: 1, CourseID: 1, Complete: false, Progress: 0, ActualPoints: 0, Game: "pond-tutor", StartTime: time.Now()},
        {UserID: 2, CourseID: 1, Complete: false, Progress: 0, ActualPoints: 0, Game: "puzzle", StartTime: time.Now()},
        {UserID: 2, CourseID: 1, Complete: false, Progress: 0, ActualPoints: 0, Game: "maze", StartTime: time.Now()}}
	db.Create(&tests)
	log.Println(tests)
    // end test data
    DB = db
}