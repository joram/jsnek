package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func Static(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	router := httprouter.New()
	router.POST("/start", Start)
	router.POST("/move", Move)
	router.POST("/end", End)
	router.POST("/ping", Ping)
	router.GET("/debug", Debug)
	router.GET("/debug/:gameID", Debug)
	router.GET("/static/*filepath", Static)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	// Add filename into logging messages
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("Running server on port %s...\n", port)
	http.ListenAndServe(":"+port, LoggingHandler(router))
}
