package websocket

type IO struct {
	hub *Hub
}

func NewIO(hub *Hub) *IO {
	return &IO{hub}
}

func (i *IO) Broadcast(data []byte) {
	i.hub.broadcast <- data
}

func (i *IO) To(id string, data []byte) {
	i.hub.clients[id].send <- data
}
