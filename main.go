package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "log"
    "github.com/thelmholtz/astroweather/weather"
)

func main(){
    router := mux.NewRouter()

    router.Path("/clima").Queries("dia", "{dia}").HandlerFunc(Weather)

    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatal(err)
    }
}

func Weather(w http.ResponseWriter, r *http.Request) {

    day := r.FormValue("dia")
    value, err := weather.Day(day)
    if err != nil {
        log.Print("Error retrieving weather by day", err)
        w.Write([]byte("{\n\t\"error\": \"" + err.Error() + "\"\n\t}"))
        return
    }

    w.Write([]byte("{\n\t\"dia\": " + day + ",\n\t\"clima\": \"" + value + "\"\n}")) //TODO This needs a less half-assed rewriting

}
