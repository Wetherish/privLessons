package db

import (
	"database/sql"
	"fmt"
	"privLessons/models"

	"github.com/go-playground/locales/id"
	_ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "User"
    password = "Password"
    dbname   = "LessonDB"
)

var database* sql.DB
func ConnectDB()  {

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    database, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }else{
        fmt.Println("Connect to DB")
    }
    defer database.Close()

    err = database.Ping()
    if err != nil {
        panic(err)
    }

}

func GetStudentByID(int64 id) (models.Student, error) {
    var student models.Student
    rows, err := database.Query("SELECT * FROM studentList  WHERE ID = ?", id)
    if err != nil{
        return nil, fmt.Errorf()
    }
}
