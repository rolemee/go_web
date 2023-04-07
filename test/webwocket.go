package test
import (

	"fmt"

	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gorilla/websocket"
)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
  // gin.SetMode(gin.ReleaseMode)
  r := gin.Default()
  r.GET("/",func(ctx *gin.Context) {
	fmt.Println("success")
	ws ,err := upgrader.Upgrade(ctx.Writer,ctx.Request, nil )
	if err !=nil{
		fmt.Println(err)
		return 
	}
	defer ws.Close()
	for {
		//Read Message from client
		mt, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
	fmt.Println(string(message))
		//If client message is ping will return pong
		
			message = []byte("pong")
		
		//Response message to client
		err = ws.WriteMessage(mt, message)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
  })
  r.Run()

}