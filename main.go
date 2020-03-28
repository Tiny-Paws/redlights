package main

import (
	"fmt"
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
	fmt.Printf("%v", r.Pattern())
	for {
		fmt.Printf("%v\n", redlight.ColorState(r, time.Now()))
		time.Sleep(1 * time.Second)
	}
	g := gin.Default()
	g.GET("/status", status)
}

func status(g *gin.Context) {

}
