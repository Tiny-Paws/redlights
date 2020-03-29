package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"gitlab.com/shizzuru/redlights/internal/redlight"
)

func main() {
	r := redlight.NewRedlight(time.Now())
	r.SetPattern(redlight.ColorPattern{
		[]int{redlight.Red, 15},
		[]int{redlight.Green, 8},
		[]int{redlight.Orange, 3},
		[]int{redlight.Orange | redlight.Red, 1},
	})
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.GET("/status", func(g *gin.Context) {
		g.Header("Access-Control-Allow-Origin", "*")
		g.JSON(http.StatusOK, gin.H{
			"state": redlight.ColorState(r, time.Now()),
		})
	})
	go func() {
		err := g.Run(":5600")
		if err != nil {
			panic(err)
		}
	}()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//r.Header.Set("Access-Control-Allow-Origin", "*")
	http.FileServer(http.Dir("website")).ServeHTTP(w, r)
}
