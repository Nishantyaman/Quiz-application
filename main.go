package main

import(
	"fmt"
	"net/http"

	"quiz/handlersa"
)

func main(){
	http.HandleFunc("/questions",handlersa.HandleQuestions)
	http.HandleFunc("/answers",handlersa.HandleAnswer)
	http.HandleFunc("/score",handlersa.HandleScore)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080",nil)

}