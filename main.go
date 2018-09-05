package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    var listen_port string = ":" + os.Getenv("LISTEN_PORT")
    router := mux.NewRouter()
    router.HandleFunc("/healthcheck", HealthCheck).Methods("GET")
    log.Fatal(http.ListenAndServe(listen_port, router))
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
  appname := os.Getenv("APP_NAME")
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, "Hey Bro, %v is Alive!", appname)
}

