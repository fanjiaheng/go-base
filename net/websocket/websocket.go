package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Data struct {
	Ip       string   `json:"ip"`
	User     string   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
}

var h = hub{
	c: make(map[*connection]bool),
	b: make(chan []byte),
	r: make(chan *connection),
	u: make(chan *connection),
}

var wu = &websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func MainWebSocket() {

	router := mux.NewRouter()
	go h.run()

	router.HandleFunc("/ws", myws)

	fmt.Println("server running ... ")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("err: ", err)
	}
}

func myws(w http.ResponseWriter, r *http.Request) {
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	c := &connection{
		sc:   make(chan []byte, 256),
		ws:   ws,
		data: &Data{},
	}

	h.r <- c
	go c.writer()
	c.reader()

	defer func() {
		c.data.Type = "logout"
		user_list = del(user_list, c.data.User)
		c.data.UserList = user_list
		c.data.Content = c.data.User
		data_b, _ := json.Marshal(c.data)
		h.b <- data_b
		h.r <- c
	}()
}

var user_list = []string{}

func del(slice []string, user string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}

	var n_slice = []string{}
	for i := range slice {
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			n_slice = append(slice[:i], slice[i+1:]...)
			break
		} else {

		}
	}
	fmt.Println(n_slice)
	return n_slice
}

// 连接处理相关
type connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *Data
}

func (c *connection) writer() {
	for m := range c.sc {
		c.ws.WriteMessage(websocket.TextMessage, m)
	}
	c.ws.Close()
}

func (c *connection) reader() {
	for {
		_, m, err := c.ws.ReadMessage()
		if err != nil {
			h.r <- c
			break
		}
		json.Unmarshal(m, &c.data)
		switch c.data.Type {
		case "login":
			c.data.User = c.data.Content
			c.data.From = c.data.User
			user_list = append(user_list, c.data.User)
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		case "user":
			c.data.Type = "user"
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		case "logout":
			c.data.Type = "logout"
			user_list = del(user_list, c.data.User)
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
			h.r <- c
		default:
			fmt.Println("*******************default********************")
		}
	}
}

// hub相关
type hub struct {
	c map[*connection]bool
	b chan []byte
	r chan *connection
	u chan *connection
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.r:
			h.c[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()

		}
	}
}
