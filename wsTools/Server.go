package wsTools
import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"encoding/binary"
)
const (
	IntSize = 8
)

type wsServer struct {
	writer *http.ResponseWriter
	request *http.Request
	conn *websocket.Conn
}

type WsMessage struct {
	EventName string
	Message   []byte
}

var ErrorMessageNotBinary = Error{
	Err: "Error, message is not binary",
}

var ErrorMessageWrongFormat = Error{
	Err: "Error, message has wrong format",
}


func (w *wsServer) send(msgType string, data []byte) error {

}

func (w *wsServer) sendMessage(msg WsMessage) error{

}

func (w *wsServer) receive() (*WsMessage, error){
	mt, message, err := w.conn.ReadMessage()
	if err != nil{
		return nil, err
	}
	if mt != websocket.BinaryMessage{
		return nil, ErrorMessageNotBinary
	}
	if len(message) < IntSize {
		return nil, ErrorMessageWrongFormat
	}
	eventStringLength := binary.BigEndian.Uint64(message[0: IntSize])
	eventString := string(message[IntSize: eventStringLength])
}

func (w * wsServer) close(){

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
		writer: w,
		request:r,
		conn: c,
	}

	return res, nil
}