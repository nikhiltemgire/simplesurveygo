package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	mgo "gopkg.in/mgo.v2"
)

type Movies struct {
	Movies []Movies `json:"movies"`
}

type Movie struct {
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Director string `json:"director"`
	Cast     string `json:"cast"`
	Genre    string `json:"string"`
	Notes    string `json:"notes"`
}

var movies []Movie
var MgoSession *mgo.Session

func Init() {
	if MgoSession == nil {
		var err error
		MgoSession, err = mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
	}
}

func f(payload Movie) {
	session := MgoSession.Clone()
	defer session.Close()

	sessionClctn := session.DB("Movies").C("movie_details")

	sessionClctn.Insert(payload)
}

func main() {
	url := "https://raw.githubusercontent.com/prust/wikipedia-movie-data/master/movies.json"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	go Init()

	body, _ := ioutil.ReadAll(res.Body)
	//text := string(body)
	//fmt.Println(body)

	json.Unmarshal(body, &movies)

	//fmt.Println(movies)
	var wg sync.WaitGroup

	movie_len := len(movies)
	for i := 0; i < movie_len; i++ {
		var payload_1 = movies[0]
		movies = movies[1:len(movies)]
		var payload_2 = movies[0]
		movies = movies[1:len(movies)]
		var payload_3 = movies[0]
		movies = movies[1:len(movies)]
		var payload_4 = movies[0]
		movies = movies[1:len(movies)]

		go f(payload_1)
		go f(payload_2)
		go f(payload_3)
		go f(payload_4)

		wg.Wait()
	}
}
