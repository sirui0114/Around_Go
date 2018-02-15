package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"strconv"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Post struct {
	// `json:"user"` is for the json parsing of this User field. Otherwise, by default it's 'User'.
	User     string `json:"user"`
	Message  string  `json:"message"`
	Location Location `json:"location"`
}

const (
	DISTANCE = "200km"
)

func main() {
	fmt.Println("started-service")
	http.HandleFunc("/post", handlerPost)
	http.HandleFunc("/search", handlerSearch)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func handlerSearch(w http.ResponseWriter, r *http.Request){
	fmt.Println("Received one request for search")
	lat := r.URL.Query().Get("lat")
	lt, _ := strconv.ParseFloat(lat, 64)
	lon := r.URL.Query().Get("lon")
	ln, _:= strconv.ParseFloat(lon, 64)


	ran := DISTANCE
	if val := r.URL.Query().Get("range"); val != ""{
		ran = val + "km"
	}

	fmt.Fprintf(w, "Search received: %f %f %s", lat, lon, ran)

	p := &Post{
		User:"1111",
		Message:"一生必去的100个地方",
		Location: Location{
			Lat:lt,
			Lon:ln,
		},
	}

	js, err := json.Marshal(p)
	if err != nil {
		panic(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)


}

func handlerPost(w http.ResponseWriter, r *http.Request) { //传指针快
	fmt.Println("Received a requires for post")
	decoder := json.NewDecoder(r.Body)
	var p Post
	if err := decoder.Decode(&p); err != nil{
		panic(err)
		return
	}
	fmt.Fprintf(w, "Post received: %s\n", p.Message)

}

