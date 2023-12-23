package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	server := http.FileServer(http.Dir("./static"))

	http.Handle("/", server)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting Server at http://localhost:8080")
	
	isServerUp := http.ListenAndServe(":8080", nil)
	if isServerUp != nil {
		log.Fatal(isServerUp)
	}
}

func formHandler(resp http.ResponseWriter, req *http.Request) {

	form := req.ParseForm()
	if form != nil{
		fmt.Fprintf(resp,"Parse Form error : %v ",form)
		return
	}

	name := req.FormValue("name")
	age := req.FormValue("age")

	fmt.Fprintf(resp,"Name : %s\n",name)
	fmt.Fprintf(resp,"Age : %s\n",age)
}

func helloHandler(resp http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(resp, "Hello!")
}
