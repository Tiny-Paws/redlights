package main

import (
	"encoding/json"
	"net/http"
	"time"

	"gitlab.com/shizzuru/redlights/pkg/sse"

	"gitlab.com/shizzuru/redlights/pkg/redlight"
)

func main() {
	r := redlight.NewRedlight(time.Now())
	r.SetPattern(redlight.ColorPattern{
		[]int{redlight.Red, 15},
		[]int{redlight.Green, 8},
		[]int{redlight.Orange, 3},
		// []int{redlight.Orange | redlight.Red, 1},
	})

	broker := sse.NewServer()

	go func() {
		for {
			time.Sleep(time.Millisecond * 100)
			yup := map[string]int{
				"state": redlight.ColorState(r, time.Now()),
			}
			a, _ := json.Marshal(&yup)
			broker.Notifier <- []byte(a)
		}
	}()
	http.HandleFunc("/", handler)
	http.HandleFunc("/status", broker.ServeHTTP)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//r.Header.Set("Access-Control-Allow-Origin", "*")
	http.FileServer(http.Dir("website")).ServeHTTP(w, r)
}
