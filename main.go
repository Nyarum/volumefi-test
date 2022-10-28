package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func calculateHandler(w http.ResponseWriter, req *http.Request) {
	var airportsRelations [][]string

	err := json.NewDecoder(req.Body).Decode(&airportsRelations)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	src, dst := calculateDistance1(airportsRelations)

	res, err := json.Marshal([]string{src, dst})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(res)
}

func main() {
	http.HandleFunc("/calculate", calculateHandler)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
