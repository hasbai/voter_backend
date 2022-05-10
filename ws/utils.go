package ws

import "encoding/json"

func BroadcastObject(_type string, obj interface{}) {
	data, _ := json.Marshal(&Message{
		Type: _type,
		Data: obj,
	})
	Manager.Broadcast <- data
}
