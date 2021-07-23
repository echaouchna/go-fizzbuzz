package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

func getFizzbuzzResponse(w http.ResponseWriter, limit int, options FizzbuzzOptions) {
	retval := fizzbuzz(limit, options)
	fmt.Fprintf(w, "%s", retval)
}

func parseOptions(r *http.Request) FizzbuzzOptions {
	v := r.URL.Query()
	options := FizzbuzzOptions{
		Int1: parseInt(v.Get("int1")),
		Int2: parseInt(v.Get("int2")),
		Str1: v.Get("str1"),
		Str2: v.Get("str2"),
	}
	options.fillDefaults()
	return options
}

func fizzbuzzWrapper(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	options := parseOptions(r)
	limit := parseInt(vars["limit"])
	validate := validator.New()
	err := validate.Struct(options)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		responseBody := map[string]string{"error": validationErrors.Error()}
		json.NewEncoder(w).Encode(responseBody)
		return
	}
	go func() { statsChannel <- options }()
	getFizzbuzzResponse(w, limit, options)
}

func mostUsed(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	// fmt.Fprintf(w, "%s", getMostUsed())
	json.NewEncoder(w).Encode(getMostUsed())
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/calculate/{limit}", fizzbuzzWrapper)
	myRouter.HandleFunc("/most_used", mostUsed)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
