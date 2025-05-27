package handlersa

import(
	"net/http"
	"encoding/json"
	"strconv"
	"sync"
	"strings"

	"quiz/models"
	"quiz/utils"
)

var (
	questions = []models.Question{
		{ID: 1, Question: "What is 2 + 2?", Options: []string{"3", "4", "5"}, Answer: "4"},
		{ID: 2, Question: "Capital of India?", Options: []string{"Mumbai", "Delhi", "Chennai"}, Answer: "Delhi"},
	}
	scores = make(map[string]int)
	mutex = &sync.Mutex{}
)

func HandleQuestions(w http.ResponseWriter,r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Only get supported",http.StatusMethodNotAllowed)
		return
	}

	utils.WriteJSON(w,questions)
}

func HandleAnswer(w http.ResponseWriter,r *http.Request){
	if r.Method != "POST"{
		http.Error(w,"Only POST supported",http.StatusMethodNotAllowed)
		return
	}
	var req models.AnswerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil{
		http.Error(w,"Invalid JSON",http.StatusBadRequest)
		return
	}
	
	var correct bool
	for _,q := range questions{
		if q.ID == req.QuestionID{
			correct = strings.EqualFold(req.Answer,q.Answer)
			break 
		}
	}

	resp:= models.AnswerResponse{
		Correct : correct,
		Message : "Wrong!",
	}

	if correct {
		resp.Message = "Correct!"
		clientID := r.RemoteAddr
		mutex.Lock()
		scores[clientID]++
		mutex.Unlock()
	}

	utils.WriteJSON(w,resp)
}

func HandleScore(w http.ResponseWriter,r *http.Request){
	if r.Method != "GET"{
		http.Error(w,"Only GET supported",http.StatusMethodNotAllowed)
		return
	}
	clientID := r.RemoteAddr
	mutex.Lock()
	score:=scores[clientID]
	mutex.Unlock()
	utils.WriteJSON(w,map[string]string{"score":strconv.Itoa(score)})
}
