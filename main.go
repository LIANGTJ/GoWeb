package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprintf(w, "Welcome to the home page!")
	// })

	// n := negroni.Classic()
	// n.UseHandler(mux)
	// n.Run(":3000")
	//router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)

	n := negroni.New(Middleware1, Middleware2)
	// Or use a middleware with the Use() function
	n.Use(Middleware3)
	// router goes last
	n.UseHandler(router)

	n.Run(":3000")
}
