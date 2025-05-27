package models

type Question struct{
    ID int `json:"id"`
    Question string `json:"question"`
    Options []string `json:"options"`
    Answer string `json:"-"`
}

type AnswerRequest struct{
    QuestionID int `json:"question_id"`
    Answer string `json:"answer"`
}

type AnswerResponse struct{
    Correct bool `json:"correct"`
    Message string `json:"message"`
}


