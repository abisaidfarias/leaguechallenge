package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/invert", invertHandler)
	http.HandleFunc("/flatten", flattenHandler)
	http.HandleFunc("/sum", sumHandler)
	http.HandleFunc("/multiply", multiplyHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {

	//calling a function to read the file
	records, err := ReadFile(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}
func invertHandler(w http.ResponseWriter, r *http.Request) {

	//calling a function to read the file
	records, err := ReadFile(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response string
	var matrixInvert = Invert(records)
	for _, row := range matrixInvert {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response)
}
func flattenHandler(w http.ResponseWriter, r *http.Request) {

	//calling a function to read the file
	records, err := ReadFile(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s,", response, strings.Join(row, ","))
	}
	fmt.Fprint(w, response[:len(response)-1])
}
func sumHandler(w http.ResponseWriter, r *http.Request) {

	//calling a function to read the file
	records, err := ReadFile(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response string
	//calling Math function to calculate the sum of the matrix IsSum = True
	response, err = Math(records, true)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Fprint(w, response)
}
func multiplyHandler(w http.ResponseWriter, r *http.Request) {

	//calling a function to read the file
	records, err := ReadFile(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	var response string
	//calling Math function to calculate the product of the matrix parameter IsSum = False
	response, err = Math(records, false)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Fprint(w, response)
}
