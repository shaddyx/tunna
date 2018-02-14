package wsTools
import (
	"github.com/gorilla/websocket"
)

type wsStruct struct {
	upgrader websocket.Upgrader
}

func (w *wsStruct) send(msgType string, data []byte){

}







