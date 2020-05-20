package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	thisAppDir := filepath.Dir(os.Args[0])
	thisAppFullPath,_ := filepath.Abs(thisAppDir)

	domainCertFile := filepath.Join(thisAppFullPath, "domain.crt")
	domainKeyFile := filepath.Join(thisAppFullPath, "domain.key")

	// start randomizer
	rand.Seed(time.Now().UnixNano())

	// new gin
	newGin := gin.New()

	// redirect from root
	newGin.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/static/")
	})

	// static content (HTML)
	newGin.Static("/static/", filepath.Join(thisAppFullPath, "static"))

	// workers api - PULL
	newGin.GET("/api/workers", workersApi)

	// push API - HTTP2 streams
	newGin.GET("/api/stream", streamApi)

	// push API - WebSockets
	newGin.GET("/ws", socketApi)

	log.Println("Starting all API's (except WebSocket) on TCP/8443 (HTTPS Alternative)")
	newGin.RunTLS(":8443", domainCertFile, domainKeyFile)
}

func workersApi(context *gin.Context) {
	context.JSON(200, sse.Event{
		Event: 	"timer.event",
		Id:     fmt.Sprintf("%d", time.Now().UnixNano()),
		Data:      fmt.Sprintf("Time is %s", time.Now().Format(time.RFC822)),
	})
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		time.Sleep(time.Duration(1 + rand.Intn(10)) *time.Second)
		content,_ := json.Marshal(sse.Event{
			Event: 	"timer.event",
			Id:     fmt.Sprintf("%d", time.Now().UnixNano()),
			Data:      fmt.Sprintf("Time is %s", time.Now().Format(time.RFC822)),
		})
		conn.WriteMessage(websocket.TextMessage, content)
	}
}

func socketApi(context *gin.Context) {
	socketHandler(context.Writer, context.Request)
}

func streamApi(context *gin.Context) {
	context.Stream(func(w io.Writer) bool {
		for {
			time.Sleep(time.Duration(1 + rand.Intn(10))*time.Second)
			sse.Encode(w, sse.Event{
				Event: 	"timer.event",
				Id:     fmt.Sprintf("%d", time.Now().UnixNano()),
				Data:      fmt.Sprintf("Time is %s", time.Now().Format(time.RFC822)),
			})
			return true
		}
	})
}

