package wsTools
import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
)

type wsServer struct {
	MsgTransceiver
	conn *websocket.Conn
}

var upgrader = websocket.Upgrader{}
//
//	Creates new websocket server
//
func NewWsServer(w *http.ResponseWriter, r *http.Request) (*wsServer, error){
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return nil, err
	}
	res := &wsServer{
		conn: c,
	}

	return res, nil
}