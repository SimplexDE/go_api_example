package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/go-chi/chi"
	"github.com/SimplexDE/go_api_example/internal/handlers"
	logg "github.com/sirupsen/logrus"
)

func main() {

	logg.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println(`
   __________     ___    ____  ____
  / ____/ __ \   /   |  / __ \/  _/
 / / __/ / / /  / /| | / /_/ // /  
/ /_/ / /_/ /  / ___ |/ ____// /   
\____/\____/  /_/  |_/_/   /___/   `)

	fmt.Println(" ")
	log.Println("Starting GO API service...")

	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		logg.Error(err)
	}
}