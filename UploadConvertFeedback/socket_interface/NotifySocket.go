package socket_interface

type TotalFrameMessage struct{
	TotalFrame int
	WSConnID string
}

type CurrentFrameMessage struct {
	CurrentFrame int
	WSConnID string
}

type ISocketNotify interface {
	NotifyTotalFrame(msg TotalFrameMessage)
	NotifyCurrentFrame(msg CurrentFrameMessage)
}