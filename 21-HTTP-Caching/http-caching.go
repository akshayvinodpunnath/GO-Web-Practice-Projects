package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	CONNECT_HOST = "localhost"
	CONNECT_PORT = "8080"
)

var newCache *cache.Cache

//creating new cache and setting its expiry
func init() {
	newCache = cache.New(5*time.Minute, 10*time.Minute)
	newCache.Set("foo", "bar", cache.DefaultExpiration)
}

func getFromCache(w http.ResponseWriter, r *http.Request) {
	foo, found := newCache.Get("foo")
	if found {
		log.Print("Key Found in Cache with value as :: ", foo.(string))
		fmt.Fprintf(w, "Hello "+foo.(string))
	} else {
		log.Print("Key Not Found in Cache :: ", "foo")
		fmt.Fprintf(w, "Key Not Found in Cache")
	}
}

func main() {
	http.HandleFunc("/", getFromCache)
	err := http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
