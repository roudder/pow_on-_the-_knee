package main

import (
	"fmt"
	"github.com/roudder/powServer/pow"
	"log"
	"math/rand"
	"net/http"
)

//Jack Chan
var quotes = map[int]string{
	0: "Don’t try to be like Jackie. There is only one Jackie. Study computers instead.",
	1: "I’m crazy, but I’m not stupid",
	2: "I only want my work to make people happy.",
	3: "I’m good for some things, bad for a lot of things.",
}

func main() {
	http.HandleFunc("/req", handler)
	log.Fatal(http.ListenAndServe(":9797", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	hcHeader := r.Header.Get("X-Hashcash")
	if len(hcHeader) == 0 {
		w.Write([]byte("where is hashcash"))
	}
	//TODO just do pow check with string (not convers. into struct)
	hc, err := pow.NewHashCash(hcHeader)
	if err != nil {
		w.WriteHeader(400)
	}
	proof := pow.NewProofOfWork(*hc)

	if proof.Validate() {
		fmt.Println("pow is okay")
		w.WriteHeader(200)
		w.Write([]byte(quotes[rand.Intn(3)]))
	} else {
		fmt.Println("try again")
		w.WriteHeader(400)
		w.Write([]byte("you should do it better"))
	}
}
