package main

import (
	"github.com/erikdubbelboer/fasthttp"
)
func requestHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf8")
	ctx.WriteString("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")
}

func main() {
	fasthttp.ListenAndServe("localhost:8081", requestHandler)
}