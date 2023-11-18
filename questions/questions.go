package questions

import (
	"encoding/json"
	"log"
	"net/http"
)

type Question struct {
	Id   int64
	Body string
}

type QuestionList []Question

func (ql *QuestionList) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		questionsJson, err := json.Marshal(ql)

		if err != nil {
			log.Fatal("error marshalling")
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(questionsJson)
	case "POST":
		if err := req.ParseForm(); err != nil {
			log.Fatal("error parsing form")
		}
		body := req.FormValue("body")
		list := *ql
		lastId := list[len(list)-1].Id
		q := Question{Id: lastId + 1, Body: body}
		*ql = append(*ql, q)

		questionJson, err := json.Marshal(q)
		if err != nil {
			log.Fatalf("error marshalling error: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(questionJson)
	default:
		log.Fatal("unsupported method")
	}

}

func (ql *QuestionList) addToList(q Question) QuestionList {
	return append(*ql, q)
}
