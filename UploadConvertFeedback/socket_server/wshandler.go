package socket_server

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
	"fmt"
	"sync"
	si "github.com/TechMaster/LearnGo/UploadConvertFeedback/socket_interface"
	"github.com/TechMaster/LearnGo/UploadConvertFeedback/encode_server"
)
type SocketServer struct {
	ws *websocket.Server
	totalFrameChannel chan si.TotalFrameMessage
	currentFrameChannel chan si.CurrentFrameMessage
	mutex *sync.Mutex
	connections map[string] websocket.Connection
}
var socketServer *SocketServer
var once sync.Once



func GetSocketServer() *SocketServer {
	return socketServer
}
//Singleton pattern
func New(app *iris.Application) (*SocketServer) {
	once.Do(func() {
		socketServer = &SocketServer {
			ws: websocket.New(websocket.Config{}),
			totalFrameChannel: make (chan si.TotalFrameMessage),
			currentFrameChannel: make (chan si.CurrentFrameMessage),
			mutex: new(sync.Mutex),
			connections: make(map[string] websocket.Connection),
		}

		// register the server on an endpoint.
		// see the inline javascript code i the websockets.html, this endpoint is used to connect to the server.
		app.Get("/ws_endpoint", socketServer.ws.Handler())

		// serve the javascript built'n client-side library,
		// see websockets.html script tags, this path is used.
		app.Any("/iris-ws.js", func(ctx iris.Context) {
			ctx.Write(websocket.ClientSource)
		})
		socketServer.ws.OnConnection(socketServer.handleConnection)
	})

	go func() {
		for {
			select {
			case msg := <- socketServer.totalFrameChannel:
				socketServer.connections[msg.WSConnID].Emit("TotalFrame", msg.TotalFrame)
			case msg := <- socketServer.currentFrameChannel:
				socketServer.connections[msg.WSConnID].Emit("CurrentFrame", msg.CurrentFrame)
			}
		}
	}()

	return socketServer
}

func  (ss *SocketServer) handleConnection(c websocket.Connection) {
	fmt.Println(c.ID())
	ss.mutex.Lock()
	ss.connections[c.ID()] = c
	ss.mutex.Unlock()

	c.On("Encode", func(videoFile string) {
		es := encode_server.GetEncodeServer()
		es.EnqueueVideoEncodeRequest(videoFile, c.ID())
	})

	c.OnDisconnect(func() {
		ss.mutex.Lock()
		delete(ss.connections, c.ID())
		ss.mutex.Unlock()
		fmt.Printf("\nConnection %s has been disconnected!\n", c.ID())
	})

}

func (socketServer *SocketServer) NotifyTotalFrame(msg si.TotalFrameMessage) {
	socketServer.totalFrameChannel <- msg
}

func (socketServer *SocketServer) NotifyCurrentFrame(msg si.CurrentFrameMessage) {
	socketServer.currentFrameChannel <- msg
}