package wsTools

import (
	"encoding/binary"
	"github.com/gorilla/websocket"
)

type MsgTransceiver struct {
	conn *websocket.Conn
}


var ErrorMessageNotBinary = Error{
	Err: "Error, message is not binary",
}

var ErrorMessageWrongFormat = Error{
	Err: "Error, message has wrong format",
}

func (w *MsgTransceiver) Send(eventName *string, message *[]byte) error {
	return w.SendMessage(&WsMessage{
		EventName: eventName,
		Message:   message,
	})
}

func (w *MsgTransceiver) SendMessage(msg *WsMessage) error {
	l := len(*msg.EventName)
	msgL := len(*msg.Message)
	bs := make([]byte, IntSize + l + msgL)

	binary.BigEndian.PutUint64(bs, uint64(l))
	return w.conn.WriteMessage(websocket.BinaryMessage, bs)
}

func (w *MsgTransceiver) Receive() (*WsMessage, error){
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
	eventString := string(message[IntSize: eventStringLength + IntSize])
	eventData := message[eventStringLength + IntSize:]
	return &WsMessage{
		EventName: &eventString,
		Message: &eventData,
	}, nil
}

func (w * MsgTransceiver) close(){
	w.conn.Close()
}