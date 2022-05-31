package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/mmhil/simple-go-math-service/pkg/domath"
)

func main() {
	// do math, the web server
	// this will listen on port :8000 for web requests.

	http.HandleFunc("/math", mathHandler)
	log.Println("Listening for requests at Http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func mathHandler(w http.ResponseWriter, r *http.Request) {
	num1str, ok1 := r.URL.Query()["num1"]
	num2str, ok2 := r.URL.Query()["num2"]

	if !ok1 || len(num1str[0]) < 1 {
		log.Println("Url Param 'num1' is missing")
		return
	}
	if !ok2 || len(num2str[0]) < 1 {
		log.Println("Url Param 'num2' is missing")
		return
	}

	num1, err1 := strconv.Atoi(num1str[0])
	num2, err2 := strconv.Atoi(num2str[0])

	if err1 != nil || err2 != nil {
		if err1 != nil {
			log.Println("num1 is not an integer")
			io.WriteString(w, "num 1 is not an integer\n")
		}
		if err2 != nil {
			log.Println("num2 is not an integer")
			io.WriteString(w, "num 2 is not an integer\n")
		}
		return
	}

	resultStringArray := domath.ReturnStringsOfAllMathFunctions(num1, num2)

	for i := 0; i < len(resultStringArray); i++ {
		io.WriteString(w, resultStringArray[i])
	}

}
