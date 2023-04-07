package api
import (
	"outsourcing/module"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var connection = make(map[string]*websocket.Conn)
func Queryy(ctx *gin.Context){
	authHeader := ctx.Query("Authorization")
	if !module.Checkjwt(authHeader){
		ctx.JSON(200, Response{PARAMETER_ERROR, myjson{} , "token不合法"})
		return
	}
	ws ,err := upgrader.Upgrade(ctx.Writer,ctx.Request, nil )
	if err !=nil{
		fmt.Println(err)
		return 
	}
	userinfo :=module.ParseJwt(authHeader)
	username := userinfo.UserName
	connection[username] = ws
	defer func (ws *websocket.Conn)  {
		delete(connection,username)
		ws.Close()
	}(ws)
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		var m WebsocketAuctionReceive
		err = json.Unmarshal(message, &m)
		if err != nil{
			err = ws.WriteMessage(mt, []byte(WEBSOCKET_RECEIVE_FORMAT_ERROR))
			if err != nil {
				fmt.Println(err)
				break
			}
			break
		}else{
			a := module.QueryAuction(int(m.NftId))
			module.Updateuction(int(m.NftId),int(userinfo.UserId),a.EndPrice+m.Price)
			for _, v := range connection {
				payload := WebsocketAuctionPayload{username,m.Price,float32(a.EndPrice+m.Price)}
				err = v.WriteJSON(payload)
				if err != nil {
					fmt.Println(err)
					break
				}
			}
		}
	
	}
}