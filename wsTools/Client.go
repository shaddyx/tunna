package wsTools
import (
	"github.com/gorilla/websocket"
	"net/url"
)

type wsClient struct {
	MsgTransceiver
	conn *websocket.Conn
}

//
//	Creates new websocket client
//
func NewWsClient(urlString string) (*wsClient, error) {

	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &wsClient{
		conn: c,
	}

	return res, nil
}