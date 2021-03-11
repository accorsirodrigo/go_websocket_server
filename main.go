package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

type responseObject struct {
	Response string
	Received string
	Created  time.Time
}

type lampBroadcast struct {
	Code   string `json:"code"`
	Status bool   `json:"status"`
}

func main() {
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	m := melody.New()

	//r.Use(static.Serve("/", static.LocalFile("./public", true/)))
	r.GET("/", func(c *gin.Context) {
		fmt.Println()
		code := c.Request.URL.Query().Get("code")
		dataType := c.Request.URL.Query().Get("type")
		status := c.Request.URL.Query().Get("status")
		var response lampBroadcast

		if dataType == "lamp" {
			response = lampBroadcast{
				Code:   code,
				Status: status == "on",
			}
		} else {
			response = lampBroadcast{
				Code:   "",
				Status: false,
			}
		}

		var jsonData []byte
		jsonData, err := json.Marshal(response)
		if err != nil {
			fmt.Println("error")
		}

		m.Broadcast(jsonData)

		c.JSON(200, response) // gin.H is a shortcut for map[string]interface{}
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Println(string(msg))
		response := responseObject{
			Response: "data: " + string(msg),
			Received: string(msg),
			Created:  time.Now(),
		}
		var jsonData []byte
		jsonData, err := json.Marshal(response)
		if err != nil {
			fmt.Println("ooops")
		}
		m.Broadcast(jsonData)
	})

	m.HandleConnect(func(s *melody.Session) {
		fmt.Println("Connected!")
		fmt.Println(s.Request.URL)
		fmt.Println(s.Request.Header)
	})

	m.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("Closed!")
	})

	r.Run(":5000")
}

func tryit() {

}
