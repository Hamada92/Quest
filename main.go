// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/Hamada92/Quest/config"
// 	"github.com/Hamada92/Quest/questions"
// 	_ "github.com/jackc/pgx/v5/stdlib"
// )

// type app struct {
// 	db  *sql.DB
// 	cfg config.AppConfig
// }

// func main() {
// 	err := run()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		os.Exit(1)
// 	}
// 	var repo questions.QuestionList
// 	repo = append(repo, questions.Question{Id: 1, Body: "whats your name?"})
// 	repPtr := &repo
// 	// questions = append(questions, questions.Question{id: "1", body: "whats your nme?"})
// 	http.HandleFunc("/questions", repPtr.Index)
// 	log.Fatal(http.ListenAndServe("localhost:3000", nil))
// }

// func run() (err error) {
// 	cfg, err := config.InitConfig()
// 	if err != nil {
// 		return err
// 	}

// 	app := app{cfg: cfg}

// 	err = app.db.Ping()
// 	if err != nil {
// 		return err
// 	}
// 	defer func(db *sql.DB) {
// 		err := db.Close()
// 		if err != nil {
// 			return
// 		}
// 	}(app.db)
// 	return
// }
