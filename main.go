package main

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("matrix")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		defer file.Close()

		records, err := csv.NewReader(file).ReadAll()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		var response string
		for _, row := range records {
			response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
		}
		fmt.Fprint(w, response)
	})
	http.HandleFunc("/invert", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("matrix")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		defer file.Close()

		records, err := csv.NewReader(file).ReadAll()
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
	})
	http.HandleFunc("/flatten", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("matrix")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		defer file.Close()

		records, err := csv.NewReader(file).ReadAll()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		var response string
		for _, row := range records {
			response = fmt.Sprintf("%s%s,", response, strings.Join(row, ","))
		}
		fmt.Fprint(w, response[:len(response)-1])
	})
	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("matrix")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		defer file.Close()

		records, err := csv.NewReader(file).ReadAll()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		//calling Math function to calculate the sum of the matrix IsSum = True
		response, err := Math(records, true)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Fprint(w, response)
	})
	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("matrix")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		defer file.Close()

		records, err := csv.NewReader(file).ReadAll()
		if err != nil {
			w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}
		//calling Math function to calculate the product of the matrix parameter IsSum = False
		response, err := Math(records, false)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Fprint(w, response)
	})
	http.ListenAndServe(":8080", nil)

}
